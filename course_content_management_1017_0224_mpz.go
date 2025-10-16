// 代码生成时间: 2025-10-17 02:24:29
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
    "google.golang.org/protobuf/types/known/timestamppb"

    "pb "course_content_management_proto" // Assuming the protocol buffer definitions are in this package
)

// CourseContentService represents the server structure for the Course Content Management service
type CourseContentService struct {
    pb.UnimplementedCourseContentManagementServer
    courses map[string]*pb.CourseContent
}

// NewCourseContentService creates a new instance of the CourseContentService
func NewCourseContentService() *CourseContentService {
    return &CourseContentService{
        courses: make(map[string]*pb.CourseContent),
    }
}

// AddCourseContent adds a new course content to the service
func (s *CourseContentService) AddCourseContent(ctx context.Context, req *pb.AddCourseContentRequest) (*pb.CourseContent, error) {
    if _, exists := s.courses[req.GetId()]; exists {
        return nil, fmt.Errorf("course content with ID %s already exists", req.GetId())
    }

    course := &pb.CourseContent{
        Id:       req.GetId(),
        Title:    req.GetTitle(),
        Content:  req.GetContent(),
        CreatedAt: timestamppb.Now(),
    }
    s.courses[req.GetId()] = course
    return course, nil
}

// GetCourseContent retrieves a course content by its ID
func (s *CourseContentService) GetCourseContent(ctx context.Context, req *pb.GetCourseContentRequest) (*pb.CourseContent, error) {
    course, exists := s.courses[req.GetId()]
    if !exists {
        return nil, fmt.Errorf("course content with ID %s not found", req.GetId())
    }
    return course, nil
}

// UpdateCourseContent updates an existing course content
func (s *CourseContentService) UpdateCourseContent(ctx context.Context, req *pb.UpdateCourseContentRequest) (*pb.CourseContent, error) {
    course, exists := s.courses[req.GetId()]
    if !exists {
        return nil, fmt.Errorf("course content with ID %s not found", req.GetId())
    }

    course.Title = req.GetTitle()
    course.Content = req.GetContent()
    return course, nil
}

// DeleteCourseContent removes a course content from the service
func (s *CourseContentService) DeleteCourseContent(ctx context.Context, req *pb.DeleteCourseContentRequest) (*emptypb.Empty, error) {
    if _, exists := s.courses[req.GetId()]; !exists {
        return nil, fmt.Errorf("course content with ID %s not found", req.GetId())
    }
    delete(s.courses, req.GetId())
    return &emptypb.Empty{}, nil
}

// StartServer starts the gRPC server
func StartServer(address string, service *CourseContentService) error {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        return fmt.Errorf("failed to listen: %v", err)
    }
    fmt.Printf("Listening on %s
", address)

    s := grpc.NewServer()
    pb.RegisterCourseContentManagementServer(s, service)
    if err := s.Serve(lis); err != nil {
        return fmt.Errorf("failed to serve: %v", err)
    }
    return nil
}

func main() {
    serverAddress := ":50051"
    courseService := NewCourseContentService()
    if err := StartServer(serverAddress, courseService); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}