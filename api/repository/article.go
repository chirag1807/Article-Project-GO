package repository

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"context"

	"github.com/jackc/pgx/v5"
)

type ArticleRepository interface {
	AddArticle(request.Article) error
	GetMyArticles(int64) ([]response.Article, error)
	GetArticleById(int64) (response.Article, error)
	UpdateArticle(request.Article) error
	DeleteArticle(int64) error
	IncreaseView(int64) error
	AddLike(int64, int64) error
	RemoveLike(int64, int64) error
}

type articleRepository struct {
	pgx *pgx.Conn
}

func NewArticleRepo(pgx *pgx.Conn) ArticleRepository {
	return articleRepository{
		pgx: pgx,
	}
}

func (a articleRepository) AddArticle(article request.Article) error {
	_, err := a.pgx.Exec(context.Background(), `INSERT INTO articles (title, content, image, topic, author) VALUES ($1, $2, $3, $4, $5)`, article.Title, article.Content, article.Image, article.Topic, article.Author)
	if err != nil {
		return err
	}
	return nil
}

func (a articleRepository) GetMyArticles(author int64) ([]response.Article, error) {
	articles, err := a.pgx.Query(context.Background(), `SELECT * FROM articles WHERE author = $1`, author)

	articlesSlice := make([]response.Article, 0)
	if err != nil {
		return articlesSlice, err
	}
	defer articles.Close()

	var article response.Article
	for articles.Next() {
		if err := articles.Scan(&article.ID, &article.Title, &article.Content, &article.Image, &article.Likes, &article.Views, &article.Topic, &article.Author, &article.PublishedAt); err != nil {
			return articlesSlice, err
		}
		articlesSlice = append(articlesSlice, article)
	}

	return articlesSlice, nil
}

func (a articleRepository) GetArticleById(id int64) (response.Article, error) {
	article, err := a.pgx.Query(context.Background(), `SELECT * FROM articles WHERE id = $1`, id)
	var responseArticle response.Article
	defer article.Close()

	if err != nil {
		return responseArticle, err
	}

	err = article.Scan(&responseArticle.ID, &responseArticle.Title, &responseArticle.Content, &responseArticle.Image, &responseArticle.Likes, &responseArticle.Views, &responseArticle.Topic, &responseArticle.Author, &responseArticle.PublishedAt)
	if err != nil {
		return responseArticle, err
	}
	return responseArticle, nil
}

func (a articleRepository) UpdateArticle(article request.Article) error {
	_, err := a.pgx.Exec(context.Background(), `UPDATE articles SET title = $1, content = $2, image = $3, topic = $4 WHERE id = $5`, article.Title, article.Content, article.Image, article.Topic, article.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a articleRepository) DeleteArticle(id int64) error {
	_, err := a.pgx.Exec(context.Background(), `DELETE FROM articles WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (a articleRepository) IncreaseView(articleid int64) error {
	_, err := a.pgx.Exec(context.Background(), `UPDATE articles SET noofviews = noofviews + 1 WHERE id = $1`, articleid)
	if err != nil {
		return err
	}
	return nil
}

func (a articleRepository) AddLike(articleid int64, userid int64) error {
	ctx := context.Background()
	tx, err := a.pgx.Begin(ctx)
	if err != nil {
		return err
	}

	batch := &pgx.Batch{}
	batch.Queue(`INSERT INTO likes (userid, articleid) VALUES ($1, $2)`, userid, articleid)
	batch.Queue(`UPDATE articles SET nooflikes = nooflikes + 1 WHERE id = $1`, articleid)

	results := tx.SendBatch(ctx, batch)
	defer results.Close()

	if err := results.Close(); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}

func (a articleRepository) RemoveLike(articleid int64, userid int64) error {
	ctx := context.Background()
	tx, err := a.pgx.Begin(ctx)
	if err != nil {
		return err
	}

	batch := &pgx.Batch{}
	batch.Queue(`DELETE FROM likes WHERE userid = $1 AND articleid = $2`, userid, articleid)
	batch.Queue(`UPDATE articles SET nooflikes = nooflikes - 1 WHERE id = $1`, articleid)

	results := tx.SendBatch(ctx, batch)
	defer results.Close()

	if err := results.Close(); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}
