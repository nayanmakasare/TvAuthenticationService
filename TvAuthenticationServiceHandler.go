package main

import (
	"context"
	"github.com/go-redis/redis"
	TvAuthenticationService "github.com/nayanmakasare/TvAuthenticationService/proto"
)

type TvAuthenticationServiceHandler struct {
	AuthRedisConnection *redis.Client
}

func(tvs *TvAuthenticationServiceHandler) CloudwalkerAuthWall(ctx context.Context, targetTv *TvAuthenticationService.TargetTv, isAuthorised *TvAuthenticationService.IsAuthorised) error{
	result := tvs.AuthRedisConnection.SIsMember("cloudwalkerEmac", targetTv.Emac)
	if result.Err() != nil {
		return  result.Err()
	}
	isAuthorised.Result = result.Val()
	return nil
}
