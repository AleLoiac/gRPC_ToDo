package main

import (
	"fmt"
	"gRPC_ToDo/ToDopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	ToDopb.ToDoServiceServer
}

var todoList []*ToDopb.ToDo

func (*server) CreateToDo(ctx context.Context, req *ToDopb.ToDo) (*ToDopb.ToDoResponse, error) {
	fmt.Printf("Create function is invoked with %v\n", req)
	id := req.GetId()
	title := req.GetTitle()
	description := req.GetDescription()
	done := req.GetDone()

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
		stream.Send(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Server started...")

	todoList = make([]*ToDopb.ToDo, 0)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ToDopb.RegisterToDoServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
