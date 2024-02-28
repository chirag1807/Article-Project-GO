package repository

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddArticle(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Title        string
		Content      string
		Topic        int64
		Author       int64
		Expected     interface{}
	}{
		{
			TestCaseName: "Success",
			Title:        "React JS",
			Content:      "This Article is About React JS.",
			Topic:        943765517224148993,
			Author:       945437407105286145,
			Expected:     nil,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {

			article := request.Article{
				Title:   v.Title,
				Content: v.Content,
				Topic:   v.Topic,
				Author:  v.Author,
			}

			err := NewArticleRepo(conn).AddArticle(article)
			assert.Equal(t, v.Expected, err)
		})
	}
}

func TestGetMyArticles(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Author       int64
		Expected     []response.Article
		StatusCode   int
	}{
		{
			TestCaseName: "Success",
			Author:       945437407105286145,
			Expected: []response.Article{
				{
					ID:      947136363651727361,
					Title:   "React JS",
					Content: "This Article is About React JS.",
					Likes:   0,
					Views:   0,
					Topic:   943765517224148993,
					Author:  945437407105286145,
					// PublishedAt : "2024-02-27T06:40:30.573Z",
				},
			},
			StatusCode: 200,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			response, _ := NewArticleRepo(conn).GetMyArticles(v.Author)
			assert.Equal(t, v.Expected[0].ID, response[0].ID)
		})
	}
}

func TestUpdateArticle(t *testing.T) {
	article := request.Article{
		ID:      947136363651727361,
		Title:   "ReactJS-Web Framework",
		Content: "This article is about ReactJS, A web framework.",
		Topic:   943765517224148993,
	}

	err := NewArticleRepo(conn).UpdateArticle(article)
	assert.Equal(t, nil, err)
}

func TestDeleteArticle(t *testing.T) {
	err := NewArticleRepo(conn).DeleteArticle(947136363651727361)
	assert.Equal(t, nil, err)
}
