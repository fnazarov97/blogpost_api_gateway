package clients

import (
	"blogpost/config"
	"blogpost/genprotos/article"
	"blogpost/genprotos/author"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author  author.AuthorServicesClient
	Article article.ArticleServicesClient
	conns   []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	author := author.NewAuthorServicesClient(connAuthor)

	connArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	article := article.NewArticleServicesClient(connArticle)
	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Author:  author,
		Article: article,
		conns:   append(conns, connAuthor, connArticle),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
