package service

import (
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	newsViewModel "puppy/ViewModel/news"
	"puppy/model/news"
	"puppy/repository"
	"time"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
	GetNewsById(id string) (news.News, error)
	CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error)
	IsNewsExist(id string) bool
	EditNews(userInput newsViewModel.EditNewsViewModel, imageFile *multipart.FileHeader) error
	DeleteNews(id string) error
	AddVisitCount(id string) error
	AddLike(id string) error
}

type newsService struct {
}

func NewNewsService() NewsService {
	return newsService{}
}

func (newsService) GetNewsList() ([]news.News, error) {

	newsRepository := repository.NewNewsRepository()
	newsList, err := newsRepository.GetNewsList()
	return newsList, err
}

func (s newsService) GetNewsById(id string) (news.News, error) {
	newsRepository := repository.NewNewsRepository()
	news, err := newsRepository.GetNewsById(id)
	return news, err
}

func (s newsService) CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error) {

	newsEntity := news.News{
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return "", err
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return "", err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return "", err
		}
		newsEntity.ImageName = fileName
	}

	newsRepository := repository.NewNewsRepository()
	newsId, err := newsRepository.InsertNews(newsEntity)

	return newsId, err
}

func (s newsService) IsNewsExist(id string) bool {
	newsRepository := repository.NewNewsRepository()
	_, err := newsRepository.GetNewsById(id)

	if err != nil {
		return false
	}

	return true
}

func (s newsService) EditNews(userInput newsViewModel.EditNewsViewModel, imageFile *multipart.FileHeader) error {

	newsRepository := repository.NewNewsRepository()
	newsEntity := news.News{
		Id:               userInput.Id,
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return err
		}
		oldNews, err := newsRepository.GetNewsById(userInput.Id)
		if err != nil {
			return err
		}

		wd, err := os.Getwd()

		if oldNews.ImageName != "" {
			oldImageServerPath := filepath.Join(wd, "wwwroot", "images", "news", oldNews.ImageName)
			os.Remove(oldImageServerPath)
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		imageServerPath := filepath.Join(wd, "wwwroot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return err
		}
		newsEntity.ImageName = fileName
	}

	err := newsRepository.UpdateNewsById(newsEntity)

	return err
}
func (s newsService) DeleteNews(id string) error {

	newsRepository := repository.NewNewsRepository()
	oldNews, err := newsRepository.GetNewsById(id)
	if err != nil {
		return err
	}
	if oldNews.ImageName != "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		oldImageServerPath := filepath.Join(wd, "wwwroot", "images", "news", oldNews.ImageName)
		os.Remove(oldImageServerPath)
	}

	err = newsRepository.DeleteNewsById(id)

	return err
}

func (s newsService) AddVisitCount(id string) error {

	newsRepository := repository.NewNewsRepository()
	news, err := newsRepository.GetNewsById(id)
	if err != nil {
		return err
	}
	news.VisitCount += 1
	err = newsRepository.UpdateNewsById(news)
	if err != nil {
		return err
	}

	return nil
}
func (s newsService) AddLike(id string) error {

	newsRepository := repository.NewNewsRepository()
	news, err := newsRepository.GetNewsById(id)
	if err != nil {
		return err
	}
	news.LikeCount += 1
	err = newsRepository.UpdateNewsById(news)
	if err != nil {
		return err
	}

	return nil
}
