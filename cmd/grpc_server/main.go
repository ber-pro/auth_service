package main

import (
	userAPI "auth/internal/api/user"
	"auth/internal/config"
	repository "auth/internal/repository/user"
	service "auth/internal/service/user"
	desc "auth/pkg/user_v1"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	api userAPI.API
}

func main() {
	ctx := context.Background()

	err := config.Load(".env")
	if err != nil {
		log.Fatalf("Can't load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("Can't load grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("Can't load pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("fail to connect to db: %v", err)
	}
	defer pool.Close()

	api := userAPI.NewAPI(service.NewService(repository.NewRepository(pool)))

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, api)

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
