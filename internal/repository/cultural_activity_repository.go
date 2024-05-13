package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type CulturalActivityRepositoryService struct {
	db *sql.DB
}

func InitializeCulturalActivityRepository(db *sql.DB) *CulturalActivityRepositoryService {
	return &CulturalActivityRepositoryService{
		db: db,
	}
}

func (c CulturalActivityRepositoryService) Insert(activity domain.CulturalActivity) error {
	res, err := c.db.Exec(`INSERT INTO ENCUESTA_ACTIVIDAD(
                               ID_USUARIO, JUEGOS_PREFERIDOS, PASATIEMPOS, DEPORTE_INTERES, 
                               FRECUENCIA_EJERCICIO, TIPO_TALLERES, EVENTOS_SOCIALES)
    					VALUES (?,?,?,?,?,?,?);`,
		activity.UserID, activity.PreferredGame, activity.Hobby, activity.PreferredSport,
		activity.ExerciseFrequency, activity.WorkshopType, activity.PreferredSocialEvent)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("cannot insert a cultural activity")
	}

	return nil
}

func (c CulturalActivityRepositoryService) Delete(userID int) error {
	res, err := c.db.Exec(`DELETE FROM ENCUESTA_ACTIVIDAD WHERE ID_USUARIO = ?;`, userID)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("cannot delete a cultural activity")
	}

	return nil
}

func (c CulturalActivityRepositoryService) Update(activity domain.CulturalActivity) error {
	res, err := c.db.Exec(`UPDATE ENCUESTA_ACTIVIDAD 
							SET JUEGOS_PREFERIDOS = ?, PASATIEMPOS = ?, DEPORTE_INTERES = ?, FRECUENCIA_EJERCICIO = ?,
                              	TIPO_TALLERES = ?, EVENTOS_SOCIALES = ? WHERE ID_USUARIO = ?;`,
		activity.PreferredGame, activity.Hobby, activity.PreferredSport,
		activity.ExerciseFrequency, activity.WorkshopType, activity.PreferredSocialEvent, activity.UserID)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("cannot delete a cultural activity")
	}

	return nil
}

func (c CulturalActivityRepositoryService) GetAllOrByID(userID int) ([]domain.CulturalActivity, error) {
	var query = `SELECT * FROM ENCUESTA_ACTIVIDAD;`

	if userID > 0 {
		query = fmt.Sprintf(`SELECT * FROM ENCUESTA_ACTIVIDAD WHERE ID_USUARIO = %v;`, userID)
	}

	rows, err := c.db.Query(query)
	if err != nil {
		return []domain.CulturalActivity{}, err
	}

	defer rows.Close()

	activities := []domain.CulturalActivity{}
	for rows.Next() {
		activity := domain.CulturalActivity{}
		if err = rows.Scan(&activity.ID, &activity.UserID, &activity.PreferredGame, &activity.Hobby,
			&activity.PreferredSport, &activity.ExerciseFrequency, &activity.WorkshopType,
			&activity.PreferredSocialEvent); err != nil {
			return []domain.CulturalActivity{}, err
		}

		activities = append(activities, activity)
	}

	return activities, nil
}
