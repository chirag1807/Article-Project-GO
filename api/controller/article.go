package controller

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"articleproject/api/service"
	"articleproject/constants"
	errorhandling "articleproject/error"
	"articleproject/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ArticleController interface {
	AddArticle(w http.ResponseWriter, r *http.Request)
	GetMyArticles(w http.ResponseWriter, r *http.Request)
	GetArticleById(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
	IncreaseView(w http.ResponseWriter, r *http.Request)
	AddLike(w http.ResponseWriter, r *http.Request)
	RemoveLike(w http.ResponseWriter, r *http.Request)
}

type articleController struct {
	articleService service.ArticleService
}

func NewArticleController(a service.ArticleService) ArticleController {
	return articleController{
		articleService: a,
	}
}

func (a articleController) AddArticle(w http.ResponseWriter, r *http.Request) {
	var article request.Article

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}

	err = json.Unmarshal(body, &article)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	author := r.Context().Value("id").(int64)
	article.Author = author
	err = a.articleService.AddArticle(article)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a articleController) GetMyArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := a.articleService.GetMyArticles(r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.ArticleResponse{
		Article: articles,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a articleController) GetArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	article, err := a.articleService.GetArticleById(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	articles := make([]response.Article, 0)
	articles = append(articles, article)
	response := response.ArticleResponse{
		Article: articles,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a articleController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article request.Article
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.articleService.UpdateArticle(article)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_UPDATED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
}

func (a articleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	fmt.Println(chi.URLParam(r, "ID"))
	fmt.Println(id)
	err := a.articleService.DeleteArticle(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_DELETED,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (a articleController) IncreaseView(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := a.articleService.IncreaseView(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_VIEW_INCREASED,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (a articleController) AddLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := a.articleService.AddLike(id, r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_LIKE_ADDED,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (a articleController) RemoveLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := a.articleService.RemoveLike(id, r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_LIKE_REMOVED,
	}
	utils.ResponseGenerator(w, 200, response)
}
