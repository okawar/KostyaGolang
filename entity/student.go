package entity

type Student struct {
	Sid    uint64 `json:"sid" gorm:"primarykey"`
	Login  string
	Pass   string
	Name   string
	Access bool
}
