package main

import (
	"articleproject/api/repository"
	"articleproject/config"
	"articleproject/db"
	proto "articleproject/protoc"
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type article struct {
	proto.UnimplementedGetMyArticlesServer
}

var conn *pgx.Conn

func main() {
	config.LoadEnv("../.config/")
	conn, _, _, _ = db.DBConnection()
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	proto.RegisterGetMyArticlesServer(srv, &article{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Fatal(e)
	}
}

func (a *article) GetMyArticles(c context.Context, req *proto.GetMyArticleRequest) (*proto.ArticleResponse, error) {
	res, _ := repository.NewArticleRepo(conn).GetMyArticles(req.ID)
	var protoResponse []*proto.GetMyArticleResponse
	for _, v := range res {
		protoResponse = append(protoResponse, &proto.GetMyArticleResponse{
			ID:      v.ID,
			Title:   v.Title,
			Content: v.Content,
			Image: v.Image,
			Topic: v.Topic,
			Likes: int32(v.Likes),
			Views: int32(v.Views),
			PublishedAt: v.PublishedAt.String(),
		})
	}
	return &proto.ArticleResponse{
		Articles: protoResponse,
	}, nil
}
