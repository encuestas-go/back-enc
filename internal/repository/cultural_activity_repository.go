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
			&activity.PreferredSocialEvent, &activity.Date); err != nil {
			return []domain.CulturalActivity{}, err
		}

		activities = append(activities, activity)
	}

	return activities, nil
}

func (c *CulturalActivityRepositoryService) GetCulturalActivitiesReport(startDate, endDate string) ([]domain.CulturalActivitiesReport, error) {
	query := `
	SELECT
		SUM(CASE WHEN PASATIEMPOS LIKE '%Baile%' THEN 1 ELSE 0 END) AS Baile,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Tocar algún instrumento%' THEN 1 ELSE 0 END) AS Tocar_algun_instrumento,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Pintar%' THEN 1 ELSE 0 END) AS Pintar,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Dibujar%' THEN 1 ELSE 0 END) AS Dibujar,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Hacer ejercicio%' THEN 1 ELSE 0 END) AS Hacer_ejercicio,
		SUM(CASE WHEN PASATIEMPOS LIKE '%Leer%' THEN 1 ELSE 0 END) AS Leer,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Salir a caminar%' THEN 1 ELSE 0 END) AS Salir_a_caminar,
		SUM(CASE WHEN PASATIEMPOS  LIKE '%Series o películas%' THEN 1 ELSE 0 END) AS Series_o_peliculas,
		SUM(CASE WHEN PASATIEMPOS LIKE '%Otros%' THEN 1 ELSE 0 END) AS Otras_actividades,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Festivales%' THEN 1 ELSE 0 END) AS Festivales,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Conciertos%' THEN 1 ELSE 0 END) AS Conciertos,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Exposiciones de arte%' THEN 1 ELSE 0 END) AS Exposiciones_de_arte,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Literatura/poesía%' THEN 1 ELSE 0 END) AS Literatura_poesia,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Bailes%' THEN 1 ELSE 0 END) AS Bailes,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Charlas/Conferencias%' THEN 1 ELSE 0 END) AS Charlas_conferencias,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Parques recreativos o de diversión%' THEN 1 ELSE 0 END) AS Parques_recreativos_diversion,
		SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Otros%' THEN 1 ELSE 0 END) AS Otros_eventos
	FROM ENCUESTA_ACTIVIDAD
	WHERE FECHA BETWEEN ? AND ?;
	`
	rows, err := c.db.Query(query, startDate, endDate)
	if err != nil {
		return []domain.CulturalActivitiesReport{}, err
	}
	defer rows.Close()

	internetReport := []domain.CulturalActivitiesReport{}
	for rows.Next() {
		report := domain.CulturalActivitiesReport{}
		if err = rows.Scan(&report.Dance, &report.PlayInstrument, &report.Paint, &report.Draw, &report.DoExercise,
			&report.Read, &report.GoWalking, &report.Movies, &report.OtherActivities, &report.Festivals,
			&report.Concerts, &report.ArtExposition, &report.LiteraturePoetry, &report.Dances, &report.Conferences,
			&report.RecreationalParks, &report.OtherEvents); err != nil {
			return []domain.CulturalActivitiesReport{}, err
		}
		internetReport = append(internetReport, report)
	}
	return internetReport, nil
}
