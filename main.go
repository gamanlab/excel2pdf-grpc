package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/annlumia/excel2pdf-grpc/office2pdf"
	"github.com/annlumia/excel2pdf-grpc/pb"
	"github.com/teris-io/shortid"
	"google.golang.org/grpc"
)

// RPCServer grpc server yang disediakan untuk melayani permintaan dari aplikasi.
var RPCServer *grpc.Server

// RPCServerAddress Address of RPC server
var RPCServerAddress string

// const tempDirectory = ".temp"

var (
	flagRPCPort = flag.Uint("port", 8345, "Port number of gRPC server")
)

type routeConverterServer struct {
	pb.UnimplementedRouteConverterServer
}

// GenerateShortUID generates a short unique identifier.
func generateShortUID() string {
	return shortid.MustGenerate()
}

func (r *routeConverterServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	excelFilename := req.InputFilename

	excel2pdf := office2pdf.Excel{}
	pdf, err := excel2pdf.Export(excelFilename)
	if err != nil {
		return nil, err
	}

	return &pb.ConvertResponse{OutputFilename: pdf}, nil
}

func (r *routeConverterServer) Exit(ctx context.Context, req *pb.ExitRequest) (*pb.ExitResponse, error) {
	go func() {
		time.Sleep(time.Second * 3)
		os.Exit(int(req.Status))
	}()

	return &pb.ExitResponse{Status: req.Status}, nil
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf("127.0.0.1:%v", *flagRPCPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("E! Failed to start RPC server. %s\n", err.Error())
	}

	RPCServerAddress = listener.Addr().String()

	RPCServer = grpc.NewServer()

	pb.RegisterRouteConverterServer(RPCServer, &routeConverterServer{})

	log.Printf("I! RPC server listening on %s\n", RPCServerAddress)
	if err := RPCServer.Serve(listener); err != nil {
		log.Printf("E! Failed to serve gRPC server. %s\n", err.Error())
	}
}
