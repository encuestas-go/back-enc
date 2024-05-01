package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepository := InitializeUserRepository(db)

	user := domain.UserLogin{
		Email:    "hnco20027@gmail.com",
		Password: "123",
	}

	rows := sqlmock.NewRows([]string{"ID", "ID_TIPO_USUARIO"}).AddRow(1, 1)
	mock.ExpectQuery(`SELECT ID, ID_TIPO_USUARIO FROM USUARIO`).WillReturnRows(rows)

	// THEN
	idUser, idTypeUser, err := userRepository.Login(user)

	// EXPECT
	assert.NoError(t, err)
	assert.Equal(t, 1, idUser)
	assert.Equal(t, 1, idTypeUser)
}
