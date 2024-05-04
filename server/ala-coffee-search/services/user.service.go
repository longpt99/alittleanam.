package services

import (
	"ala-coffee-search/models"
	"ala-coffee-search/repositories"
	"ala-coffee-search/utils"
	"ala-coffee-search/utils/errs"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

func (s *UserService) CreateAccount(body *models.SignUpAccount) (interface{}, error) {
	hashedPw, _ := utils.HashPassword(body.Password)
	body.Password = hashedPw

	id, err := s.UserRepo.InsertOne(body)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == errs.UniqueViolationErrCode {
				user, _ := s.UserRepo.FindUserByEmail(body.Email)
				if user.Status == models.UserStatusInactive {
					err = s.UserRepo.DeleteUserInactive(user.ID)
					if err != nil {
						return "", err
					}

					id, err = s.UserRepo.InsertOne(body)
					if err != nil {
						return "", err
					}

					return id, nil
				}
			}

			return "", errs.E("account has exists", http.StatusBadRequest)
		}
	}

	return id, nil
}

func (s *UserService) SignInAccount(body *models.SignInAccount) (interface{}, error) {
	user, err := s.UserRepo.FindUserByConditions(body.Credential)

	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errs.E("account un-exists", http.StatusBadRequest)
		}

		return nil, err
	}

	if user.Status == models.UserStatusInactive {
		return nil, errs.E("your account is unverified", http.StatusBadRequest)
	}

	err = utils.ComparePassword(body.Password, user.Password)
	if err != nil {
		return nil, errs.E("wrong password", http.StatusBadRequest)
	}

	return user, nil
}
