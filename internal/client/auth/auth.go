package auth

import (
	"context"
	"fmt"

	"github.com/solumD/auth/pkg/access_v1"
	"github.com/solumD/auth/pkg/auth_v1"
	"github.com/solumD/auth/pkg/user_v1"
	"github.com/solumD/chat-client/internal/client"
	"github.com/solumD/chat-client/internal/model"

	"google.golang.org/grpc/metadata"
)

type authServerClient struct {
	userClient   user_v1.UserV1Client
	authClient   auth_v1.AuthV1Client
	accessClient access_v1.AccessV1Client
}

func New(userClient user_v1.UserV1Client, authClient auth_v1.AuthV1Client, accessClient access_v1.AccessV1Client) client.AuthServerClient {
	return &authServerClient{
		userClient:   userClient,
		authClient:   authClient,
		accessClient: accessClient,
	}
}

func (cl *authServerClient) CreateUser(ctx context.Context, user *model.UserToCreate) (int64, error) {
	req := &user_v1.CreateUserRequest{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
	}

	res, err := cl.userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func (cl *authServerClient) Login(ctx context.Context, user *model.UserToLogin) (string, string, error) {
	req := &auth_v1.LoginRequest{
		Username: user.Name,
		Password: user.Password,
	}

	res, err := cl.authClient.Login(ctx, req)
	if err != nil {
		return "", "", err
	}

	return res.GetRefreshToken(), res.GetAccessToken(), nil
}

func (cl *authServerClient) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	req := &auth_v1.GetRefreshTokenRequest{
		OldRefreshToken: refreshToken,
	}

	res, err := cl.authClient.GetRefreshToken(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetRefreshToken(), nil
}

func (cl *authServerClient) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	req := &auth_v1.GetAccessTokenRequest{
		RefreshToken: refreshToken,
	}

	res, err := cl.authClient.GetAccessToken(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

func (cl *authServerClient) Check(ctx context.Context, accessToken string, endpoint string) (string, error) {
	req := &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	}

	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", accessToken)})
	mdCtx := metadata.NewOutgoingContext(ctx, md)

	res, err := cl.accessClient.Check(mdCtx, req)
	if err != nil {
		return "", err
	}

	return res.GetUsername(), nil
}
