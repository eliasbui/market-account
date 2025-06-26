package infrastructure

import (
	"context"
	"log"
	"net"
	"user-service/application"
	"user-service/domain"
	userpb "user-service/proto/user/proto/user"

	"google.golang.org/grpc"
)

// UserGRPCServer implements the generated UserServiceServer interface
type UserGRPCServer struct {
	userpb.UnimplementedUserServiceServer
	userService *application.UserService
}

// NewUserGRPCServer creates a new gRPC server for user service
func NewUserGRPCServer(userService *application.UserService) *UserGRPCServer {
	return &UserGRPCServer{
		userService: userService,
	}
}

// CreateUser implements the CreateUser RPC method
func (s *UserGRPCServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Validate request
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return &userpb.CreateUserResponse{
			Result: &userpb.CreateUserResponse_Error{
				Error: &userpb.Error{
					Code:    "INVALID_REQUEST",
					Message: "Email, username, and password are required",
				},
			},
		}, nil
	}

	// Create user via application service
	user, err := s.userService.CreateUser(ctx, req.Email, req.Username, req.Password)
	if err != nil {
		return &userpb.CreateUserResponse{
			Result: &userpb.CreateUserResponse_Error{
				Error: &userpb.Error{
					Code:    "CREATION_FAILED",
					Message: err.Error(),
				},
			},
		}, nil
	}

	// Convert domain user to protobuf user
	pbUser := convertDomainUserToProto(user)

	return &userpb.CreateUserResponse{
		Result: &userpb.CreateUserResponse_User{
			User: pbUser,
		},
	}, nil
}

// GetUser implements the GetUser RPC method
func (s *UserGRPCServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if req.Id == "" {
		return &userpb.GetUserResponse{
			Result: &userpb.GetUserResponse_Error{
				Error: &userpb.Error{
					Code:    "INVALID_REQUEST",
					Message: "User ID is required",
				},
			},
		}, nil
	}

	user, err := s.userService.GetUserByID(ctx, req.Id)
	if err != nil {
		return &userpb.GetUserResponse{
			Result: &userpb.GetUserResponse_Error{
				Error: &userpb.Error{
					Code:    "NOT_FOUND",
					Message: err.Error(),
				},
			},
		}, nil
	}

	pbUser := convertDomainUserToProto(user)

	return &userpb.GetUserResponse{
		Result: &userpb.GetUserResponse_User{
			User: pbUser,
		},
	}, nil
}

// GetUserByEmail implements the GetUserByEmail RPC method
func (s *UserGRPCServer) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUserResponse, error) {
	if req.Email == "" {
		return &userpb.GetUserResponse{
			Result: &userpb.GetUserResponse_Error{
				Error: &userpb.Error{
					Code:    "INVALID_REQUEST",
					Message: "Email is required",
				},
			},
		}, nil
	}

	user, err := s.userService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &userpb.GetUserResponse{
			Result: &userpb.GetUserResponse_Error{
				Error: &userpb.Error{
					Code:    "NOT_FOUND",
					Message: err.Error(),
				},
			},
		}, nil
	}

	pbUser := convertDomainUserToProto(user)

	return &userpb.GetUserResponse{
		Result: &userpb.GetUserResponse_User{
			User: pbUser,
		},
	}, nil
}

// ValidateUser implements the ValidateUser RPC method
func (s *UserGRPCServer) ValidateUser(ctx context.Context, req *userpb.ValidateUserRequest) (*userpb.ValidateUserResponse, error) {
	if req.Email == "" || req.Password == "" {
		return &userpb.ValidateUserResponse{
			IsValid: false,
			Error: &userpb.Error{
				Code:    "INVALID_REQUEST",
				Message: "Email and password are required",
			},
		}, nil
	}

	isValid, user, err := s.userService.ValidateCredentials(ctx, req.Email, req.Password)
	if err != nil {
		return &userpb.ValidateUserResponse{
			IsValid: false,
			Error: &userpb.Error{
				Code:    "VALIDATION_ERROR",
				Message: err.Error(),
			},
		}, nil
	}

	response := &userpb.ValidateUserResponse{
		IsValid: isValid,
	}

	if isValid && user != nil {
		response.User = convertDomainUserToProto(user)
	}

	return response, nil
}

// UpdateUser implements the UpdateUser RPC method (simplified version)
func (s *UserGRPCServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	// Simplified implementation - in real scenario you'd update specific fields
	return &userpb.UpdateUserResponse{
		Result: &userpb.UpdateUserResponse_Error{
			Error: &userpb.Error{
				Code:    "NOT_IMPLEMENTED",
				Message: "Update user is not yet implemented",
			},
		},
	}, nil
}

// DeleteUser implements the DeleteUser RPC method (simplified version)
func (s *UserGRPCServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	// Simplified implementation
	return &userpb.DeleteUserResponse{
		Success: false,
		Error: &userpb.Error{
			Code:    "NOT_IMPLEMENTED",
			Message: "Delete user is not yet implemented",
		},
	}, nil
}

// GetUsers implements the GetUsers RPC method (simplified version)
func (s *UserGRPCServer) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	// Simplified implementation
	return &userpb.GetUsersResponse{
		Users: []*userpb.User{},
		Error: &userpb.Error{
			Code:    "NOT_IMPLEMENTED",
			Message: "Get users batch is not yet implemented",
		},
	}, nil
}

// Helper function to convert domain user to protobuf user
func convertDomainUserToProto(user *domain.User) *userpb.User {
	return &userpb.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		Username:  user.Username,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Roles:     []string{}, // TODO: Convert user roles
	}
}

// StartGRPCServer starts the gRPC server on the specified port
func StartGRPCServer(userService *application.UserService, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	userGRPCServer := NewUserGRPCServer(userService)

	userpb.RegisterUserServiceServer(grpcServer, userGRPCServer)

	log.Printf("gRPC server listening on port %s", port)
	return grpcServer.Serve(lis)
}
