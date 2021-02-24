package excel2pdf

import (
	"context"
	"time"

	"github.com/annlumia/excel2pdf-grpc/pb"
	"google.golang.org/grpc"
)

var clientConnection *grpc.ClientConn

// Convert ...
func Convert(rpcAddress string, filename string) (string, error) {
	if rpcAddress == "" {
		rpcAddress = "localhost:8345"
	}

	clientConnection, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		return "", err
	}

	defer clientConnection.Close()

	client := pb.NewRouteConverterClient(clientConnection)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	reqMessage := &pb.ConvertRequest{
		InputFilename: filename,
	}

	result, err := client.Convert(ctx, reqMessage)
	if err != nil {
		return "", err
	}

	return result.OutputFilename, err
}
