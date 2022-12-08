package clients

import (
	"blogpost/config"
	"blogpost/genprotos/article"
	"blogpost/genprotos/author"
	"blogpost/genprotos/authorization"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author        author.AuthorServicesClient
	Article       article.ArticleServicesClient
	Authorization authorization.AuthServiceClient
	conns         []*grpc.ClientConn
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

	connAuthorization, err := grpc.Dial(cfg.AuthorizationServiceGrpcHost+cfg.AuthorizationServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	authorization := authorization.NewAuthServiceClient(connAuthorization)
	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Author:        author,
		Article:       article,
		Authorization: authorization,
		conns:         append(conns, connAuthor, connArticle, connAuthorization),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
