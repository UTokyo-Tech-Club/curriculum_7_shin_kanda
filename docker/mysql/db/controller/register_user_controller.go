package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// RegisterUserRequest はユーザー登録リクエスト
type RegisterUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// RegisterUserUsecaseInterfaceはユーザー登録ユースケースのインターフェース
// (mockgen用コメントは任意)
//
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=controller
type RegisterUserUsecaseInterface interface {
	Execute(name string, age int) (string, error)
}

// RegisterUserController はユーザー登録のHTTPリクエストを処理する
type RegisterUserController struct {
	registerUsecase RegisterUserUsecaseInterface
}

// NewRegisterUserController は新しいRegisterUserControllerインスタンスを作成する
func NewRegisterUserController(registerUsecase RegisterUserUsecaseInterface) *RegisterUserController {
	return &RegisterUserController{registerUsecase: registerUsecase}
}

// Handle はPOSTリクエストを処理する
func (ctrl *RegisterUserController) Handle(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json decode, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := ctrl.registerUsecase.Execute(req.Name, req.Age)
	if err != nil {
		log.Printf("fail: registerUsecase.Execute, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
