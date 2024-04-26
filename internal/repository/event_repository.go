package repository

import (
	"database/sql"
	"errors"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type EventRepositoryService struct {
	db *sql.DB
}

func InitializeEventRepository(db *sql.DB) *EventRepositoryService {
	return &EventRepositoryService{
		db: db,
	}
}

func (e EventRepositoryService) CreateEvent(event domain.Event) error {
	res, err := e.db.Exec(`
	INSERT INTO PUBLICACION_EVENTO(NOMBRE_EVENTO, LUGAR, FECHA, HORA, UBICACION, DESCRIPCION_EVENTO, CATEGORIA) 
			VALUES(?,?,?,?,?,?,?);
	`, event.EventName, event.Place, event.Date, event.Hour, event.Location, event.Description, event.Category)

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

func (e EventRepositoryService) Update(event domain.Event) error {
	res, err := e.db.Exec(`
		UPDATE PUBLICACION_EVENTO SET 
		            NOMBRE_EVENTO = ?, LUGAR = ?, FECHA = ?, HORA = ?, 
		            UBICACION = ?, DESCRIPCION_EVENTO = ?, CATEGORIA = ? 
		            WHERE ID_USUARIO = ?;
	`, event.EventName, event.Place, event.Date, event.Hour, event.Location, event.Description, event.Category, event.IDUser)

	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("cannot update a cultural activity")
	}

	return nil
}

func (e EventRepositoryService) GetEvents() ([]domain.Event, error) {
	rows, err := e.db.Query(`SELECT * FROM PUBLICACION_EVENTO;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []domain.Event

	for rows.Next() {
		var event domain.Event
		if err = rows.Scan(&event.ID, event.EventName, &event.Place, &event.Date, &event.Hour, &event.Location,
			&event.Description, &event.Category, &event.IDUser); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (e EventRepositoryService) DeleteEvent(userID int) error {
	res, err := e.db.Exec(`DELETE FROM PUBLICACION_EVENTO WHERE ID_USUARIO = ?;`, userID)
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
