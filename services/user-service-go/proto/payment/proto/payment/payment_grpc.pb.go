// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/payment/payment.proto

package paymentpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_CreatePayment_FullMethodName           = "/payment.v1.PaymentService/CreatePayment"
	PaymentService_GetPayment_FullMethodName              = "/payment.v1.PaymentService/GetPayment"
	PaymentService_GetUserPayments_FullMethodName         = "/payment.v1.PaymentService/GetUserPayments"
	PaymentService_GetOrderPayments_FullMethodName        = "/payment.v1.PaymentService/GetOrderPayments"
	PaymentService_ProcessPayment_FullMethodName          = "/payment.v1.PaymentService/ProcessPayment"
	PaymentService_ConfirmPayment_FullMethodName          = "/payment.v1.PaymentService/ConfirmPayment"
	PaymentService_CancelPayment_FullMethodName           = "/payment.v1.PaymentService/CancelPayment"
	PaymentService_RefundPayment_FullMethodName           = "/payment.v1.PaymentService/RefundPayment"
	PaymentService_UpdatePaymentStatus_FullMethodName     = "/payment.v1.PaymentService/UpdatePaymentStatus"
	PaymentService_ValidatePaymentProvider_FullMethodName = "/payment.v1.PaymentService/ValidatePaymentProvider"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Payment service definition for internal microservice communication
type PaymentServiceClient interface {
	// Create a new payment
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	// Get payment by ID
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
	// Get payments by user ID
	GetUserPayments(ctx context.Context, in *GetUserPaymentsRequest, opts ...grpc.CallOption) (*GetUserPaymentsResponse, error)
	// Get payments by order ID
	GetOrderPayments(ctx context.Context, in *GetOrderPaymentsRequest, opts ...grpc.CallOption) (*GetOrderPaymentsResponse, error)
	// Process payment (initiate payment with provider)
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
	// Confirm payment (mark as completed)
	ConfirmPayment(ctx context.Context, in *ConfirmPaymentRequest, opts ...grpc.CallOption) (*ConfirmPaymentResponse, error)
	// Cancel payment
	CancelPayment(ctx context.Context, in *CancelPaymentRequest, opts ...grpc.CallOption) (*CancelPaymentResponse, error)
	// Refund payment
	RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error)
	// Update payment status (webhook handling)
	UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error)
	// Validate payment provider
	ValidatePaymentProvider(ctx context.Context, in *ValidatePaymentProviderRequest, opts ...grpc.CallOption) (*ValidatePaymentProviderResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetUserPayments(ctx context.Context, in *GetUserPaymentsRequest, opts ...grpc.CallOption) (*GetUserPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserPaymentsResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetUserPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetOrderPayments(ctx context.Context, in *GetOrderPaymentsRequest, opts ...grpc.CallOption) (*GetOrderPaymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderPaymentsResponse)
	err := c.cc.Invoke(ctx, PaymentService_GetOrderPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_ProcessPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ConfirmPayment(ctx context.Context, in *ConfirmPaymentRequest, opts ...grpc.CallOption) (*ConfirmPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_ConfirmPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CancelPayment(ctx context.Context, in *CancelPaymentRequest, opts ...grpc.CallOption) (*CancelPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_CancelPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) RefundPayment(ctx context.Context, in *RefundPaymentRequest, opts ...grpc.CallOption) (*RefundPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefundPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_RefundPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdatePaymentStatus(ctx context.Context, in *UpdatePaymentStatusRequest, opts ...grpc.CallOption) (*UpdatePaymentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentStatusResponse)
	err := c.cc.Invoke(ctx, PaymentService_UpdatePaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) ValidatePaymentProvider(ctx context.Context, in *ValidatePaymentProviderRequest, opts ...grpc.CallOption) (*ValidatePaymentProviderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidatePaymentProviderResponse)
	err := c.cc.Invoke(ctx, PaymentService_ValidatePaymentProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
//
// Payment service definition for internal microservice communication
type PaymentServiceServer interface {
	// Create a new payment
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	// Get payment by ID
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	// Get payments by user ID
	GetUserPayments(context.Context, *GetUserPaymentsRequest) (*GetUserPaymentsResponse, error)
	// Get payments by order ID
	GetOrderPayments(context.Context, *GetOrderPaymentsRequest) (*GetOrderPaymentsResponse, error)
	// Process payment (initiate payment with provider)
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	// Confirm payment (mark as completed)
	ConfirmPayment(context.Context, *ConfirmPaymentRequest) (*ConfirmPaymentResponse, error)
	// Cancel payment
	CancelPayment(context.Context, *CancelPaymentRequest) (*CancelPaymentResponse, error)
	// Refund payment
	RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error)
	// Update payment status (webhook handling)
	UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error)
	// Validate payment provider
	ValidatePaymentProvider(context.Context, *ValidatePaymentProviderRequest) (*ValidatePaymentProviderResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetUserPayments(context.Context, *GetUserPaymentsRequest) (*GetUserPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPayments not implemented")
}
func (UnimplementedPaymentServiceServer) GetOrderPayments(context.Context, *GetOrderPaymentsRequest) (*GetOrderPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderPayments not implemented")
}
func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServiceServer) ConfirmPayment(context.Context, *ConfirmPaymentRequest) (*ConfirmPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPayment not implemented")
}
func (UnimplementedPaymentServiceServer) CancelPayment(context.Context, *CancelPaymentRequest) (*CancelPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPayment not implemented")
}
func (UnimplementedPaymentServiceServer) RefundPayment(context.Context, *RefundPaymentRequest) (*RefundPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundPayment not implemented")
}
func (UnimplementedPaymentServiceServer) UpdatePaymentStatus(context.Context, *UpdatePaymentStatusRequest) (*UpdatePaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentStatus not implemented")
}
func (UnimplementedPaymentServiceServer) ValidatePaymentProvider(context.Context, *ValidatePaymentProviderRequest) (*ValidatePaymentProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePaymentProvider not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetUserPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetUserPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetUserPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetUserPayments(ctx, req.(*GetUserPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetOrderPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetOrderPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetOrderPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetOrderPayments(ctx, req.(*GetOrderPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ConfirmPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ConfirmPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ConfirmPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ConfirmPayment(ctx, req.(*ConfirmPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CancelPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CancelPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CancelPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CancelPayment(ctx, req.(*CancelPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_RefundPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).RefundPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_RefundPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).RefundPayment(ctx, req.(*RefundPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_UpdatePaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_UpdatePaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).UpdatePaymentStatus(ctx, req.(*UpdatePaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_ValidatePaymentProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePaymentProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ValidatePaymentProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ValidatePaymentProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ValidatePaymentProvider(ctx, req.(*ValidatePaymentProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.v1.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _PaymentService_GetPayment_Handler,
		},
		{
			MethodName: "GetUserPayments",
			Handler:    _PaymentService_GetUserPayments_Handler,
		},
		{
			MethodName: "GetOrderPayments",
			Handler:    _PaymentService_GetOrderPayments_Handler,
		},
		{
			MethodName: "ProcessPayment",
			Handler:    _PaymentService_ProcessPayment_Handler,
		},
		{
			MethodName: "ConfirmPayment",
			Handler:    _PaymentService_ConfirmPayment_Handler,
		},
		{
			MethodName: "CancelPayment",
			Handler:    _PaymentService_CancelPayment_Handler,
		},
		{
			MethodName: "RefundPayment",
			Handler:    _PaymentService_RefundPayment_Handler,
		},
		{
			MethodName: "UpdatePaymentStatus",
			Handler:    _PaymentService_UpdatePaymentStatus_Handler,
		},
		{
			MethodName: "ValidatePaymentProvider",
			Handler:    _PaymentService_ValidatePaymentProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payment/payment.proto",
}
