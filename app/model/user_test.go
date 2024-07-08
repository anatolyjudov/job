package model

import "testing"

func TestName(t *testing.T) {
	u := User{name: "TestFoo"}

	if u.Name() != "TestFoo" {
		t.Error("User.Name() test error")
	}
}

func TestAvatarFromString(t *testing.T) {
	u := User{}

	u.SetAvatarFromString("FO")

	runes := u.Avatar()
	if runes[0] != rune('F') || runes[1] != rune('O') {
		t.Error("User.SetAvatarFromString test error")
	}
}
