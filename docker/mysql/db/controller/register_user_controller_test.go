// controller/register_user_controller_test.go
package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// RegisterUserUsecaseのモック
type mockRegisterUserUsecase struct {
	executeFunc func(name string, age int) (string, error)
}

func (m *mockRegisterUserUsecase) Execute(name string, age int) (string, error) {
	return m.executeFunc(name, age)
}

func TestRegisterUserController_Handle(t *testing.T) {
	// テストケースを定義
	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		mockExecuteID  string
		mockExecuteErr error
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "有効なリクエストで登録成功",
			requestBody:    map[string]interface{}{"name": "佐藤健", "age": 35},
			mockExecuteID:  "01F8Z5MRXN5RCCPA0FQEMJA0S",
			mockExecuteErr: nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   map[string]string{"id": "01F8Z5MRXN5RCCPA0FQEMJA0S"},
		},
		{
			name:           "無効なリクエストで登録失敗",
			requestBody:    map[string]interface{}{"name": "", "age": 35},
			mockExecuteID:  "",
			mockExecuteErr: errors.New("invalid user data"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name:           "数値型でない年齢を含むリクエスト",
			requestBody:    map[string]interface{}{"name": "伊藤洋子", "age": "三十歳"},
			mockExecuteID:  "",
			mockExecuteErr: nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	// 各テストケースを実行
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// リクエストボディを作成
			body, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// モックユースケースの設定
			mockUsecase := &mockRegisterUserUsecase{
				executeFunc: func(name string, age int) (string, error) {
					return tc.mockExecuteID, tc.mockExecuteErr
				},
			}

			// テスト対象のコントローラーを作成
			controller := NewRegisterUserController(mockUsecase)

			// HTTPリクエストとレスポンスレコーダーを設定
			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// コントローラーのHandleメソッドを実行
			controller.Handle(rec, req)

			// レスポンスステータスの検証
			if rec.Code != tc.expectedStatus {
				t.Errorf("expected status %d but got %d", tc.expectedStatus, rec.Code)
			}

			// レスポンスボディの検証（成功時のみ）
			if tc.expectedBody != nil {
				var responseBody map[string]string
				if err := json.NewDecoder(rec.Body).Decode(&responseBody); err != nil {
					t.Errorf("failed to decode response body: %v", err)
					return
				}

				if responseBody["id"] != tc.expectedBody["id"] {
					t.Errorf("expected response body %v but got %v", tc.expectedBody, responseBody)
				}
			}
		})
	}
}
