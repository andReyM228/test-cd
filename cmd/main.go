package main

//import (
//	"black_box_backend/logger/zaplog"
//	"black_box_backend/repository"
//	"black_box_backend/server"
//	"black_box_backend/server/handlers"
//	"context"
//	"fmt"
//	"github.com/zeebo/errs"
//	"net"
//	"os"
//)

func main() {
	//log := zaplog.NewLog()
	//ctx := context.Background()
	//
	//listener, err := net.Listen("tcp", ":5554")
	//if err != nil {
	//	log.Fatal("creating a listener", err)
	//}
	//
	//rep, err := repository.New(ctx, "", fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
	//	os.Getenv("MONGO_USER"),
	//	os.Getenv("MONGO_PASSWORD"),
	//	os.Getenv("MONGO_IP"),
	//	os.Getenv("MONGO_PORT"),
	//	os.Getenv("MONGO_DB")))
	//
	//if err != nil {
	//	log.Fatal("creating a repository", err)
	//}
	//
	//h := handlers.New(log, rep)
	//serv := server.NewServer(log, listener, h)
	//
	//runErr := serv.Run(ctx)
	//closeErr := serv.Close()
	//if runErr != nil || closeErr != nil {
	//	log.Error("server error", errs.Combine(runErr, closeErr))
	//}
}
