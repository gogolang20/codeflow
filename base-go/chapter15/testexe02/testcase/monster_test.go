package monster

import (
	"testing"
)

// 测试序列化
func TestStore(t *testing.T) {

	monster := &Monster{
		Name:  "齐天大圣",
		Age:   800,
		Skill: "筋斗云",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store 错误 希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store 成功")
}

// 测试反序列化
func TestReStore(t *testing.T) {

	// 先创建一个 Monster 实例 不需要指定字段的值
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore 错误 希望为=%v 实际为=%v", true, res)
	}

	// 进一步判断
	if monster.Name != "齐天大圣" {
		t.Fatalf("monster.Store 错误 希望为=%v 实际为=%v", "齐天大圣", monster.Name)
	}

	t.Logf("monster.ReStore 成功")
}
