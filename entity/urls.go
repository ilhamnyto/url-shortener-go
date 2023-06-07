package entity

import (
	"os"
	"time"
)

type Url struct {
	ID        	int    		`json:"id"`
	UserID    	int    		`json:"user_id"`
	LongURL   	string 		`json:"long_url"`
	ShortURL  	string 		`json:"short_url"`
	Visits		int			`json:"visits"`
	CreatedAt 	time.Time	`json:"created_at"`
}

type CreateUrlRequest struct {
	LongURL   	string 		`json:"long_url"`
	UserID    	int    		`json:"user_id"`
	ULID 		string		`json:"ulid"`
	CreatedAt 	time.Time	`json:"created_at"`
}

type UserUrlResponse struct {
	LongURL		string		`json:"long_url"`
	ShortURL  	string 		`json:"short_url"`
	Visits		int			`json:"visits"`
	CreatedAt 	time.Time	`json:"created_at"`
}

func (u *UserUrlResponse) ParseEntityToResponse(url *Url) {
	u.LongURL = url.LongURL
	u.ShortURL = os.Getenv("BASE_URL") + url.ShortURL
	u.Visits = url.Visits
	u.CreatedAt = url.CreatedAt
}
