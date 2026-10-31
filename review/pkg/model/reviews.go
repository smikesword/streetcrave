package model

import(
	"time"
)

type Reviews struct{
	Id int `json:"id"`
	UserId int `json:"user_id"`
	BusinessId int `json:"business_id"`
	Rating float32 `json:"rating"`
	Comment string `json:"comment"`
	PostedAt time.Time `json:"posted_at"` 
}


