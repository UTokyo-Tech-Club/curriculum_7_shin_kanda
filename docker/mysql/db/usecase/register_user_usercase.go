package usecase

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/dao"
	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/model"
)

// RegisterUserUsecaseはユーザー情報の登録を行うユースケース
type RegisterUserUsecase struct {
	userDAO dao.UserDAOInterface
}

// NewRegisterUserUsecaseはRegisterUserUsecaseの新しいインスタンスを返す
func NewRegisterUserUsecase(userDAO dao.UserDAOInterface) *RegisterUserUsecase {
	return &RegisterUserUsecase{userDAO: userDAO}
}

// Executeはユーザー情報の登録を行う
func (u *RegisterUserUsecase) Execute(name string, age int) (string, error) {
	user := &model.User{
		Name: name,
		Age:  age,
	}

	if !user.IsValid() {
		return "", errors.New("invalid user data")
	}

	// IDを生成
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id, _ := ulid.New(ulid.Timestamp(t), entropy)
	user.ID = id.String()
	if err := u.userDAO.Create(user); err != nil {
		return "", err
	}

	return user.ID, nil
}
