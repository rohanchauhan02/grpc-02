package main

import (
	"context"
	"log"
	"net"

	"githiub.com/rohanchauhan02/grpc-02/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	// invoicer.InvoicerServer
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Msg: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("can not create listener:%v", err)
	}
	log.Printf("Server running on port: %v", "8089")
	servierRegister := grpc.NewServer()
	invoicer.RegisterInvoicerServer(servierRegister, &myInvoicerServer{})
	err = servierRegister.Serve(lis)

	if err != nil {
		log.Fatalf("impossible to serve: %v", err)
	}
}
