package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	pb "restful/restful"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "doankimhuy"
	password = "mysql"
	dbname   = "mydb"
)

type server struct {
	pb.UnimplementedRestfulServer
}

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}
	return db, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	taskId := in.GetId()
	if taskId <= 0 {
		return &pb.GetResponse{Id: taskId}, errors.New("invalid request info")
	}
	db, err := ConnectDB()
	if err != nil {
		return &pb.GetResponse{Id: taskId}, err
	}

	var (
		title  string
		status string
	)
	err = db.QueryRow(fmt.Sprintf("SELECT title, status FROM todo WHERE id = %v;", taskId)).Scan(&title, &status)
	if err != nil {
		return &pb.GetResponse{Id: taskId}, err
	}
	return &pb.GetResponse{Id: taskId, Title: title, Status: status}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	taskId := in.GetId()
	taskTitle := in.GetTitle()
	taskStatus := in.GetStatus()
	if taskId <= 0 || taskTitle == "" || taskStatus == "" {
		return &pb.CreateResponse{Status: "error"}, errors.New("invalid request info")
	}
	db, err := ConnectDB()
	if err != nil {
		return &pb.CreateResponse{Status: "error"}, err
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO todo VALUES (%v, '%s', '%s');", taskId, taskTitle, taskStatus))
	if err != nil {
		return &pb.CreateResponse{Status: "error"}, err
	}
	return &pb.CreateResponse{Status: "ok"}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	taskId := in.GetId()
	taskTitle := in.GetTitle()
	taskStatus := in.GetStatus()
	if taskId <= 0 {
		return &pb.UpdateResponse{Status: "error"}, errors.New("invalid request info")
	}
	db, err := ConnectDB()
	if err != nil {
		return &pb.UpdateResponse{Status: "error"}, err
	}
	if taskTitle != "" && taskStatus != "" {
		err = Exec(db, fmt.Sprintf("UPDATE todo SET title='%s', status='%s' WHERE id='%v';", taskTitle, taskStatus, taskId))
	} else if taskTitle != "" {
		err = Exec(db, fmt.Sprintf("UPDATE todo SET title='%s' WHERE id='%v';", taskTitle, taskId))
	} else if taskStatus != "" {
		err = Exec(db, fmt.Sprintf("UPDATE todo SET status='%s' WHERE id='%v';", taskStatus, taskId))
	}
	if err != nil {
		return &pb.UpdateResponse{Status: "error"}, err
	}
	return &pb.UpdateResponse{Status: "ok"}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	taskId := in.GetId()
	if taskId <= 0 {
		return &pb.DeleteResponse{Status: "error"}, errors.New("invalid request info")
	}
	db, err := ConnectDB()
	if err != nil {
		return &pb.DeleteResponse{Status: "error"}, err
	}
	_, err = db.Exec(fmt.Sprintf("DELETE FROM todo WHERE id=%v;", taskId))
	if err != nil {
		return &pb.DeleteResponse{Status: "error"}, err
	}
	return &pb.DeleteResponse{Status: "ok"}, nil
}

func Exec(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7501))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRestfulServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
