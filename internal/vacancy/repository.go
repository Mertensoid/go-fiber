package vacancy

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	dbpool *pgxpool.Pool
	logger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger) *VacancyRepository {
	r := &VacancyRepository{
		dbpool: dbpool,
		logger: logger,
	}
	return r
}

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) {

}
