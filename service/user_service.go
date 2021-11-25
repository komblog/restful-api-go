package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/komblog/restful-api-go/helper"
	"github.com/komblog/restful-api-go/model/domain"
	"github.com/komblog/restful-api-go/model/web"
	"github.com/komblog/restful-api-go/repository"
)

type UserService interface {
	Create(ctx context.Context, user web.CreateUserRequest) web.UserDTO
	FindById(ctx context.Context, userId int) web.UserDTO
	FindAll(ctx context.Context) []web.UserDTO
	Update(ctx context.Context, user web.CreateUserRequest) web.UserDTO
	Delete(ctx context.Context, userId int)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		validator:      validator,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, user web.CreateUserRequest) web.UserDTO {
	errValidation := service.validator.Struct(user)
	helper.PanicIfError(errValidation)

	tx, errBegin := service.DB.Begin()
	helper.PanicIfError(errBegin)

	defer func() {
		errRecover := recover()
		if errRecover != nil {
			errRollback := tx.Rollback()
			helper.PanicIfError(errRollback)
			panic(errRecover)
		} else {
			errorCommit := tx.Commit()
			panic(errorCommit.Error())
		}
	}()

	userResult := domain.User{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email.NullString,
		Birth_Date: user.Birth_Date.NullTime,
	}

	userResult = service.UserRepository.Save(ctx, tx, userResult)

	return helper.ToUserDTO(userResult)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserDTO {
	tx, errBegin := service.DB.Begin()
	helper.PanicIfError(errBegin)
	defer helper.RollbackOrCommit(tx)

	userResult, err := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserDTO(userResult)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserDTO {
	tx, errBegin := service.DB.Begin()
	helper.PanicIfError(errBegin)
	defer helper.RollbackOrCommit(tx)

	usersResult := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUsersDTO(usersResult)
}

func (service *UserServiceImpl) Update(ctx context.Context, user web.CreateUserRequest) web.UserDTO {
	errValidation := service.validator.Struct(user)
	helper.PanicIfError(errValidation)

	tx, errBegin := service.DB.Begin()
	helper.PanicIfError(errBegin)
	defer helper.RollbackOrCommit(tx)

	userResult, _ := service.UserRepository.FindById(ctx, tx, user.Id)

	userResult.FirstName = user.FirstName
	userResult.LastName = user.LastName
	userResult.Email = user.Email.NullString
	userResult.Birth_Date = user.Birth_Date.NullTime

	userUpdate := service.UserRepository.Update(ctx, tx, userResult)
	return helper.ToUserDTO(userUpdate)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, errBegin := service.DB.Begin()
	helper.PanicIfError(errBegin)
	defer helper.RollbackOrCommit(tx)

	userResult, errFind := service.UserRepository.FindById(ctx, tx, userId)
	helper.PanicIfError(errFind)

	service.UserRepository.Delete(ctx, tx, userResult.Id)
}
