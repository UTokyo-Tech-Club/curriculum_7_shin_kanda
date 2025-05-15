package model

// userはユーザー情報を表す構造体
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// IsValidはユーザー情報が有効かどうかを検証
func (u *User) IsValid() bool {
	return u.Name != "" && len(u.Name) <= 50 && u.Age >= 20 && u.Age <= 80
}
