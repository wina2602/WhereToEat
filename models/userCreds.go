package models

type Usercreds struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"uname" gorm:"notNull"`
	Salt     string `json:"salt" gorm:"notNull"`
	Password string `json:"password" gorm:"notNull"`
}
