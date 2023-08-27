package main

import (
	"fmt"
	"net"

	db "github.com/cvzamannow/service-auth/internal/db"
	pb "github.com/cvzamannow/service-auth/internal/proto"
	"github.com/cvzamannow/service-auth/internal/repository"
	"github.com/cvzamannow/service-auth/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC Server is running....")
	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		panic("failed to listen gRPC Server!")
	}

	godotenv.Load(".env")

	// // AWS Relational Database Credential
	// RDS_HOST := os.Getenv("RDS_HOST")
	// RDS_USER := os.Getenv("RDS_USER")
	// RDS_PASSWORD := os.Getenv("RDS_PASSWORD")
	// RDS_DB := os.Getenv("RDS_DB")
	// RDS_PORT := os.Getenv("RDS_PORT")

	conn := db.NewDBConnection("postgres", "postgres://root:root@localhost:5434/cvz_auth?sslmode=disable")
	authRepository := repository.NewAuthRepository(conn)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &service.Server{
		Repository: authRepository,
	})

	fmt.Printf("Server running on %v", lis.Addr())

	if errS := s.Serve(lis); errS != nil {
		panic("Failed to start gRPC Server!")
	}

}
