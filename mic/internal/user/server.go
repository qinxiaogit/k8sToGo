package user

import (
	"context"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/user/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func NewServer(logger *log.Logger, repo Repository) v1.UserServiceServer {
	return &Server{
		logger: logger,
		repo:   repo,
	}
}

type Server struct {
	v1.UnimplementedUserServiceServer
	logger *log.Logger
	repo   Repository
}

func (s Server) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	id := req.GetId()
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete user: %v", err)
	}
	return &v1.DeleteUserResponse{
		Success: true,
	}, nil
}

func (s Server) ListUsersByIDs(ctx context.Context, req *v1.ListUsersByIDsRequest) (*v1.ListUsersByIDsResponse, error) {
	ids := req.GetIds()
	users, err := s.repo.ListUsersByIDs(ctx, ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user list by ids: %v", err)
	}
	protoUsers := make([]*v1.User, len(users))
	for i, user := range users {
		protoUsers[i] = entityToProtobuf(user)
	}
	resp := &v1.ListUsersByIDsResponse{
		Users: protoUsers,
	}
	return resp, nil
}

func (s Server) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	id := req.GetId()
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	resp := &v1.GetUserResponse{
		User: entityToProtobuf(user),
	}

	return resp, nil
}

func (s Server) GetUserByEmail(ctx context.Context, req *v1.GetUserByEmailRequest) (*v1.GetUserResponse, error) {
	email := req.GetEmail()
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user by email: %v", err)
	}
	// 如果传递了密码，则需要验证密码
	if req.GetPassword() != "" {
		ok := isCorrectPassword(user.Password, req.GetPassword())
		if !ok {
			return nil, status.Errorf(codes.Internal, "incorrect password")
		}
	}
	resp := &v1.GetUserResponse{
		User: entityToProtobuf(user),
	}

	return resp, nil
}

func (s Server) GetUserByUsername(ctx context.Context, req *v1.GetUserByUsernameRequest) (*v1.GetUserResponse, error) {
	username := req.GetUsername()
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user by username: %v", err)
	}
	// 如果传递了密码，则需要验证密码
	if req.GetPassword() != "" {
		ok := isCorrectPassword(user.Password, req.GetPassword())
		if !ok {
			return nil, status.Errorf(codes.Internal, "incorrect password")
		}
	}

	resp := &v1.GetUserResponse{
		User: entityToProtobuf(user),
	}

	return resp, nil
}

func (s Server) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	password, err := generateFromPassword(req.GetUser().GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
	}
	user := &User{
		UUID:     uuid.New().String(),
		Username: req.GetUser().GetUsername(),
		Email:    req.GetUser().GetEmail(),
		Avatar:   req.GetUser().GetAvatar(),
		Password: password,
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	resp := &v1.CreateUserResponse{
		User: entityToProtobuf(user),
	}

	return resp, nil
}

func (s Server) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	user, err := s.repo.Get(ctx, req.GetUser().GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get user: %v", err)
	}

	if req.GetUser().GetUsername() != "" {
		user.Username = req.GetUser().GetUsername()
	}
	if req.GetUser().Email != "" {
		user.Email = req.GetUser().Email
	}
	if req.GetUser().Avatar != "" {
		user.Avatar = req.GetUser().Avatar
	}
	if req.GetUser().GetPassword() != "" {
		password, err := generateFromPassword(req.GetUser().GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to bcrypt generate password: %v", err)
		}
		user.Password = password
	}
	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	resp := &v1.UpdateUserResponse{
		Success: true,
	}

	return resp, nil
}

func generateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func isCorrectPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func entityToProtobuf(user *User) *v1.User {
	return &v1.User{
		Id:        user.ID,
		Uuid:      user.UUID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
