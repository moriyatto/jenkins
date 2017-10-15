package calc

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result == 4 {
		t.Log("テスト成功")
	} else {
		t.Error("テストに失敗")
	}
}
