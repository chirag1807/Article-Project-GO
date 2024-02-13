package route

import (
	"articleproject/api/controller"
	"articleproject/api/middleware"
	"articleproject/api/repository"
	"articleproject/api/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func UsersRoutes(conn *pgx.Conn) (*chi.Mux) {
	r := chi.NewRouter()

	authRepository := repository.NewAuthRepo(conn)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	topicRepostiry := repository.NewTopicRepo(conn)
	topicService := service.NewTopicService(topicRepostiry)
	topicController := controller.NewTopicController(topicService)

	articleRepository := repository.NewArticleRepo(conn)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	followerRepository := repository.NewFollowerRepository(conn)
	followerService := service.NewFollowerService(followerRepository)
	followerController := controller.NewFollowerController(followerService)

	// r.Use(middlewares.SetDBConn(conn))
 
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/registration", authController.UserRegistration)

		r.Post("/login", authController.UserLogin)

		r.With(middleware.RetrieveRefreshToken()).Post("/refresh-token", authController.RefreshToken)
	})

	r.Route("/api/admin", func(r chi.Router) {
		r.Use(middleware.VerifyAccessToken(1))

		r.Post("/add-topic", topicController.AddTopic)

		r.Get("/get-all-topics", topicController.GetAllTopics)

		r.Put("/update-topic", topicController.UpdateTopic)

		r.Delete("/delete-topic/{ID}", topicController.DeleteTopic)
	})

	r.Route("/api/article", func(r chi.Router) {
		r.Use(middleware.VerifyAccessToken(0))

		r.Post("/add-article", articleController.AddArticle)

		r.Get("/get-my-articles", articleController.GetMyArticles)

		r.Get("/get-article-by-id/{ID}", articleController.GetArticleById)

		r.Put("/update-article", articleController.UpdateArticle)

		r.Delete("/delete-article/{ID}", articleController.DeleteArticle)

		r.Patch("/increase-view/{ID}", articleController.IncreaseView)

		r.Patch("/add-like/{ID}", articleController.AddLike)

		r.Patch("/remove-like/{ID}", articleController.RemoveLike)
	})

	r.Route("/api/user", func(r chi.Router) {
		r.Use(middleware.VerifyAccessToken(0))

		r.Post("/follow/{ID}", followerController.Follow)

		r.Delete("/unfollow/{ID}", followerController.UnFollow)

		r.Get("/myfollowers", followerController.MyFollowers)

		r.Get("/myfollwings", followerController.MyFollowings)

		r.Get("/someonefollowers/{ID}", followerController.SomeoneFollowers)

		r.Get("/someonefollowings/{ID}", followerController.SomeoneFollowings)
	})

	return r

}
