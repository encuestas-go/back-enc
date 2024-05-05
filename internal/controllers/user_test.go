package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/encuestas-go/back-enc/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	t.Run("Create User succesfully", func(t *testing.T) {
		// GIVEN
		user := domain.User{
			Name:        "Flora",
			MiddleName:  "Lopez",
			LastName:    "Gonzalez",
			Email:       "flolg@gmail.com",
			PhoneNumber: "7771234567",
			Username:    "flo18",
			Password:    "hola123",
			IDUserType:  1,
		}

		// Server
		e := echo.New()
		userJSON, err := json.Marshal(user)
		assert.NoError(t, err)

		request := httptest.NewRequest(http.MethodPost, "/crear/usuario", strings.NewReader(string(userJSON))) // Create a request to pass to our handler.
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)                                   //Create a ResponseRecorder to record the response.
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// Database
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		userRepository := repository.InitializeUserRepository(db)

		mock.ExpectExec(`INSERT INTO USUARIO`).
			WithArgs(user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username,
				user.Password, user.IDUserType).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// WHEN
		err = InitUserController(userRepository).Create((echoContext))

		// THEN
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)

			expectedControllerMessage := ControllerMessageResponse{
				StatusCode: http.StatusCreated,
				Message:    fmt.Sprintf("Creater %v successfully", user.Username),
			}
			expectedBody, err := json.Marshal(expectedControllerMessage)
			assert.Nil(t, err)

			assert.Contains(t, recorder.Body.String(), string(expectedBody))
		}
	})

	t.Run("Create user fails due to invalid JSON", func(t *testing.T) {
		// GIVEN
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/crear/usuario", strings.NewReader(`{"}`))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder := httptest.NewRecorder()
		echoContext := e.NewContext(request, recorder)

		// WHEN
		err := InitUserController(nil).Create((echoContext))

		// THEN
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	})
}

func TestUpdateUser(t *testing.T) {

}
