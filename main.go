package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"grpc-go/infra"
	pbUser "grpc-go/protos/user"
	repoUser "grpc-go/repository/user"
	"grpc-go/usecase/user"
	"log"
	"net"
	"os"
)

func main() {
	db := infra.NewGormDB()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userRepo := repoUser.NewUserRepo(db)
	userUC := user.NewUserUseCase(userRepo)

	s := grpc.NewServer()
	pbUser.RegisterUserServiceServer(s, userUC)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func init() {
	err := configSetENV()

	if err != nil {
		panic(err)
	}

	systemEnv := os.Getenv("APP_ENV")

	err = godotenv.Load(systemEnv + ".env")
	if err != nil {
		err = godotenv.Load("../" + systemEnv + ".env")
		if err != nil {
			panic(err)
		}
	}

}

func configSetENV() (err error) {
	if len(os.Args) > 1 {
		err = os.Setenv("APP_ENV", os.Args[1])
	} else {
		err = os.Setenv("APP_ENV", "local")
	}
	return
}
