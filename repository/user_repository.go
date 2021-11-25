package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/komblog/restful-api-go/helper"
	"github.com/komblog/restful-api-go/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId int)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	queryInsert := "INSERT INTO USER(FIRSTNAME, LASTNAME, EMAIL, BIRTH_DATE) VALUES(?,?,?,?)"
	result, errInsert := tx.ExecContext(ctx, queryInsert, user.FirstName, user.LastName, user.Email.String, user.Birth_Date.Time)
	helper.PanicIfError(errInsert)

	id, errId := result.LastInsertId()
	if errId != nil {
		panic(errId.Error())
	}

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	querySelect := "SELECT ID, FIRSTNAME, LASTNAME, EMAIL, BIRTH_DATE FROM USER WHERE ID = ?"
	result, errSelect := tx.QueryContext(ctx, querySelect, userId)
	helper.PanicIfError(errSelect)
	defer result.Close()

	user := domain.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Birth_Date)
		if err != nil {
			panic(err.Error())
		}
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	querySelect := "SELECT ID, FIRSTNAME, LASTNAME, EMAIL, BIRTH_DATE FROM USER"
	result, errSelect := tx.QueryContext(ctx, querySelect)
	helper.PanicIfError(errSelect)
	defer result.Close()

	users := []domain.User{}

	for result.Next() {
		user := domain.User{}
		err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Birth_Date)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	queryUpdate := "UPDATE USER SET FIRSTNAME = ?, LASTNAME = ?, EMAIL = ?, BIRTH_DATE = ? WHERE ID = ?"
	_, errUpdate := tx.ExecContext(ctx, queryUpdate, user.FirstName, user.LastName, user.Email.String, user.Birth_Date.Time)
	helper.PanicIfError(errUpdate)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	queryDelete := "DELETE FROM USER WHERE ID = ?"
	_, errDelete := tx.ExecContext(ctx, queryDelete, userId)
	helper.PanicIfError(errDelete)
}
