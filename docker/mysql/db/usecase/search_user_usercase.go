package usecase

import (
	"errors"

	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/dao"
	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/model"
)

// SearchUserUsecaseはユーザー情報の検索を行うユースケース
type SearchUserUsecase struct {
	userDAO *dao.UserDAO
}

// NewSearchUserUsecaseはSearchUserUsecaseの新しいインスタンスを返す
func NewSearchUserUsecase(userDAO *dao.UserDAO) *SearchUserUsecase {
	return &SearchUserUsecase{userDAO: userDAO}
}

// Executeはユーザー情報の検索を行う
func (u *SearchUserUsecase) Execute(name string) ([]*model.User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return u.userDAO.FindByName(name)
}
