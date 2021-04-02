// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/matchfunction.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_MatchFunction_Run_0(ctx context.Context, marshaler runtime.Marshaler, client MatchFunctionClient, req *http.Request, pathParams map[string]string) (MatchFunction_RunClient, runtime.ServerMetadata, error) {
	var protoReq RunRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.Run(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterMatchFunctionHandlerServer registers the http handlers for service MatchFunction to "mux".
// UnaryRPC     :call MatchFunctionServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMatchFunctionHandlerFromEndpoint instead.
func RegisterMatchFunctionHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MatchFunctionServer) error {

	mux.Handle("POST", pattern_MatchFunction_Run_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterMatchFunctionHandlerFromEndpoint is same as RegisterMatchFunctionHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMatchFunctionHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMatchFunctionHandler(ctx, mux, conn)
}

// RegisterMatchFunctionHandler registers the http handlers for service MatchFunction to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMatchFunctionHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMatchFunctionHandlerClient(ctx, mux, NewMatchFunctionClient(conn))
}

// RegisterMatchFunctionHandlerClient registers the http handlers for service MatchFunction
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MatchFunctionClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MatchFunctionClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MatchFunctionClient" to call the correct interceptors.
func RegisterMatchFunctionHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MatchFunctionClient) error {

	mux.Handle("POST", pattern_MatchFunction_Run_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MatchFunction_Run_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MatchFunction_Run_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MatchFunction_Run_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "matchfunction"}, "run", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_MatchFunction_Run_0 = runtime.ForwardResponseStream
)
