package services

import (
	"crypto/rand"
	"database/sql"
	"time"

	"github.com/ilhamnyto/url-shortener-go/entity"
	"github.com/ilhamnyto/url-shortener-go/repositories"
	"github.com/ilhamnyto/url-shortener-go/utils"
	"github.com/oklog/ulid"
)

type InterfaceUrlService interface {
	CreateUrl(req *entity.CreateUrlRequest) *entity.CustomError
	GetUrlByShortUrl(shortUrl string) (string, *entity.CustomError)
	GetUrlsByUserId(userId int) ([]*entity.UserUrlResponse, *entity.CustomError)
	GetUrlsByUsername(username string) ([]*entity.UserUrlResponse, *entity.CustomError)
}

type UrlService struct {
	repo repositories.InterfaceUrlRepository
}

func NewUrlServices(repo repositories.InterfaceUrlRepository) InterfaceUrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) CreateUrl(req *entity.CreateUrlRequest) *entity.CustomError {
	entropy := ulid.Monotonic(rand.Reader, 0)
	now := time.Now()
	id := ulid.MustNew(ulid.Timestamp(now), entropy)

	
	if !utils.IsValidURL(req.LongURL) {
		return entity.BadRequestError("Invalid URL.")
	}

	req.CreatedAt = time.Now()
	req.ULID = id.String()

	urlId, err := s.repo.SaveLongUrl(req)

	if err != nil {
		return entity.RepositoryError(err.Error())
	}

	shortUrl := utils.EncodeID(urlId)

	if err := s.repo.CreateShortUrl(shortUrl, urlId); err != nil {
		return entity.RepositoryError(err.Error())
	}

	return nil
}

func (s *UrlService) GetUrlByShortUrl(shortUrl string) (string, *entity.CustomError) {
	longUrl, err := s.repo.GetUrlByShortUrl(shortUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", entity.NotFoundError("URL not found.")
		}

		return "", entity.RepositoryError(err.Error())
	}

	if err := s.repo.AddVisitCount(shortUrl); err != nil {
		return "", entity.RepositoryError(err.Error())
	}

	return longUrl, nil
}

func (s *UrlService) GetUrlsByUserId(userId int) ([]*entity.UserUrlResponse, *entity.CustomError) {
	urls, err := s.repo.GetUrlsByUserId(userId)

	if err != nil {
		return nil, entity.RepositoryError(err.Error())
	}

	var urlResp []*entity.UserUrlResponse

	for _, url := range(urls) {
		tempUrl := new(entity.UserUrlResponse)
		tempUrl.ParseEntityToResponse(url)
		urlResp = append(urlResp, tempUrl)
	}

	return urlResp, nil
}

func (s *UrlService) GetUrlsByUsername(username string) ([]*entity.UserUrlResponse, *entity.CustomError) {
	urls, err := s.repo.GetUrlsByUsername(username)

	if err != nil {
		return nil, entity.RepositoryError(err.Error())
	}

	var urlResp []*entity.UserUrlResponse

	for _, url := range(urls) {
		tempUrl := new(entity.UserUrlResponse)
		tempUrl.ParseEntityToResponse(url)
		urlResp = append(urlResp, tempUrl)
	}

	return urlResp, nil
}
