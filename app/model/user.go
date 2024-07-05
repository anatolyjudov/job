package model

type User struct {
	id     int
	name   string
	avatar [2]rune
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Avatar() [2]rune {
	return u.avatar
}

func (u *User) AvatarAsString() string {
	return string(u.avatar[0:2])
}

func (u *User) SetId(id int) *User {
	u.id = id
	return u
}

func (u *User) SetName(name string) *User {
	u.name = name
	return u
}

func (u *User) SetAvatar(avatar [2]rune) *User {
	u.avatar = avatar
	return u
}

func (u *User) SetAvatarFromString(avatar string) *User {
	runedAvatar := []rune(avatar)
	copy(u.avatar[:], runedAvatar[:min(len(runedAvatar), len(u.avatar))])
	return u
}
