package repository

import (
	"database/sql"
	"log"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type UserRepositoryService struct {
	db *sql.DB
}

func InitializeUserRepository(db *sql.DB) *UserRepositoryService {
	return &UserRepositoryService{
		db: db,
	}
}

func (u *UserRepositoryService) Insert(user domain.User) error {
	result, err := u.db.Exec(`
	INSERT INTO USUARIO (NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, CORREO_ELECTRONICO, NUMERO_TELEFONO, USUARIO, CONTRASENA, ID_TIPO_USUARIO)
    VALUES(?, ?, ?, ?, ?, ?, ?, ?);
	`, user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.IDUserType)
	if err != nil {
		log.Println("Unable to insert into the USER table, the error is: ", err)
		return nil
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return nil
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to User table")
		return nil
	} else if rowsInserted == 0 {
		log.Println("Cannot add data into the User table", err)
		return nil
	}

	return nil
}

func (u *UserRepositoryService) Update(user domain.User, id int) error {
	result, err := u.db.Exec(`
	UPDATE USUARIO SET NOMBRE = ?,
                   APELLIDO_PATERNO = ?,
                   APELLIDO_MATERNO = ?,
                   NUMERO_TELEFONO = ?,
                   USUARIO = ?,
                   ID_TIPO_USUARIO = ?
                   WHERE ID = ?;
	`, user.Name, user.MiddleName, user.LastName, user.PhoneNumber, user.Username, user.IDUserType, id)
	if err != nil {
		log.Println("Data could not be updated into User table, the error was:", err)
		return nil
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return nil
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into User table")
		return nil
	} else if rowsUpdated == 0 {
		log.Println("Data could not be update into User table")
		return nil
	}

	return nil
}

func (u *UserRepositoryService) Delete(user domain.User, id int) error {
	result, err := u.db.Exec("DELETE FROM USUARIO WHERE ID = ?;", id)
	if err != nil {
		log.Println("Could not delete the id on user table, the error was: ", err)
		return nil
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Println("Could not delete information with the requested id: ", err)
		return nil
	}

	if rowsDeleted > 0 {
		log.Printf("ID %v was successfully deleted from the user table", id)
		return nil
	} else if rowsDeleted == 0 {
		log.Println("Could not delete the requested ID in the user table.")
		return nil
	}
	return nil
}

func (u *UserRepositoryService) Get(user domain.User) error {
	return nil
}
