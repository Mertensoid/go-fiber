package vacancy

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (role, company, sphere, salary, location, email) VALUES (@role, @company, @sphere, @salary, @location, @email)`
	args := pgx.NamedArgs{
		"role":     form.Role,
		"company":  form.Company,
		"sphere":   form.Sphere,
		"salary":   form.Salary,
		"location": form.Location,
		"email":    form.Email,
	}
	_, err := r.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать вакансию: %w", err)
	}
	return nil
}
