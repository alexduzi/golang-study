package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/alexduzi/golang-study/golanggrpc/internal/database"
	pb "github.com/alexduzi/golang-study/golanggrpc/internal/pb"
	"github.com/alexduzi/golang-study/golanggrpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	s := grpc.NewServer()
	pb.RegisterCategoryServiceServer(s, categoryService)

	// Register reflection service on gRPC server.
	reflection.Register(s) // This is the crucial step

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
