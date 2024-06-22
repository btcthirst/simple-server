package database

import (
	"errors"
	"fmt"
	"math"
	"simple-server/internal/models"
)

type FakeStorage struct {
	articles []models.Article
}

var fakeArticles = []models.Article{
	{
		ID:          1,
		Title:       "test ti",
		Body:        "some lorem ipsum text",
		Description: "a litle desc",
		Author:      "dev",
	},
	{
		ID:          2,
		Title:       "test ti1",
		Body:        "some1 lorem ipsum text",
		Description: "a1 litle desc",
		Author:      "dev",
	},
	{
		ID:          3,
		Title:       "tesi",
		Body:        "some loremipsum text",
		Description: "a litdesc",
		Author:      "dev",
	},
	{
		ID:          4,
		Title:       "ttti1",
		Body:        "soorem ipsum text",
		Description: "alitle desc",
		Author:      "dev",
	},
}

func New() *FakeStorage {
	return &FakeStorage{
		articles: fakeArticles,
	}
}

func (f *FakeStorage) Create(art models.ArticleDTO) error {
	id := int(math.Round(100))
	newArt := models.Article{
		ID:          id,
		Title:       art.Title,
		Body:        art.Body,
		Description: art.Description,
		Author:      art.Author,
	}
	f.articles = append(f.articles, newArt)
	return nil
}

func (f *FakeStorage) Get() ([]models.Article, error) {
	var err error
	if len(f.articles) == 0 {
		err = errors.New("no articles")
	}
	return f.articles, err
}

func (f *FakeStorage) GetBy(id int) (models.Article, error) {
	find := models.Article{}
	for _, a := range f.articles {
		if a.ID == id {
			find = a
		}
	}
	if find.ID == 0 {
		return find, errors.New("not found")
	}
	return find, nil
}

func (f *FakeStorage) Update(art models.Article) (models.Article, error) {
	err := errors.New("not found")
	for i, a := range f.articles {
		if a.ID == art.ID {
			if !(art.Title == "" || art.Title == " ") {
				f.articles[i].Title = art.Title
			}
			if !(art.Body == "" || art.Body == " ") {
				f.articles[i].Body = art.Body
			}
			if !(art.Description == "" || art.Description == " ") {
				f.articles[i].Description = art.Description
			}
			if !(art.Author == "" || art.Author == " ") {
				f.articles[i].Author = art.Author
			}

			err = nil
			art = f.articles[i]
		}
	}
	return art, err
}

func (f *FakeStorage) Delete(id int) error {
	err := errors.New("not found")
	index := 0
	for i, a := range f.articles {
		if a.ID == id {
			index = i
			err = nil
		}
	}
	if err != nil {
		return err
	}
	f.articles = append(f.articles[0:index], f.articles[index+1:]...)
	fmt.Println(f.articles)
	return nil
}
