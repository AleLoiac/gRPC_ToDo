package main

import (
	"fmt"
	"gRPC_ToDo/ToDopb"
	"github.com/dgraph-io/badger/v3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
	ToDopb.ToDoServiceServer
}

var todoList []*ToDopb.ToDo

var serial int

func (*server) CreateToDo(ctx context.Context, req *ToDopb.NewToDo) (*ToDopb.ToDoResponse, error) {
	fmt.Printf("Create function is invoked with %v\n", req)
	id := strconv.Itoa(serial)
	title := req.GetTitle()
	description := req.GetDescription()
	done := false

	serial++

	newTodo := &ToDopb.ToDo{
		Id:          id,
		Title:       title,
		Description: description,
		Done:        done,
	}

	//sending a response, look at the structure in the generated code
	res := &ToDopb.ToDoResponse{
		Todo: newTodo,
	}

	err := db.Update(func(txn *badger.Txn) error {
		data, err := proto.Marshal(newTodo)
		if err != nil {
			return err
		}
		return txn.Set([]byte(id), data)
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	gettodo := &ToDopb.ToDo{}

	err2 := db.View(func(txn *badger.Txn) error {
		data, err3 := txn.Get([]byte(id))
		data.Value(func(val []byte) error {
			err = proto.Unmarshal(val, gettodo)
			if err != nil {
				return err
			}
			return nil
		})
		//need to transform data in bytes
		fmt.Println(gettodo)
		return err3
	})
	if err2 != nil {
		return nil, status.Error(codes.Internal, err2.Error())
	}

	todoList = append(todoList, newTodo)

	fmt.Println(todoList)
	return res, nil
}

func (*server) ListToDos(req *ToDopb.Empty, stream ToDopb.ToDoService_ListToDosServer) error {
	fmt.Println("ListToDos function is invoked with an empty request")

	for i := range todoList {
		res := &ToDopb.ToDoResponse{
			Todo: &ToDopb.ToDo{
				Id:          todoList[i].GetId(),
				Title:       todoList[i].GetTitle(),
				Description: todoList[i].GetDescription(),
				Done:        todoList[i].GetDone(),
			},
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (*server) CheckUncheck(ctx context.Context, req *ToDopb.ToDoId) (*ToDopb.ToDoResponse, error) {
	fmt.Printf("CheckUncheck function is invoked with %v\n", req)
	id := req.GetId()

	for i := range todoList {
		if id == todoList[i].GetId() {
			todoList[i].Done = !todoList[i].GetDone()
			res := &ToDopb.ToDoResponse{
				Todo: &ToDopb.ToDo{
					Id:          todoList[i].GetId(),
					Title:       todoList[i].GetTitle(),
					Description: todoList[i].GetDescription(),
					Done:        todoList[i].GetDone(),
				},
			}
			return res, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find todo with id: %v", req.GetId()))
}

func (*server) DeleteToDo(ctx context.Context, req *ToDopb.ToDoId) (*ToDopb.Empty, error) {
	fmt.Printf("DeleteToDo function is invoked with %v\n", req)
	id := req.GetId()

	for i := range todoList {
		if id == todoList[i].GetId() {
			todoList = append(todoList[:i], todoList[i+1:]...)
			res := &ToDopb.Empty{}
			fmt.Println("Todo deleted")
			return res, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find todo with id: %v", req.GetId()))
}

var db *badger.DB

func main() {
	fmt.Println("Server started...")

	db, _ = badger.Open(badger.DefaultOptions("/tmp/badger"))

	defer db.Close()

	todoList = make([]*ToDopb.ToDo, 0)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ToDopb.RegisterToDoServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
