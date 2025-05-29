package dao

import (
	"database/sql"

	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/model"
)

// UserDAOはユーザー情報のデータアクセスオブジェクト
type UserDAO struct {
	db *sql.DB
}

// UserDAOInterfaceはUserDAOのインターフェース
type UserDAOInterface interface {
	Create(user *model.User) error
	FindByName(name string) ([]*model.User, error)
}

// NewUserDAOはUserDAOの新しいインスタンスを返す
func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{db: db}
}

// Createは新しいユーザーをデータベースに作成する
func (dao *UserDAO) Create(user *model.User) error {
	tx, err := dao.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users (id, name, age) VALUES (?, ?, ?)", user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}

	return tx.Commit()

}

// FindByNameは名前に一致するユーザーをデータベースから検索する
func (dao *UserDAO) FindByName(name string) ([]*model.User, error) {
	rows, err := dao.db.Query("SELECT id, name, age FROM users WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}
