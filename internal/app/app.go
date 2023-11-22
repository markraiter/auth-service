package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/markraiter/auth-service/internal/app/grpc"
)

type App struct {
	GRPCSrv  *grpcapp.App
	grpcPort int
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: initialize storage

	// TODO: init service layer

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
