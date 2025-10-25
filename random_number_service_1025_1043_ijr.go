// 代码生成时间: 2025-10-25 10:43:30
package main

import (
    "fmt"
    "log"
    "math/rand"
    "net"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/protobuf/types/known/wrapperspb"
    "golang.org/x/net/context"
)

// Define the proto message and service in a separate file (random_number.proto)
// and generate Go code with `protoc --go_out=. --go-grpc_out=. random_number.proto`

// RandomNumberService is the server API for RandomNumber service.
type RandomNumberService struct {}

// GenerateRandomNumber generates a random number within the specified range.
func (s *RandomNumberService) GenerateRandomNumber(ctx context.Context, in *wrapperspb.Int32Value) (*wrapperspb.Int32Value, error) {
    if in.Value == nil {
        return nil, fmt.Errorf("input value is required")
    }
    rangeStart := *in.Value
    if rangeStart < 0 {
        return nil, fmt.Errorf("range start cannot be negative")
    }
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Int31n(int32(rangeStart)) + 1
    return &wrapperspb.Int32Value{Value: &randomNumber}, nil
}

// startService starts the gRPC server on the specified address.
func startService(address string) error {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        return err
    }
    fmt.Printf("Serving on %s
", address)
    s := grpc.NewServer()
    // Register the generated service.
    // randomNumberService.Register(s)
    return s.Serve(lis)
}

func main() {
    // Define the address to serve on.
    address := ":50051"
    if err := startService(address); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
