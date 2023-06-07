package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/url-shortener-go/entity"
)

var (
	querySaveLongUrl = `INSERT INTO urls(user_id, ulid, long_url, created_at) VALUES ($1, $2, $3, $4) RETURNING ulid`

	queryCreateShortUrl = `UPDATE urls SET short_url = $1 WHERE ulid = $2`

	queryGetUrlByShortUrl = `SELECT long_url from urls WHERE short_url = $1`

	queryGetUrlsByUserId = `SELECT long_url, short_url, visits, date_trunc('second', created_at) as created_at from urls where user_id = $1 ORDER BY created_at DESC`

	queryGetUrlsByUsername = `
				SELECT u.long_url, u.short_url, u.visits, date_trunc('second', u.created_at) as created_at from urls as u 
				left join users as us on u.user_id = us.id where us.username = $1 ORDER BY u.created_at DESC`

	queryAddVisitCount = `UPDATE urls SET visits = visits + 1 WHERE short_url = $1`
)

type InterfaceUrlRepository interface {
	SaveLongUrl(req *entity.CreateUrlRequest) (string, error)
	CreateShortUrl(shortUrl string, urlId string) error
	GetUrlByShortUrl(shortUrl string) (string, error)
	GetUrlsByUserId(userId int) ([]*entity.Url, error)
	GetUrlsByUsername(username string) ([]*entity.Url, error)
	AddVisitCount(shortUrl string) error
}

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) InterfaceUrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) SaveLongUrl(req *entity.CreateUrlRequest) (string, error) {
	
	stmt, err := r.db.Prepare(querySaveLongUrl)

	if err != nil {
		return "", err
	}

	var id string

	if err := stmt.QueryRow(req.UserID, req.ULID, req.LongURL, req.CreatedAt).Scan(&id); err != nil {
		return "", err
	}


	return id, nil
}

func(r *UrlRepository) CreateShortUrl(shortUrl string, urlId string) error {
	stmt, err := r.db.Prepare(queryCreateShortUrl)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(shortUrl, urlId); err != nil {
		return err
	}

	return nil
}

func (r *UrlRepository) GetUrlByShortUrl(shortUrl string) (string, error) {
	stmt, err := r.db.Prepare(queryGetUrlByShortUrl)

	if err != nil {
		return "", err
	}

	var url string

	if err := stmt.QueryRow(shortUrl).Scan(&url); err != nil {
		return "", err
	}

	return url, nil
}

func (r *UrlRepository) GetUrlsByUserId(userId int) ([]*entity.Url, error) {
	stmt, err := r.db.Prepare(queryGetUrlsByUserId)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var urls []*entity.Url

	for rows.Next() {
		tempUrl := new(entity.Url)
		if err := rows.Scan(&tempUrl.LongURL, &tempUrl.ShortURL, &tempUrl.Visits, &tempUrl.CreatedAt); err != nil {
			return nil, err
		}

		urls = append(urls, tempUrl)
	}

	return urls, nil
}

func (r *UrlRepository) GetUrlsByUsername(username string) ([]*entity.Url, error) {
	stmt, err := r.db.Prepare(queryGetUrlsByUsername)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var urls []*entity.Url

	for rows.Next() {
		tempUrl := new(entity.Url)
		if err := rows.Scan(&tempUrl.LongURL, &tempUrl.ShortURL, &tempUrl.Visits, &tempUrl.CreatedAt); err != nil {
			return nil, err
		}

		urls = append(urls, tempUrl)
	}

	return urls, nil
}

func (r *UrlRepository) AddVisitCount(shortUrl string) error {
	stmt, err := r.db.Prepare(queryAddVisitCount)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(shortUrl); err != nil {
		return err
	}

	return nil
}