package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type ServicesRepositoryService struct {
	db *sql.DB
}

func InitializeServicesRepository(db *sql.DB) *ServicesRepositoryService {
	return &ServicesRepositoryService{
		db: db,
	}
}

func (s ServicesRepositoryService) Insert(service domain.Services) error {
	res, err := s.db.Exec(`
	INSERT INTO ENCUESTA_SERVICIO(ID_USUARIO, PROVEEDOR_LUZ, PROVEEDOR_AGUA, PROVEEDOR_INTERNET, 
	                              PROVEEDOR_TELEFONO, PROVEEDOR_TELEVISION, VENCIMIENTO_PAGOS, 
	                              PAGOS_ADICIONALES, GASTOS_SERVICIOS)
        VALUES(?,?,?,?,?,?,?,?,?);
	`, service.UserID, service.EnergyProvider, service.WaterProvider, service.InternetProvider,
		service.PhoneProvider, service.TvProvider, service.PaymentDueDate, service.AdditionalPayments,
		service.ServicesBill)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if result == 0 {
		return errors.New("no rows inserted")
	}

	return nil
}

func (s ServicesRepositoryService) Update(service domain.Services) error {
	res, err := s.db.Exec(`
		UPDATE ENCUESTA_SERVICIO SET PROVEEDOR_LUZ = ?, PROVEEDOR_AGUA = ?, PROVEEDOR_INTERNET = ?, PROVEEDOR_TELEFONO = ?,
                             PROVEEDOR_TELEVISION = ?, VENCIMIENTO_PAGOS = ?, PAGOS_ADICIONALES = ?, GASTOS_SERVICIOS = ?
                            WHERE ID_USUARIO = ?;
	`, service.EnergyProvider, service.WaterProvider, service.InternetProvider,
		service.PhoneProvider, service.TvProvider, service.PaymentDueDate, service.AdditionalPayments,
		service.ServicesBill, service.UserID)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if result == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (s ServicesRepositoryService) Delete(userID int) error {
	res, err := s.db.Exec(`
		DELETE FROM ENCUESTA_SERVICIO WHERE ID_USUARIO = ?;
	`, userID)
	if err != nil {
		return err
	}

	result, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if result == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}

func (s ServicesRepositoryService) GetAllOrByID(userID int) ([]domain.Services, error) {
	var query = `SELECT * FROM ENCUESTA_SERVICIO;`

	if userID > 0 {
		query = fmt.Sprintf(`SELECT * FROM ENCUESTA_SERVICIO WHERE ID_USUARIO = %v;`, userID)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.Services{}, err
	}

	defer rows.Close()

	services := []domain.Services{}
	for rows.Next() {
		var service domain.Services
		if err = rows.Scan(&service.ID, &service.UserID, &service.EnergyProvider, &service.WaterProvider,
			&service.InternetProvider, &service.PhoneProvider, &service.TvProvider, &service.PaymentDueDate,
			&service.AdditionalPayments, &service.ServicesBill); err != nil {
			return []domain.Services{}, err
		}

		services = append(services, service)
	}

	return services, nil
}
