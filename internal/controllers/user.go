package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// UserController it's a struct for User
type UserController struct {
	UserRepository *repository.UserRepositoryService
}

func InitUserController(repo *repository.UserRepositoryService) *UserController {
	return &UserController{
		UserRepository: repo,
	}
}

type userLoginResponse struct {
	IDUser     int `json:"id_user,omitempty"`
	IDTypeUser int `json:"id_type_user,omitempty"`
}

func (u *UserController) Login(c echo.Context) error {
	userLogin := domain.UserLogin{}
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while binding the request body: %v", err),
		})
	}

	idUser, idTypeUser, err := u.UserRepository.Login(userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("An error happened when trying to login, the body is: %v, the error is: %v", userLogin, err),
		})
	}

	if idUser == 0 || idTypeUser == 0 {
		log.Printf("ID user %d ,ID Type User %d", idUser, idTypeUser)
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid email or password provided",
		})
	}

	idUserConverted := strconv.Itoa(idUser)
	idTypeUserConverted := strconv.Itoa(idTypeUser)

	expires := time.Now().Add(24 * time.Hour)

	cookieIDUser := new(http.Cookie)
	cookieIDUser.Name = "id_user"
	cookieIDUser.Value = idUserConverted
	cookieIDUser.Expires = expires
	cookieIDUser.Path = "/"
	// cookieIDUser.Domain = "127.0.0.1"
	cookieIDUser.HttpOnly = false
	cookieIDUser.Secure = false
	cookieIDUser.SameSite = http.SameSiteLaxMode

	cookieIDTypeUser := new(http.Cookie)
	cookieIDTypeUser.Name = "id_type_user"
	cookieIDTypeUser.Value = idTypeUserConverted
	cookieIDTypeUser.Expires = expires
	cookieIDTypeUser.Path = "/"
	// cookieIDTypeUser.Domain = "127.0.0.1"
	cookieIDTypeUser.HttpOnly = false
	cookieIDTypeUser.Secure = false
	cookieIDTypeUser.SameSite = http.SameSiteLaxMode

	c.SetCookie(cookieIDUser)
	c.SetCookie(cookieIDTypeUser)

	return c.JSON(http.StatusOK, userLoginResponse{
		IDUser:     idUser,
		IDTypeUser: idTypeUser,
	})
}

func (u *UserController) ResetPassword(c echo.Context) error {
	email := c.QueryParam("email")
	if len(email) == 0 {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Email is required"),
		})
	}

	if err := u.sendEmailToUser(c, email); err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when sending email to user: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Password reset successfully",
	})
}

func (u *UserController) Create(c echo.Context) error {
	user := domain.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = u.UserRepository.Insert(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened when trying to insert the user, the body is: %v, the error is: %v", user, err),
		})
	}

	return c.JSON(http.StatusCreated, ControllerMessageResponse{
		StatusCode: http.StatusCreated,
		Message:    fmt.Sprintf("Created %v successfully", user.Username),
	})
}

func (u *UserController) Update(c echo.Context) error {
	user := domain.User{}
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid ID user requested: %v", err),
		})
	}

	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error happened trying to bind the body, err: %v", err),
		})
	}

	err = u.UserRepository.Update(user, userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred when trying to update the user: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("User with ID %s successfully updated", userID),
	})
}

func (u *UserController) Delete(c echo.Context) error {
	userID := c.QueryParam("user_id")
	userIDConverted, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Invalid userID requested: %v", err),
		})
	}
	err = u.UserRepository.Delete(userIDConverted)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("An error occurred while trying to delete the user: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("User with ID %s successfully deleted", userID),
	})
}

func (u *UserController) Get(c echo.Context) error {
	userIDString := c.QueryParam("id")
	if userIDString == "" {
		userIDString = "0"
	}

	getOnlyStudents := false
	onlyStudents := c.QueryParam("only_students")
	if onlyStudents == "true" {
		getOnlyStudents = true
	}

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ControllerMessageResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid input data: %s", err),
		})
	}

	res, err := u.UserRepository.GetAllOrByID(userID, getOnlyStudents)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ControllerMessageResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to get information of USER: %v", err),
		})
	}

	return c.JSON(http.StatusOK, ControllerMessageResponse{
		StatusCode: http.StatusOK,
		Message:    "Information of the user succesfully retrieved",
		Data:       res,
	})
}

func (u *UserController) sendEmailToUser(c echo.Context, email string) error {
	smtpHost := "mail.privateemail.com"
	smtpPort := "587"
	username := os.Getenv("smtp_username")
	password := os.Getenv("smtp_password")

	from := username
	to := []string{email}

	rand.Seed(time.Now().UnixNano())

	tmpPassword := strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10)) +
		strconv.Itoa(rand.Intn(10))

	err := u.UserRepository.UpdateOnlyPassword(email, tmpPassword)
	if err != nil {
		return err
	}

	message := []byte(fmt.Sprintf("To: %v \r\n", email) +
		fmt.Sprintf("Subject: ¡Cambio de contraseña!\r\n") +
		"\r\n" +
		fmt.Sprintf("Usa esta contraseña: %v para entrar a tu cuenta.\r\n", tmpPassword))

	auth := smtp.PlainAuth("", username, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserController) getUserCookie(c echo.Context) (int, int) {
	cookieIDUser, err := c.Cookie("id_user")
	if err != nil {
		return 0, 0
	}

	IDUserConverted, err := strconv.Atoi(cookieIDUser.Value)
	if err != nil {
		return 0, 0
	}

	cookieIDTypeUser, err := c.Cookie("id_type_user")
	if err != nil {
		return 0, 0
	}

	IDTypeUserConverted, err := strconv.Atoi(cookieIDTypeUser.Value)
	if err != nil {
		return 0, 0
	}

	return IDUserConverted, IDTypeUserConverted
}
