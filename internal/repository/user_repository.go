package repository

import (
	"database/sql"
	"errors"
	"fmt"
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

func (u *UserRepositoryService) Login(userlogin domain.UserLogin) (int, int, error) {
	var (
		id           int
		id_user_type int
	)

	err := u.db.QueryRow(`
	SELECT ID, ID_TIPO_USUARIO FROM USUARIO WHERE CORREO_ELECTRONICO = ? AND CONTRASENA = SHA2(?,256);
	`, userlogin.Email, userlogin.Password).Scan(&id, &id_user_type)
	if err != nil {
		log.Println("Could not retrieve information with requested email and password")
		return 0, 0, err
	}

	return id, id_user_type, nil
}

func (u *UserRepositoryService) Insert(user domain.User) error {
	result, err := u.db.Exec(`
	INSERT INTO USUARIO (NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, CORREO_ELECTRONICO, NUMERO_TELEFONO, USUARIO, CONTRASENA, ID_TIPO_USUARIO)
    VALUES(?, ?, ?, ?, ?, ?, SHA2(?,256), ?);
	`, user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.IDUserType)
	if err != nil {
		log.Println("Unable to insert into the USER table, the error is: ", err)
		return err
	}

	rowsInserted, err := result.RowsAffected()
	if err != nil {
		log.Println("Unable to obtained values from the added columns", err)
		return err
	}

	if rowsInserted > 0 {
		log.Printf("Data successfully added to User table")
		return nil
	} else if rowsInserted == 0 {
		return errors.New("cannot add data into the user table")
	}

	return nil
}

func (u *UserRepositoryService) Update(user domain.User, id int) error {
	result, err := u.db.Exec(`
	UPDATE USUARIO SET NOMBRE = ?,
                   APELLIDO_PATERNO = ?,
                   APELLIDO_MATERNO= ?,
                   CORREO_ELECTRONICO = ?,
                   NUMERO_TELEFONO = ?,
                   USUARIO = ?,
                   CONTRASENA = SHA2(?,256),
                   ID_TIPO_USUARIO = ?
                   WHERE ID = ?;
	`, user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber, user.Username, user.Password, user.IDUserType, id)
	if err != nil {
		log.Println("Data could not be updated into User table, the error was:", err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("Cannot update values on the corresponding columns", err)
		return err
	}

	if rowsUpdated > 0 {
		log.Println("Successfully updated into User table")
		return nil
	} else if rowsUpdated == 0 {
		return errors.New("data could not be update into User table")
	}

	return nil
}

func (u *UserRepositoryService) Delete(id int) error {
	// Start a transaction
	tx, err := u.db.Begin()
	if err != nil {
		log.Println("Could not start transaction: ", err)
		return err
	}

	// Define the delete queries for related tables
	relatedTables := []string{
		"ENCUESTA_ACTIVIDAD",
		"ENCUESTA_SATISFACCION",
		"ENCUESTA_SERVICIO",
		"ENCUESTA_INFRAESTRUCTURA_HOGAR",
		"FORO_PREGUNTA",
		"FORO_RESPUESTA",
		"ENCUESTA_NIVEL_SOCIOECONOMICO",
		"ENCUESTA_TRANSPORTE",
		"ENCUESTA_NIVEL_DEMOGRAFICO",
		"ENCUESTA_NIVEL_ECONOMICO",
		"PUBLICACION_EVENTO",
	}

	for _, table := range relatedTables {
		query := fmt.Sprintf("DELETE FROM %s WHERE ID_USUARIO = ?;", table)
		_, err = tx.Exec(query, id)
		if err != nil {
			tx.Rollback()
			log.Printf("Could not delete references in table %s: %v", table, err)
			return err
		}
	}

	// Now delete the user
	result, err := tx.Exec("DELETE FROM USUARIO WHERE ID = ?;", id)
	if err != nil {
		tx.Rollback()
		log.Println("Could not delete the id on user table, the error was: ", err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		log.Println("Could not delete information with the requested id: ", err)
		return err
	}

	if rowsDeleted > 0 {
		err = tx.Commit()
		if err != nil {
			log.Println("Could not commit transaction: ", err)
			return err
		}
		log.Printf("ID %v was successfully deleted from the user table", id)
		return nil
	} else if rowsDeleted == 0 {
		tx.Rollback()
		return errors.New("could not delete the requested ID in the user table")
	}

	return nil
}

func (u *UserRepositoryService) UpdateOnlyPassword(email string, password string) error {
	res, err := u.db.Exec(`UPDATE USUARIO SET CONTRASENA = SHA2(?,256) WHERE CORREO_ELECTRONICO = ?;`, password, email)
	if err != nil {
		return err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsUpdated == 0 {
		return errors.New("could not update the password in the user table")
	}

	return nil
}

func (u *UserRepositoryService) GetAllOrByID(id int, getOnlyStudents bool) ([]domain.User, error) {
	var query = `SELECT ID, NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, CORREO_ELECTRONICO, NUMERO_TELEFONO, 
				USUARIO,ID_TIPO_USUARIO FROM USUARIO;`

	if getOnlyStudents {
		query = fmt.Sprintf(`
			SELECT ID, NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, CORREO_ELECTRONICO, NUMERO_TELEFONO,
							USUARIO,ID_TIPO_USUARIO FROM USUARIO WHERE ID_TIPO_USUARIO = 2;
		`)
	}

	if id > 0 {
		query = fmt.Sprintf(`SELECT ID, NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, CORREO_ELECTRONICO,
							 NUMERO_TELEFONO, USUARIO,ID_TIPO_USUARIO FROM USUARIO WHERE ID = %v;`, id)
	}

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		if err = rows.Scan(&user.ID, &user.Name, &user.MiddleName, &user.LastName, &user.Email,
			&user.PhoneNumber, &user.Username, &user.IDUserType); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
