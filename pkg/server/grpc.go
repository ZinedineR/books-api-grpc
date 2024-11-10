package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"os"
	"time"
)

func NewGRPCServer(port string) (*grpc.Server, net.Listener) {
	server := grpc.NewServer()
	serverConnection, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return server, serverConnection
}

func NewGRPCGatewayClient(server *grpc.Server, net net.Listener, port string) *grpc.ClientConn {
	go func() {
		if err := server.Serve(net); err != nil {
			slog.Error("failed to serve: ", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	client, err := grpc.NewClient(
		"0.0.0.0:"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return client
}

func NewGRPCListenerClient(port string) *grpc.ClientConn {
	var (
		client *grpc.ClientConn
		err    error
	)
	for {
		client, err = grpc.NewClient(
			"0.0.0.0:"+port,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			slog.Error(fmt.Sprintf("failed to connect to port %s of grpc client", port), "error", err.Error())
			slog.Info(fmt.Sprintf("retrying to connect to port %s of grpc client in 5 seconds...", port))
			time.Sleep(5 * time.Second)
			continue
		}
		slog.Info(fmt.Sprintf("successfully connected to port %s of grpc client", port))
		break
	}
	return client
}
