package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/encuestas-go/back-enc/internal/domain"
	"github.com/stretchr/testify/assert"
)

// Successful cases for now
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

func TestInsert(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepository := InitializeUserRepository(db)

	user := domain.User{
		Name:        "Lorena",
		MiddleName:  "Valle",
		LastName:    "Gonzalez",
		Email:       "lvg@gmail.com",
		PhoneNumber: "7771234567",
		Username:    "lvg18",
		Password:    "hola123",
		IDUserType:  1,
	}

	mock.ExpectExec(`INSERT INTO USUARIO`).
		WithArgs(user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username,
			user.Password, user.IDUserType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = userRepository.Insert(user)

	// EXPECT
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepository := InitializeUserRepository(db)

	user := domain.User{
		ID:          1,
		Name:        "Lorena",
		MiddleName:  "Valle",
		LastName:    "Flores",
		Email:       "lvg@gmail.com",
		PhoneNumber: "1234567",
		Username:    "lvg109",
		Password:    "123456",
		IDUserType:  1,
	}

	mock.ExpectExec("UPDATE USUARIO").
		WithArgs(user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username,
			user.Password, user.IDUserType, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// THEN
	err = userRepository.Update(user, user.ID)

	// EXPECT
	assert.NoError(t, err)

}

func TestDelete(t *testing.T) {
	// GIVEN
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userRepository := InitializeUserRepository(db)

	userID := 1

	mock.ExpectExec("DELETE FROM USUARIO").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// THEN
	err = userRepository.Delete(userID)

	// EXPECT
	assert.NoError(t, err)
}
