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
	fmt.Println("Id not found")
	return nil, nil
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
	fmt.Println("Id not found")
	return nil, nil
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
