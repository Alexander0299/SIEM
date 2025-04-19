package model

type Inter interface {
	Inters() string
	IntersId() int
}

type User struct {
	Login string `json:"login" csv:"Login"`
	ID    int    `json:"id"`
}

func (u User) Inters() string {
	return u.Login

}
func (u User) IntersId() int {
	return u.ID

}

type Alert struct {
	Massage string `json:"massage" csv:"Massage"`
	ID      int    `json:"id" csv:"ID"`
}

func (a Alert) Inters() string {
	return a.Massage

}
func (a Alert) IntersId() int {
	return a.ID

}

type Log struct {
	Area string `json:"area" csv:"Area"`
	ID   int    `json:"id" csv:"ID"`
}

func (l Log) Inters() string {
	return l.Area

}
func (l Log) IntersId() int {
	return l.ID

}
