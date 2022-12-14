package main

import (
	"stackcoder/handler"
	pb "stackcoder/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("stackcoder"),
	)

	// Register handler
	pb.RegisterStackcoderHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
