package api

import (
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/darwishdev/devkit-api/config"
	"github.com/darwishdev/devkit-api/gen/db"
	"github.com/darwishdev/devkit-api/gen/pb/proto/devkit/v1/devkitv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"time"
)

type Server struct {
	config config.Config
	store  db.Store
	api    devkitv1connect.DevkitServiceHandler
}

func NewServer(config config.Config, store db.Store) *Server {
	return &Server{
		config: config,
		store:  store,
		api:    NewApi(config, store),
	}
}

func (s Server) NewGrpcHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler("https://darwishdev.com", http.StatusFound))
	// here we can find examples of diffrent compression method 	https://connectrpc.com/docs/go/serialization-and-compression/#compression
	compress1KB := connect.WithCompressMinBytes(1024)
	interceptors := connect.WithInterceptors(GrpcLogger())

	mux.Handle(devkitv1connect.NewDevkitServiceHandler(
		s.api,
		interceptors,
		compress1KB,
	))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(devkitv1connect.DevkitServiceName),
		compress1KB,
	))
	cors := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
	server := &http.Server{
		Addr:    s.config.GRPCServerAddress,
		Handler: h2c.NewHandler(cors.Handler(mux), &http2.Server{}),
	}
	return server

}