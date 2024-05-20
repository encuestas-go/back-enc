package repository

import (
	"database/sql"

	"github.com/encuestas-go/back-enc/internal/domain"
)

type MapRepositoryService struct {
	db *sql.DB
}

func InitializeMapRepository(db *sql.DB) *MapRepositoryService {
	return &MapRepositoryService{
		db: db,
	}
}

func (m *MapRepositoryService) GetAll() ([]domain.UserAddressMap, error) {
	query := `
	SELECT ID_USUARIO, LONGITUD, LATITUD, DIRECCION_RESIDENCIA FROM ENCUESTA_NIVEL_SOCIOECONOMICO;
	`
	queryRows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer queryRows.Close()

	var usersAddress []domain.UserAddressMap
	for queryRows.Next() {
		var info domain.UserAddressMap
		if err = queryRows.Scan(&info.IDUser, &info.CompleteAddress, &info.LongitudeAddress, &info.LatitudeAddress); err != nil {
			return nil, err
		}
		usersAddress = append(usersAddress, info)
	}
	return usersAddress, nil
}
