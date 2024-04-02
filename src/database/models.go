package database

type Todo struct {
	ID      uint
	Content string 	`gorm:"not null"`
	Done    bool 	`gorm:"not null;default:false"`
}
