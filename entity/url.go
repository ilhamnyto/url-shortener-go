package entity

import "time"

type Url struct {
	ID        int    	`json:"id"`
	UserID    int    	`json:"user_id"`
	LongURL   string 	`json:"long_url"`
	ShortURL  string 	`json:"short_url"`
	CreatedAt time.Time	`json:"created_at"`
}

