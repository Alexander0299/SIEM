package model

type Inter interface {
	Inters() string
}

type User struct {
	Login string
}

func (u User) Inters() string {
	return u.Login
}

type Alert struct {
	Massage string
}

func (a Alert) Inters() string {
	return a.Massage
}

type Log struct {
	Area string
}

func (l Log) Inters() string {
	return l.Area
}
