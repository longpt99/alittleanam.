package repositories

import (
	"ala-coffee-notification/models"
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type psqlStorage struct {
	db  *pgxpool.Pool
	ctx context.Context
}

// func NewStorage(pg *models.PostgresConfig) *psqlStorage {
// 	return &psqlStorage{
// 		db:  pg.DB,
// 		ctx: pg.Ctx,
// 	}
// }

type RepositoryConfig interface {
	InitRepositories(*models.PostgresConfig) *Repository
}

type Repository struct {
	UserRepo UserRepository
}

func InitRepositories(store *models.PostgresConfig) *Repository {
	log.Println("Init Repositories Successfully! 🚀")

	return &Repository{
		UserRepo: InitUserRepository(store),
	}
}

func scanStruct(rows pgx.Row, dest interface{}) error {
	val := reflect.ValueOf(dest)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("destination should be a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return errors.New("destination should be a pointer to a struct")
	}

	// Prepare a slice of pointers to struct fields.
	fields := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Field(i).Addr().Interface()
	}

	// Scan the row into the field pointers.
	return rows.Scan(fields...)
}
