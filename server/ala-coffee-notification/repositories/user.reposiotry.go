package repositories

import (
	"ala-coffee-notification/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// type UserRepository struct {

// }

type UserRepository interface {
	InsertOne(body *models.SignUpAccount) (string, error)
	FindUserByEmail(email string) (models.User, error)
	FindUserByConditions(credential string) (models.User, error)
	DeleteUserInactive(id string) error
	// List() (*[]Product, error)
	// DetailByID(id string) (*Product, error)
	// Delete(id string) error
}

type repo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func InitUserRepository(store *models.PostgresConfig) UserRepository {
	return &repo{
		db:  store.DB,
		ctx: store.Ctx,
	}
}

func (s *repo) InsertOne(body *models.SignUpAccount) (string, error) {
	query := `INSERT INTO users (first_name, last_name, email, password, date_of_birth) 
		VALUES (@firstName, @lastName, @email, @password, @dateOfBirth) 
		RETURNING id
	`
	args := pgx.NamedArgs{
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"email":       body.Email,
		"password":    body.Password,
		"dateOfBirth": body.DateOfBirth,
	}

	var id string
	err := s.db.QueryRow(s.ctx, query, args).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *repo) FindUserByEmail(email string) (models.User, error) {
	row, _ := s.db.Query(s.ctx, `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL LIMIT 1`, email)
	acc, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.User])

	if err != nil {
		return acc, err
	}

	return acc, nil
}

func (s *repo) DeleteUserInactive(id string) error {
	_, err := s.db.Exec(s.ctx, `UPDATE users SET deleted_at = NOW() where id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *repo) FindUserByConditions(credential string) (models.User, error) {
	row, _ := s.db.Query(s.ctx, `SELECT * FROM users WHERE (email = $1 OR username = $1) AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 1`, credential)
	acc, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.User])

	if err != nil {
		return acc, err
	}

	return acc, nil
}
