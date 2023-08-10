package user

import (
	"context"
	"errors"
	"grpc-go/domain"
	"grpc-go/helper"
	pb "grpc-go/protos/user"
	repoUser "grpc-go/repository/user"
	"log"
)

type UseCaseUser struct {
	repoUser *repoUser.RepoUser
	pb.UnimplementedUserServiceServer
}

type UseCaseUserContract interface {
	CreateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error)
	GetListUser(ctx context.Context, request *pb.GetUserListRequest) (*pb.UserListResponse, error)
	UpdateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error)
	GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.UserResponse, error)
	DeleteUser(ctx context.Context, request *pb.GetUserRequest) (*pb.DeleteUserResponse, error)
}

func NewUserUseCase(repoUser *repoUser.RepoUser) *UseCaseUser {
	return &UseCaseUser{
		repoUser: repoUser,
	}
}

func (u *UseCaseUser) CreateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	reqUser := &domain.User{
		Email:    request.GetEmail(),
		Password: helper.GetMD5Hash(request.GetPassword()),
		Name:     request.GetUserName(),
	}

	log.Println(request.GetEmail())
	err := u.repoUser.CreateUser(ctx, reqUser)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		ID:       reqUser.ID,
		Email:    reqUser.Email,
		UserName: reqUser.Name,
	}, nil
}

func (u *UseCaseUser) GetListUser(ctx context.Context, request *pb.GetUserListRequest) (*pb.UserListResponse, error) {
	offset := request.GetPageNumber() - 1*request.GetPageSize()
	limit := request.GetPageSize()
	res, total, err := u.repoUser.GetListUser(ctx, int(limit), int(offset))
	if err != nil {
		return nil, err
	}

	totalPage := int32(total) / limit
	meta := &pb.PageInfo{
		TotalItems:  int32(total),
		TotalPages:  totalPage,
		CurrentPage: request.GetPageSize(),
	}

	return &pb.UserListResponse{
		Users:    mapUsersToResponseUsers(res),
		PageInfo: meta,
	}, nil
}

func (u *UseCaseUser) UpdateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	reqUser := &domain.User{
		Email:    request.GetEmail(),
		Password: helper.GetMD5Hash(request.GetPassword()),
		Name:     request.GetUserName(),
	}

	user, err := u.repoUser.GetUserByEmail(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user with given email not found")
	}

	err = u.repoUser.UpdateUser(ctx, reqUser)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		ID:       user.ID,
		Email:    reqUser.Email,
		UserName: reqUser.Name,
	}, nil
}

func (u *UseCaseUser) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := u.repoUser.GetUser(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.Name,
	}, nil
}

func (u *UseCaseUser) DeleteUser(ctx context.Context, request *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {

	user, err := u.repoUser.GetUser(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user with given email not found")
	}

	err = u.repoUser.DeleteUser(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{Message: "Success"}, nil
}

func mapUsersToResponseUsers(e []domain.User) []*pb.User {
	var res []*pb.User
	for _, v := range e {
		res = append(res, &pb.User{
			ID:       v.ID,
			Email:    v.Email,
			UserName: v.Name,
		})
	}
	return res
}
