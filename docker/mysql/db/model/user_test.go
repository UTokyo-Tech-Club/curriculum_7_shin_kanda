// model/user_test.go
package model

import "testing"

func TestUser_IsValid(t *testing.T) {
	// テストケースを定義
	testCases := []struct {
		name     string
		user     User
		expected bool
	}{
		{
			name:     "有効なユーザー情報",
			user:     User{Name: "田中太郎", Age: 30},
			expected: true,
		},
		{
			name:     "名前が空",
			user:     User{Name: "", Age: 30},
			expected: false,
		},
		{
			name:     "名前が長すぎる（51文字）",
			user:     User{Name: "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんあいうえおかきくけこ", Age: 30},
			expected: false,
		},
		{
			name:     "年齢が下限未満（19歳）",
			user:     User{Name: "田中太郎", Age: 19},
			expected: false,
		},
		{
			name:     "年齢が下限（20歳）",
			user:     User{Name: "田中太郎", Age: 20},
			expected: true,
		},
		{
			name:     "年齢が上限（80歳）",
			user:     User{Name: "田中太郎", Age: 80},
			expected: true,
		},
		{
			name:     "年齢が上限超え（81歳）",
			user:     User{Name: "田中太郎", Age: 81},
			expected: false,
		},
	}

	// 各テストケースを実行
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.user.IsValid()
			if actual != tc.expected {
				t.Errorf("%s: expected %v, but got %v", tc.name, tc.expected, actual)
			}
		})
	}
}