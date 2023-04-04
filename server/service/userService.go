package service

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ramonfsk/ibank-backend/server/domain"
	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
)

//go:generate mockgen -destination=../mocks/service/mockUserService.go -package=service github.com/ramonfsk/ibank-backend/server/service UserService
type UserService interface {
	GetAllUsers(status string) ([]dto.UserResponse, *errs.AppError)
	GetUser(id string) (*dto.UserResponse, *errs.AppError)
	NewUser(request dto.UserRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

var rangeNumbers = []rune("0123456789")

func (s DefaultUserService) GetAllUsers(status string) ([]dto.UserResponse, *errs.AppError) {
	users, err := s.repo.FindAll(status)

	if len(users) == 0 {
		return nil, errs.NewValidationError("No have users for this bank on database")
	}

	response := make([]dto.UserResponse, 0)
	for _, user := range users {
		response = append(response, user.ToDTO())
	}

	return response, err
}

func (s DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.AppError) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := user.ToDTO()

	return &response, nil
}

func (s DefaultUserService) NewUser(user dto.UserRequest) (*dto.NewAccountResponse, *errs.AppError) {

	account, err := s.repo.RegisterNewUser(user, generateNewAccount())
	if err != nil {
		return nil, err
	}

	response := account.ToDTONewAccount()

	return &response, nil
}

func generateNewAccount() domain.Account {
	return domain.Account{
		Agency:     generateAgency(),
		Number:     generateNumberAccount(),
		CheckDigit: generateCheckDigit(),
	}
}

func generateAgency() string {
	rand.Seed(time.Now().UnixNano())
	agency := make([]rune, 4)
	for index := range agency {
		agency[index] = rangeNumbers[rand.Intn(len(rangeNumbers))]
	}

	return string(agency)
}

func generateNumberAccount() string {
	rand.Seed(time.Now().UnixNano())
	account := make([]rune, 8)
	for index := range account {
		account[index] = rangeNumbers[rand.Intn(len(rangeNumbers))]
	}

	return string(account)
}

func generateCheckDigit() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9))
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repository}
}
