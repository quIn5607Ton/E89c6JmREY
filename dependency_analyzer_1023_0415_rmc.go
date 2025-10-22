// 代码生成时间: 2025-10-23 04:15:33
// 依赖关系分析器
// 该程序用于分析项目依赖关系。

package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    "path/filepath"
    "strings"
)

// DependencyAnalyzerService 定义分析服务接口
type DependencyAnalyzerService struct{}

// AnalyzeDependencies 分析依赖关系
func (s *DependencyAnalyzerService) AnalyzeDependencies(ctx context.Context, req *AnalyzeRequest) (*AnalyzeResponse, error) {
    // 检查请求是否有效
    if req == nil || req.ProjectPath == "" {
        return nil, fmt.Errorf("invalid request")
    }

    // 递归分析依赖关系
    dependencies, err := analyzeProjectDependencies(req.ProjectPath)
    if err != nil {
        return nil, err
    }

    // 构造响应
    response := &AnalyzeResponse{Dependencies: dependencies}
    return response, nil
}

// analyzeProjectDependencies 递归分析项目依赖关系
func analyzeProjectDependencies(projectPath string) ([]string, error) {
    var dependencies []string
    // 这里假设有一个解析项目依赖关系的逻辑，例如解析Go Modules的go.mod文件
    // 为了示例，我们模拟依赖关系解析
    // 假设项目依赖了两个库："example.com/library1" 和 "example.com/library2"
    dependencies = []string{"example.com/library1", "example.com/library2"}
    return dependencies, nil
}

// AnalyzeRequest 定义分析请求结构体
type AnalyzeRequest struct {
    ProjectPath string
}

// AnalyzeResponse 定义分析响应结构体
type AnalyzeResponse struct {
    Dependencies []string
}

// 定义gRPC服务
type server struct{
    DependencyAnalyzerService
}

// 服务端主函数
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")
    s := grpc.NewServer()
    RegisterDependencyAnalyzerServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterDependencyAnalyzerServiceServer 注册服务到gRPC服务器
func RegisterDependencyAnalyzerServiceServer(s *grpc.Server, srv *server) {
    s.RegisterService(&_DependencyAnalyzerService_serviceDesc, srv)
}

// DependencyAnalyzerServiceServer 是gRPC服务接口
type DependencyAnalyzerServiceServer interface {
    AnalyzeDependencies(context.Context, *AnalyzeRequest) (*AnalyzeResponse, error)
}

// 服务描述
var _DependencyAnalyzerService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "DependencyAnalyzerService",
    HandlerType: (*AnalyzeRequest)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "AnalyzeDependencies",
            Handler: _DependencyAnalyzerService_AnalyzeDependencies_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "proto/dependency_analyzer.proto",
}

// _DependencyAnalyzerService_AnalyzeDependencies_Handler 实现AnalyzeDependencies方法
func _DependencyAnalyzerService_AnalyzeDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(AnalyzeRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(DependencyAnalyzerServiceServer).AnalyzeDependencies(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod:="/DependencyAnalyzerService/AnalyzeDependencies",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(DependencyAnalyzerServiceServer).AnalyzeDependencies(ctx, req.(*AnalyzeRequest))
    }
    return interceptor(ctx, in, info, handler)
}
