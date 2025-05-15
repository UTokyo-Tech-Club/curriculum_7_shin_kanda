package controller

import (
    "encoding/json"
    "log"
    "net/http"
    
    "github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/usecase"
)

// UserResponse はHTTPレスポンス用のユーザー情報
type UserResponse struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// SearchUserController はユーザー検索のHTTPリクエストを処理する
type SearchUserController struct {
    searchUsecase *usecase.SearchUserUsecase
}

// NewSearchUserController は新しいSearchUserControllerインスタンスを作成する
func NewSearchUserController(searchUsecase *usecase.SearchUserUsecase) *SearchUserController {
    return &SearchUserController{searchUsecase: searchUsecase}
}

// Handle はGETリクエストを処理する
func (ctrl *SearchUserController) Handle(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    
    users, err := ctrl.searchUsecase.Execute(name)
    if err != nil {
        log.Printf("fail: searchUsecase.Execute, %v\n", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    // レスポンス用の構造体に変換
    responses := make([]UserResponse, len(users))
    for i, user := range users {
        responses[i] = UserResponse{
            ID:   user.ID,
            Name: user.Name,
            Age:  user.Age,
        }
    }
    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responses); err != nil {
        log.Printf("fail: json.Encode, %v\n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}