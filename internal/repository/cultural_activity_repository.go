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

func (c *CulturalActivityRepositoryService) GetCulturalActivitiesReport() ([]domain.CulturalActivitiesReport, error) {
	query := `
	SELECT 'PASATIEMPOS' AS CATEGORIA, 'Baile' AS ACTIVIDAD, DATE_FORMAT(FECHA, '%Y-%m') AS Mes, SUM(CASE WHEN PASATIEMPOS LIKE '%Baile%' THEN 1 ELSE 0 END) AS Conteo
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Tocar algún instrumento', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Tocar algún instrumento%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Pintar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Pintar%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Dibujar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Dibujar%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Hacer ejercicio', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Hacer ejercicio%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Leer', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Leer%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Salir a caminar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Salir a caminar%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Series o películas', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Series o películas%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'PASATIEMPOS', 'Otros', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Otros%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Festivales', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Festivales%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Conciertos', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Conciertos%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Exposiciones de arte', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Exposiciones de arte%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Literatura/poesía', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Literatura/poesía%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Bailes', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Bailes%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Charlas/Conferencias', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Charlas/Conferencias%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Parques recreativos o de diversión', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Parques recreativos o de diversión%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
	UNION ALL
	SELECT 'EVENTOS_SOCIALES', 'Otros', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Otros%' THEN 1 ELSE 0 END)
	FROM ENCUESTA_ACTIVIDAD
	GROUP BY DATE_FORMAT(FECHA, '%Y-%m');
	`

	rows, err := c.db.Query(query)
	if err != nil {
		return []domain.CulturalActivitiesReport{}, err
	}
	defer rows.Close()

	internetReport := []domain.CulturalActivitiesReport{}
	for rows.Next() {
		report := domain.CulturalActivitiesReport{}
		if err = rows.Scan(&report.Category, &report.Activity, &report.Month, &report.Count); err != nil {
			return []domain.CulturalActivitiesReport{}, err
		}
		internetReport = append(internetReport, report)
	}
	return internetReport, nil
}

/*
func (c *CulturalActivityRepositoryService) GetCulturalActivitiesReport(startDate, endDate time.Time) ([]domain.CulturalActivitiesReport, error) {
	query := `
    SELECT 'PASATIEMPOS' AS CATEGORIA, 'Baile' AS ACTIVIDAD, DATE_FORMAT(FECHA, '%Y-%m') AS Mes, SUM(CASE WHEN PASATIEMPOS LIKE '%Baile%' THEN 1 ELSE 0 END) AS Conteo
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Tocar algún instrumento', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Tocar algún instrumento%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Pintar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Pintar%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Dibujar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Dibujar%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Hacer ejercicio', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Hacer ejercicio%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Leer', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Leer%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Salir a caminar', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Salir a caminar%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Series o películas', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Series o películas%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'PASATIEMPOS', 'Otros', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN PASATIEMPOS LIKE '%Otros%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Festivales', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Festivales%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Conciertos', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Conciertos%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Exposiciones de arte', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Exposiciones de arte%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Literatura/poesía', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Literatura/poesía%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Bailes', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Bailes%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Charlas/Conferencias', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Charlas/Conferencias%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Parques recreativos o de diversión', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Parques recreativos o de diversión%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m')
    UNION ALL
    SELECT 'EVENTOS_SOCIALES', 'Otros', DATE_FORMAT(FECHA, '%Y-%m'), SUM(CASE WHEN EVENTOS_SOCIALES LIKE '%Otros%' THEN 1 ELSE 0 END)
    FROM ENCUESTA_ACTIVIDAD
    WHERE FECHA BETWEEN ? AND ?
    GROUP BY DATE_FORMAT(FECHA, '%Y-%m');
    `

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []domain.CulturalActivitiesReport
	for rows.Next() {
		var report domain.CulturalActivitiesReport
		if err := rows.Scan(&report.Category, &report.Activity, &report.Month, &report.Count); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

*/
