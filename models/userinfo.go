package models

type UserInfo struct {
	ID 		  int      `gorm:"column:id;       NOT NULL; PRIMARY KEY;" json:"id"`
	Username  string   `gorm:"column:username; NOT NULL;"              json:"username"`
	Password  string   `gorm:"column:password; NOT NULL"               json:"password"`
	Email	  string   `gorm:"column:email;    NOT NULL"               json:"email"`
	//Nickname  string   `gorm:"column:nickname; NOT NULL"               json:"nickname"`
	
}

