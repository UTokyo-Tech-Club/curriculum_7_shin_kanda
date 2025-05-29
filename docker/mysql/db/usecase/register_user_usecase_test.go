// usecase/register_user_usecase_test.go
package usecase

import (
	"errors"
	"testing"

	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/model"
)

// UserDAOのモック
type mockUserDAO struct {
	createFunc func(user *model.User) error
}

func (m *mockUserDAO) Create(user *model.User) error {
	return m.createFunc(user)
}

func (m *mockUserDAO) FindByName(name string) ([]*model.User, error) {
	// このテストでは使用しないのでダミー実装
	return nil, nil
}

func TestRegisterUserUsecase_Execute(t *testing.T) {
	// テストケースを定義
	testCases := []struct {
		name          string
		userName      string
		userAge       int
		mockCreateErr error
		expectErr     bool
		expectErrMsg  string
	}{
		{
			name:          "有効なユーザー情報で登録成功",
			userName:      "鈴木一郎",
			userAge:       25,
			mockCreateErr: nil,
			expectErr:     false,
		},
		{
			name:          "名前が空で登録失敗",
			userName:      "",
			userAge:       25,
			mockCreateErr: nil,
			expectErr:     true,
			expectErrMsg:  "invalid user data",
		},
		{
			name:          "名前が長すぎる（51文字）で登録失敗",
			userName:      "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんあいうえおかきくけこ",
			userAge:       25,
			mockCreateErr: nil,
			expectErr:     true,
			expectErrMsg:  "invalid user data",
		},
		{
			name:          "年齢が下限未満（19歳）で登録失敗",
			userName:      "佐藤次郎",
			userAge:       19,
			mockCreateErr: nil,
			expectErr:     true,
			expectErrMsg:  "invalid user data",
		},
		{
			name:          "年齢が上限超え（81歳）で登録失敗",
			userName:      "高橋三郎",
			userAge:       81,
			mockCreateErr: nil,
			expectErr:     true,
			expectErrMsg:  "invalid user data",
		},
		{
			name:          "データベースエラーで登録失敗",
			userName:      "山田花子",
			userAge:       30,
			mockCreateErr: errors.New("database error"),
			expectErr:     true,
			expectErrMsg:  "database error",
		},
	}

	// 各テストケースを実行
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// モックDAOの設定
			mockDAO := &mockUserDAO{
				createFunc: func(user *model.User) error {
					return tc.mockCreateErr
				},
			}

			// テスト対象のユースケースを作成
			usecase := NewRegisterUserUsecase(mockDAO)

			// ユースケースを実行
			id, err := usecase.Execute(tc.userName, tc.userAge)

			// エラーチェック
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tc.expectErrMsg != "" && err.Error() != tc.expectErrMsg {
					t.Errorf("expected error message %q, but got %q", tc.expectErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
					return
				}
				if id == "" {
					t.Errorf("expected non-empty ID but got empty string")
				}
			}
		})
	}
}
