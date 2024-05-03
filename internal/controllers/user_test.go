package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{
		"name" : "Lorena ",
		"middle_name" : "Valle",
		"last_name" : "Gonzalez",
		"email" : "lore1@gmail.com",
		"phone_number" :"7771338587",
		"username":"lore123",
		"password" :"1234",
		"id_user_type": 2
	}`
)

func TestCreateUser(t *testing.T) {

	t.Run("Create User succesfully", func(t *testing.T) {
		//GIVEN
		e := echo.New()
		// Create a request to pass to our handler.
		request := httptest.NewRequest(http.MethodPost, "/crear/usuario", strings.NewReader(userJSON))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		//Create a ResponseRecorder to record the response.
		recorder := httptest.NewRecorder()

		echoContext := e.NewContext(request, recorder)
		userController := InitUserController()

		//THEN
		err := userController.Create(echoContext)

		//EXPECT
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, recorder.Code)
		}
	})
}
