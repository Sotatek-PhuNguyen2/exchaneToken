package models

type NetWork struct {
	Id            int    `json:"networkID" gorm:"type: int; not null;primaryKey;autoIncrement"`
	NameOfNetwork string `json:"nameOfNetwork" gorm:"type: varchar(100); not null"`
	Dec           int    `json:"dec" gorm:"type: int; not null"`
}

// type User struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID string
// 	Company   Company `gorm:"references:CompanyID"` // use Company.CompanyID as references
// }

// type Company struct {
// 	CompanyID int
// 	Code      string
// 	Name      string
// }
