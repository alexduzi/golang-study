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

// to generate the proto buffer interfaces you need to execute this below command
// protoc --go_out=. --go-grpc_out=. proto/course_category.proto
// check if the server is working
// grpcurl -plaintext localhost:50051 list
// use evans to make a call to gRPC server
// evans --host localhost -p 50051 -r
// go to the package with package pb in this case
// service CategoryService to set the service
// call 'Name of Method' to call the endpoint

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
