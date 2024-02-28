package controller

import (
	testcase "articleproject/api/model/dto/test-cases"
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestAddArticle(t *testing.T) {
	testCases := []struct {
		TestCaseName string
		Title        string
		Content      string
		Topic        int64
		Author       int64
		Expected     string
		StatusCode   int
	}{
		{
			TestCaseName: "Success",
			Title:        "React JS",
			Content:      "This Article is About React JS.",
			Topic:        943765517224148993,
			Author:       945437407105286145,
			Expected:     "Article Added Successfully.",
			StatusCode:   200,
		},
	}

	for _, v := range testCases {
		t.Run(v.TestCaseName, func(t *testing.T) {
			r.Post("/api/article/add-article", NewArticleController(articleService).AddArticle)

			article := request.Article{
				Title:   v.Title,
				Content: v.Content,
				Topic:   v.Topic,
			}
			jsonValue, _ := json.Marshal(article)

			req, _ := http.NewRequest("POST", "/api/article/add-article", bytes.NewBuffer(jsonValue))
			ctx := context.WithValue(req.Context(), ContextKeyID, v.Author)
			req = req.WithContext(ctx)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response testcase.Article
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, v.Expected, response.Message)
			assert.Equal(t, v.StatusCode, w.Code)
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
					ID:      946881922904522753,
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
			r.Get("/api/article/get-my-articles", NewArticleController(articleService).GetMyArticles)

			req, _ := http.NewRequest("GET", "/api/article/get-my-articles", http.NoBody)
			ctx := context.WithValue(req.Context(), ContextKeyID, v.Author)
			req = req.WithContext(ctx)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response response.ArticleResponse
			json.Unmarshal(w.Body.Bytes(), &response)

			assert.Equal(t, v.Expected[0].ID, response.Article[0].ID)
			assert.Equal(t, v.StatusCode, w.Code)
		})
	}
}

func TestUpdateArticle(t *testing.T) {
	r.Put("/api/article/update-article", NewArticleController(articleService).UpdateArticle)

	article := request.Article{
		ID:      946881922904522753,
		Title:   "ReactJS-Web Framework",
		Content: "This article is about ReactJS, A web framework.",
		Topic:   943765517224148993,
	}
	jsonValue, _ := json.Marshal(article)

	req, _ := http.NewRequest("PUT", "/api/article/update-article", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response testcase.Article
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Article Updated Successfully.", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteArticle(t *testing.T) {
	r.Delete("/api/article/delete-article/:ID", NewArticleController(articleService).DeleteArticle)

	req, _ := http.NewRequest("DELETE", "/api/article/delete-article/:ID", http.NoBody)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("ID", "946881922904522753")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	fmt.Println(chi.URLParam(req, "ID"))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response testcase.Article
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Article Deleted Successfully.", response.Message)
	assert.Equal(t, 200, w.Code)
}
