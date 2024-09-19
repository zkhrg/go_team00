package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	DataStream_StreamData_FullMethodName = "/datastream.DataStream/StreamData"
)

type DataStreamClient interface {
	StreamData(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataMessage], error)
}

type dataStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewDataStreamClient(cc grpc.ClientConnInterface) DataStreamClient {
	return &dataStreamClient{cc}
}

func (c *dataStreamClient) StreamData(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DataMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataStream_ServiceDesc.Streams[0], DataStream_StreamData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRequest, DataMessage]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataStream_StreamDataClient = grpc.ServerStreamingClient[DataMessage]

type DataStreamServer interface {
	StreamData(*StreamRequest, grpc.ServerStreamingServer[DataMessage]) error
	mustEmbedUnimplementedDataStreamServer()
}

type UnimplementedDataStreamServer struct{}

func (UnimplementedDataStreamServer) StreamData(*StreamRequest, grpc.ServerStreamingServer[DataMessage]) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedDataStreamServer) mustEmbedUnimplementedDataStreamServer() {}
func (UnimplementedDataStreamServer) testEmbeddedByValue()                    {}

type UnsafeDataStreamServer interface {
	mustEmbedUnimplementedDataStreamServer()
}

func RegisterDataStreamServer(s grpc.ServiceRegistrar, srv DataStreamServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataStream_ServiceDesc, srv)
}

func _DataStream_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataStreamServer).StreamData(m, &grpc.GenericServerStream[StreamRequest, DataMessage]{ServerStream: stream})
}

type DataStream_StreamDataServer = grpc.ServerStreamingServer[DataMessage]

var DataStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datastream.DataStream",
	HandlerType: (*DataStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamData",
			Handler:       _DataStream_StreamData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/api/data_stream.proto",
}
