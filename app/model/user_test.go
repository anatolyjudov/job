package model

import "testing"

func TestName(t *testing.T) {
	u := User{name: "TestFoo"}

	if u.Name() != "TestFoo" {
		t.Error("User.Name() test error")
	}
}

func TestAvatarFromString(t *testing.T) {
	var runes [2]rune

	u := User{}

	u.SetAvatarFromString("FO")

	runes = u.Avatar()
	if runes[0] != rune('F') || runes[1] != rune('O') || len(runes) != 2 {
		t.Error("User.SetAvatarFromString test error")
	}

	u.SetAvatarFromString("ANATOLY")

	runes = u.Avatar()
	if runes[0] != rune('A') || runes[1] != rune('N') || len(runes) != 2 {
		t.Error("User.SetAvatarFromString test error")
	}
}
