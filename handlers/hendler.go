package handlers

import (
	"blogpost/clients"
	"blogpost/config"
)

type Handler struct {
	Conf        config.Config
	grpcClients *clients.GrpcClients
}

func NewHandler(conf config.Config, grpcClients *clients.GrpcClients) Handler {
	return Handler{
		Conf:        conf,
		grpcClients: grpcClients,
	}
}
