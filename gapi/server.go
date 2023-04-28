package gapi

import (
	"fmt"

	db "github.com/nc-minh/tinybank/db/sqlc"
	"github.com/nc-minh/tinybank/pb"
	"github.com/nc-minh/tinybank/token"
	"github.com/nc-minh/tinybank/utils"
	"github.com/nc-minh/tinybank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedTinyBankServer
	config          utils.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker, taskDistributor: taskDistributor}

	return server, nil
}
