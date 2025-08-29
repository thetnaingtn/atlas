package v1

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"atlas/internal/config"
	apiv1 "atlas/proto/gen/api/v1"
	"atlas/server/frontend"
	"atlas/store"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type APIV1Service struct {
	apiv1.UnimplementedAtlasServiceServer
	apiv1.UnimplementedProductServiceServer
	store      store.Store
	grpcServer *grpc.Server
	config     *config.Config
}

// NewAPIV1Service creates a new instance of APIV1Service
func NewAPIV1Service(grpcServer *grpc.Server, storeInstance store.Store, cfg *config.Config) *APIV1Service {
	apiService := &APIV1Service{
		store:      storeInstance,
		grpcServer: grpcServer,
		config:     cfg,
	}

	apiv1.RegisterAtlasServiceServer(grpcServer, apiService)
	apiv1.RegisterProductServiceServer(grpcServer, apiService)

	return apiService
}

func (s *APIV1Service) RegisterGateway(ctx context.Context, mux *http.ServeMux) error {
	address := fmt.Sprintf("%s:%s", s.config.Server.Addr, s.config.Server.Port)
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
	)

	if err != nil {
		return err
	}

	gwmux := runtime.NewServeMux()

	if err := apiv1.RegisterAtlasServiceHandler(ctx, gwmux, conn); err != nil {
		return err
	}
	if err := apiv1.RegisterProductServiceHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	grpcWebOptions := []grpcweb.Option{
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	}
	frontendService := frontend.NewFrontendService(&s.store, s.config)
	grpcWebProxy := grpcweb.WrapServer(s.grpcServer, grpcWebOptions...)

	handler := func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Path) >= 8 && r.URL.Path[:8] == "/api.v1." {
			fmt.Println("Handling gRPC-Web request")
			grpcWebProxy.ServeHTTP(w, r)
			return
		}

		if len(r.URL.Path) >= 4 && r.URL.Path[:4] == "/v1/" {
			fmt.Println("Handling gRPC-Gateway request")
			gwmux.ServeHTTP(w, r)
			return
		}

		frontendService.ServeHTTP(w, r)
	}

	mux.Handle("/", http.HandlerFunc(handler))

	return nil
}
