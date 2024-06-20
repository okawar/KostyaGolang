package entity

type Vds struct {
	Vid           uint64 `json:"uid" gorm:"primarykey"`
	Uid           uint64 // список пользователей
	Name          string //название курса
	Description   string // описание курса
	Price         uint64
	Course_Holder string // тот кто выдал курс
}
