package blog

import (
	"context"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/blog/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	userv1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/user/v1"
	authv1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var prefix = "/" + v1.BlogService_ServiceDesc.ServiceName + "/"

var AuthMethod = map[string]bool{
	prefix + "SignUp":               false, //不需要验证token
	prefix + "SignIn":               false,
	prefix + "CreatePost":           true, // 需要验证
	prefix + "UpdatePost":           true,
	prefix + "GetPost":              false,
	prefix + "ListPosts":            false,
	prefix + "DeletePost":           true,
	prefix + "CreateComment":        true,
	prefix + "UpdateComment":        true,
	prefix + "DeleteComment":        true,
	prefix + "ListCommentsByPostID": false,
}

func NewServer(logger *log.Logger, conf *config.Config, authClient authv1.AuthServiceClient, userClient userv1.UserServiceClient) v1.BlogServiceServer {
	return &Server{
		logger:     logger,
		conf:       conf,
		authClient: authClient,
		userClient: userClient,
	};
}

type Server struct {
	v1.UnimplementedBlogServiceServer
	logger     *log.Logger
	conf       *config.Config
	userClient userv1.UserServiceClient
	authClient authv1.AuthServiceClient
}

func (s Server) SignUp(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	email := req.GetEmail()

	usernameResp, err := s.userClient.GetUserByUsername(ctx, &userv1.GetUserByUsernameRequest{
		Username: username,
	})
	s.logger.Info("usernameResp:", usernameResp)
	if err == nil && usernameResp.GetUser().GetId() != 0 {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}
	emailResp, err := s.userClient.GetUserByEmail(ctx, &userv1.GetUserByEmailRequest{
		Email: email,
	})
	if err == nil && emailResp.GetUser().GetId() != 0 {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	userResp, err := s.userClient.CreateUser(ctx, &userv1.CreateUserRequest{
		User: &userv1.User{
			Username: username,
			Email:    email,
			Password: password,
		},
	})
	if err != nil {
		s.logger.Error(err)
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	authResp, err := s.authClient.GenerateToken(ctx, &authv1.GenerateTokenRequest{
		UserId: userResp.GetUser().GetId(),
	})

	if err != nil {
		s.logger.Error(err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}
	return &v1.SignUpResponse{Token: authResp.GetToken()}, nil
}

func (s Server) SignIn(ctx context.Context, request *v1.SignInRequest) (*v1.SignInResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	email := request.GetEmail()

	if username == "" && email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email and username cannot be empty");
	}
	var userID uint64

	if email != "" {
		resp, err := s.userClient.GetUserByEmail(ctx, &userv1.GetUserByEmailRequest{
			Email:    email,
			Password: password,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user by email: %v", err)
		}
		user := resp.GetUser()
		userID = user.GetId()
	} else {
		req, err := s.userClient.GetUserByUsername(ctx, &userv1.GetUserByUsernameRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user by username: %v", err)
		}
		user := req.GetUser()
		userID = user.GetId()
	}

	authResp, err := s.authClient.GenerateToken(ctx, &authv1.GenerateTokenRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}
	return &v1.SignInResponse{Token: authResp.GetToken()}, nil

}
func (s Server) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	return &v1.CreatePostResponse{}, nil
}

func (s Server) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	return &v1.GetPostResponse{}, nil
}
func (s Server) ListPosts(ctx context.Context, req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	return &v1.ListPostsResponse{}, nil
}
func (s Server) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	return &v1.UpdatePostResponse{}, nil
}

func (s Server) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	return &v1.DeletePostResponse{}, nil
}

func (s Server) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	return &v1.CreateCommentResponse{}, nil
}
func (s Server) UpdateComment(ctx context.Context, req *v1.UpdateCommentRequest) (*v1.UpdateCommentResponse, error) {
	return &v1.UpdateCommentResponse{}, nil
}

func (s Server) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	return &v1.DeleteCommentResponse{}, nil
}

func (s Server) ListCommentsByPostID(ctx context.Context, req *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	return &v1.ListCommentsByPostIDResponse{}, nil
}
