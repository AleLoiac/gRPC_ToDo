package main

import (
	"bufio"
	"fmt"
	"gRPC_ToDo/ToDopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //later on it will need to be secured
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	reader := bufio.NewReader(os.Stdin)

	c := ToDopb.NewToDoServiceClient(cc)

	doUnaryCreate(c, reader)
	doUnaryCreate(c, reader)

	doServerStreamingList(c)

	doUnaryCheck(c, reader)

	doServerStreamingList(c)

	doUnaryDelete(c, reader)

	doServerStreamingList(c)
}

func doUnaryCreate(c ToDopb.ToDoServiceClient, reader *bufio.Reader) {
	fmt.Println("Starting Unary RPC...")

	var title, desc string

	fmt.Println("Create new ToDo, write a title:")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Create new ToDo, write a description:")
	desc, _ = reader.ReadString('\n')
	desc = strings.TrimSpace(desc)

	req := &ToDopb.NewToDo{
		Title:       title,
		Description: desc,
	}

	//we call the function generated
	res, err := c.CreateToDo(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateToDo RPC: %v", err)
	}
	log.Printf("Response from CreateToDo: %v", res.GetTodo())
}

func doServerStreamingList(c ToDopb.ToDoServiceClient) {
	fmt.Println("Starting Server Streaming RPC...")

	req := &ToDopb.Empty{}

	resStream, err := c.ListToDos(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListToDos RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}
		log.Printf("Listing ToDos: %v", msg.GetTodo())
	}
}

func doUnaryCheck(c ToDopb.ToDoServiceClient, reader *bufio.Reader) {
	fmt.Println("Starting Unary RPC...")

	var id string

	fmt.Println("Check or Uncheck a ToDo, insert id:")
	id, _ = reader.ReadString('\n')
	id = strings.TrimSpace(id)

	req := &ToDopb.ToDoId{
		Id: id,
	}

	//we call the function generated
	res, err := c.CheckUncheck(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CheckUncheck RPC: %v", err)
	}
	log.Printf("Response from CheckUncheck: %v", res.GetTodo())
}

func doUnaryDelete(c ToDopb.ToDoServiceClient, reader *bufio.Reader) {
	fmt.Println("Starting Unary RPC...")

	var id string

	fmt.Println("Delete a ToDo, insert id:")
	id, _ = reader.ReadString('\n')
	id = strings.TrimSpace(id)

	req := &ToDopb.ToDoId{
		Id: id,
	}

	//we call the function generated
	_, err := c.DeleteToDo(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteToDo RPC: %v", err)
	}
	log.Printf("Response from DeleteToDo: Deleted")
}
