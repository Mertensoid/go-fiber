package users

import (
	"context"
	"go-fiber/pkg/cryptograf"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
	logger *zerolog.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger) *UserRepository {
	r := &UserRepository{
		dbpool: dbpool,
		logger: logger,
	}
	return r
}

// Получение количества всех пользователей из базы данных
func (r *UserRepository) CountAll() int {
	query := `SELECT count(*) FROM users`
	var count int
	r.dbpool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

// Регистрация нового пользователя
func (r *UserRepository) addUser(form RegistrationForm) error {
	passwordHash, err := cryptograf.HashPassword(form.Password)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}
	query := `INSERT INTO users (email, password, name, registered) 
				VALUES (@email, @password, @name, @registered)
				`
	args := pgx.NamedArgs{
		"email":      form.Email,
		"password":   passwordHash,
		"name":       form.Name,
		"registered": time.Now(),
	}
	_, err = r.dbpool.Exec(context.Background(), query, args)
	return err
}
