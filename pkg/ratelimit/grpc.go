// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ratelimit

import (
	"context"
	"net"
	"strconv"
	"strings"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/auth"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// GrpcKeyFunc calculates the rate limiting key from the request context and the full method name.
// Returning an empty string means that no rate-limiting should be applied.
type GrpcKeyFunc func(ctx context.Context, fullMethod string) string

// GrpcMaxWaitFunc returns the maximum duration we are allowed to wait for rate limiting
// tokens to become available.
type GrpcMaxWaitFunc func(ctx context.Context, fullMethod string) time.Duration

// GrpcRemoteIP is a GrpcKeyFunc that rate limits requests based on the remote IP address.
// Returns an empty string if request is coming from a cluster peer (TODO: Is this desired behaviour?)
func GrpcRemoteIP(ctx context.Context, fullMethod string) string {
	if md := rpcmetadata.FromIncomingContext(ctx); md.XForwardedFor != "" {
		xff := strings.Split(md.XForwardedFor, ",")
		return strings.Trim(xff[0], " ")
	}
	if p, ok := peer.FromContext(ctx); ok && p.Addr != nil && p.Addr.String() != "pipe" {
		if host, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
			return host
		}
	}
	return ""
}

// GrpcAuthID is a GrpcKeyFunc that rate limits requests based on the authentication token ID.
// Returns an empty string if no token ID is found.
func GrpcAuthID(ctx context.Context, fullMethod string) string {
	if authValue := rpcmetadata.FromIncomingContext(ctx).AuthValue; authValue != "" {
		_, id, _, err := auth.SplitToken(authValue)
		if err != nil {
			return "unauthenticated"
		}
		return id
	}
	return "unauthenticated"
}

// GrpcMaxWait is a GrpcMaxWaitFunc that allows waiting for a preset time duration.
func GrpcMaxWait(t time.Duration) GrpcMaxWaitFunc {
	return func(context.Context, string) time.Duration {
		return t
	}
}

var (
	errRateLimitExceeded = errors.DefineResourceExhausted("rate_limit_exceeded", "rate limit exceeded")
)

// GrpcUnaryServerInterceptor returns a gRPC unary server interceptor that rate limits gRPC calls.
func GrpcUnaryServerInterceptor(c Config, keyFunc GrpcKeyFunc, waitFunc GrpcMaxWaitFunc) grpc.UnaryServerInterceptor {
	l := c.New()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if key := keyFunc(ctx, info.FullMethod); key != "" {
			md, ok := l.WaitMaxDuration(key, waitFunc(ctx, info.FullMethod))
			grpc.SetHeader(ctx, metadata.Pairs(
				"x-rate-limit-limit", strconv.FormatInt(md.Limit, 10),
				"x-rate-limit-available", strconv.FormatInt(md.Available, 10),
				"x-rate-limit-reset", strconv.FormatInt(md.ResetSeconds, 10),
			))
			if !ok {
				return nil, errRateLimitExceeded.New()
			}
			time.Sleep(md.Wait)
		}
		return handler(ctx, req)
	}
}