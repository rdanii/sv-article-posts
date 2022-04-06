package structs

import "time"

type Posts struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Title        string    `json:"title" type:"varchar(200)" binding:"required"`
	Content      string    `json:"content" type:"text" binding:"required"`
	Category     string    `json:"category" type:"varchar(100)" binding:"required"`
	Created_date time.Time `json:"created_date" type:"timestamp"`
	Updated_date time.Time `json:"updated_date" type:"timestamp"`
	Status       string    `json:"status" type:"varchar(100)" binding:"required"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
