package entity

// *Book struct represents book table in database
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"tyep:text" json:"description"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreign_key:userID;constraint:onUpdate:CASECADE,onDelete:CASECADE" json:"user"`
}
