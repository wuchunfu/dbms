// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: dbms_master.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	UpsertDatabase(ctx context.Context, in *UpsertDatabaseRequest, opts ...grpc.CallOption) (*UpsertDatabaseResponse, error)
	DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*DeleteDatabaseResponse, error)
	ShowDatabase(ctx context.Context, in *ShowDatabaseRequest, opts ...grpc.CallOption) (*ShowDatabaseResponse, error)
	UpsertDatasource(ctx context.Context, in *UpsertDatasourceRequest, opts ...grpc.CallOption) (*UpsertDatasourceResponse, error)
	DeleteDatasource(ctx context.Context, in *DeleteDatasourceRequest, opts ...grpc.CallOption) (*DeleteDatasourceResponse, error)
	ShowDatasource(ctx context.Context, in *ShowDatasourceRequest, opts ...grpc.CallOption) (*ShowDatasourceResponse, error)
	UpsertStructMigrateTask(ctx context.Context, in *UpsertStructMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertStructMigrateTaskResponse, error)
	DeleteStructMigrateTask(ctx context.Context, in *DeleteStructMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteStructMigrateTaskResponse, error)
	ShowStructMigrateTask(ctx context.Context, in *ShowStructMigrateTaskRequest, opts ...grpc.CallOption) (*ShowStructMigrateTaskResponse, error)
	UpsertStmtMigrateTask(ctx context.Context, in *UpsertStmtMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertStmtMigrateTaskResponse, error)
	DeleteStmtMigrateTask(ctx context.Context, in *DeleteStmtMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteStmtMigrateTaskResponse, error)
	ShowStmtMigrateTask(ctx context.Context, in *ShowStmtMigrateTaskRequest, opts ...grpc.CallOption) (*ShowStmtMigrateTaskResponse, error)
	UpsertSqlMigrateTask(ctx context.Context, in *UpsertSqlMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertSqlMigrateTaskResponse, error)
	DeleteSqlMigrateTask(ctx context.Context, in *DeleteSqlMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteSqlMigrateTaskResponse, error)
	ShowSqlMigrateTask(ctx context.Context, in *ShowSqlMigrateTaskRequest, opts ...grpc.CallOption) (*ShowSqlMigrateTaskResponse, error)
	UpsertCsvMigrateTask(ctx context.Context, in *UpsertCsvMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertCsvMigrateTaskResponse, error)
	DeleteCsvMigrateTask(ctx context.Context, in *DeleteCsvMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteCsvMigrateTaskResponse, error)
	ShowCsvMigrateTask(ctx context.Context, in *ShowCsvMigrateTaskRequest, opts ...grpc.CallOption) (*ShowCsvMigrateTaskResponse, error)
	UpsertDataCompareTask(ctx context.Context, in *UpsertDataCompareTaskRequest, opts ...grpc.CallOption) (*UpsertDataCompareTaskResponse, error)
	DeleteDataCompareTask(ctx context.Context, in *DeleteDataCompareTaskRequest, opts ...grpc.CallOption) (*DeleteDataCompareTaskResponse, error)
	ShowDataCompareTask(ctx context.Context, in *ShowDataCompareTaskRequest, opts ...grpc.CallOption) (*ShowDataCompareTaskResponse, error)
	OperateTask(ctx context.Context, in *OperateTaskRequest, opts ...grpc.CallOption) (*OperateTaskResponse, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) UpsertDatabase(ctx context.Context, in *UpsertDatabaseRequest, opts ...grpc.CallOption) (*UpsertDatabaseResponse, error) {
	out := new(UpsertDatabaseResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*DeleteDatabaseResponse, error) {
	out := new(DeleteDatabaseResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowDatabase(ctx context.Context, in *ShowDatabaseRequest, opts ...grpc.CallOption) (*ShowDatabaseResponse, error) {
	out := new(ShowDatabaseResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertDatasource(ctx context.Context, in *UpsertDatasourceRequest, opts ...grpc.CallOption) (*UpsertDatasourceResponse, error) {
	out := new(UpsertDatasourceResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteDatasource(ctx context.Context, in *DeleteDatasourceRequest, opts ...grpc.CallOption) (*DeleteDatasourceResponse, error) {
	out := new(DeleteDatasourceResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowDatasource(ctx context.Context, in *ShowDatasourceRequest, opts ...grpc.CallOption) (*ShowDatasourceResponse, error) {
	out := new(ShowDatasourceResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertStructMigrateTask(ctx context.Context, in *UpsertStructMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertStructMigrateTaskResponse, error) {
	out := new(UpsertStructMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertStructMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteStructMigrateTask(ctx context.Context, in *DeleteStructMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteStructMigrateTaskResponse, error) {
	out := new(DeleteStructMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteStructMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowStructMigrateTask(ctx context.Context, in *ShowStructMigrateTaskRequest, opts ...grpc.CallOption) (*ShowStructMigrateTaskResponse, error) {
	out := new(ShowStructMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowStructMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertStmtMigrateTask(ctx context.Context, in *UpsertStmtMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertStmtMigrateTaskResponse, error) {
	out := new(UpsertStmtMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertStmtMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteStmtMigrateTask(ctx context.Context, in *DeleteStmtMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteStmtMigrateTaskResponse, error) {
	out := new(DeleteStmtMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteStmtMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowStmtMigrateTask(ctx context.Context, in *ShowStmtMigrateTaskRequest, opts ...grpc.CallOption) (*ShowStmtMigrateTaskResponse, error) {
	out := new(ShowStmtMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowStmtMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertSqlMigrateTask(ctx context.Context, in *UpsertSqlMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertSqlMigrateTaskResponse, error) {
	out := new(UpsertSqlMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertSqlMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteSqlMigrateTask(ctx context.Context, in *DeleteSqlMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteSqlMigrateTaskResponse, error) {
	out := new(DeleteSqlMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteSqlMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowSqlMigrateTask(ctx context.Context, in *ShowSqlMigrateTaskRequest, opts ...grpc.CallOption) (*ShowSqlMigrateTaskResponse, error) {
	out := new(ShowSqlMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowSqlMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertCsvMigrateTask(ctx context.Context, in *UpsertCsvMigrateTaskRequest, opts ...grpc.CallOption) (*UpsertCsvMigrateTaskResponse, error) {
	out := new(UpsertCsvMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertCsvMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteCsvMigrateTask(ctx context.Context, in *DeleteCsvMigrateTaskRequest, opts ...grpc.CallOption) (*DeleteCsvMigrateTaskResponse, error) {
	out := new(DeleteCsvMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteCsvMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowCsvMigrateTask(ctx context.Context, in *ShowCsvMigrateTaskRequest, opts ...grpc.CallOption) (*ShowCsvMigrateTaskResponse, error) {
	out := new(ShowCsvMigrateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowCsvMigrateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpsertDataCompareTask(ctx context.Context, in *UpsertDataCompareTaskRequest, opts ...grpc.CallOption) (*UpsertDataCompareTaskResponse, error) {
	out := new(UpsertDataCompareTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/UpsertDataCompareTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) DeleteDataCompareTask(ctx context.Context, in *DeleteDataCompareTaskRequest, opts ...grpc.CallOption) (*DeleteDataCompareTaskResponse, error) {
	out := new(DeleteDataCompareTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/DeleteDataCompareTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ShowDataCompareTask(ctx context.Context, in *ShowDataCompareTaskRequest, opts ...grpc.CallOption) (*ShowDataCompareTaskResponse, error) {
	out := new(ShowDataCompareTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/ShowDataCompareTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) OperateTask(ctx context.Context, in *OperateTaskRequest, opts ...grpc.CallOption) (*OperateTaskResponse, error) {
	out := new(OperateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.Master/OperateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	UpsertDatabase(context.Context, *UpsertDatabaseRequest) (*UpsertDatabaseResponse, error)
	DeleteDatabase(context.Context, *DeleteDatabaseRequest) (*DeleteDatabaseResponse, error)
	ShowDatabase(context.Context, *ShowDatabaseRequest) (*ShowDatabaseResponse, error)
	UpsertDatasource(context.Context, *UpsertDatasourceRequest) (*UpsertDatasourceResponse, error)
	DeleteDatasource(context.Context, *DeleteDatasourceRequest) (*DeleteDatasourceResponse, error)
	ShowDatasource(context.Context, *ShowDatasourceRequest) (*ShowDatasourceResponse, error)
	UpsertStructMigrateTask(context.Context, *UpsertStructMigrateTaskRequest) (*UpsertStructMigrateTaskResponse, error)
	DeleteStructMigrateTask(context.Context, *DeleteStructMigrateTaskRequest) (*DeleteStructMigrateTaskResponse, error)
	ShowStructMigrateTask(context.Context, *ShowStructMigrateTaskRequest) (*ShowStructMigrateTaskResponse, error)
	UpsertStmtMigrateTask(context.Context, *UpsertStmtMigrateTaskRequest) (*UpsertStmtMigrateTaskResponse, error)
	DeleteStmtMigrateTask(context.Context, *DeleteStmtMigrateTaskRequest) (*DeleteStmtMigrateTaskResponse, error)
	ShowStmtMigrateTask(context.Context, *ShowStmtMigrateTaskRequest) (*ShowStmtMigrateTaskResponse, error)
	UpsertSqlMigrateTask(context.Context, *UpsertSqlMigrateTaskRequest) (*UpsertSqlMigrateTaskResponse, error)
	DeleteSqlMigrateTask(context.Context, *DeleteSqlMigrateTaskRequest) (*DeleteSqlMigrateTaskResponse, error)
	ShowSqlMigrateTask(context.Context, *ShowSqlMigrateTaskRequest) (*ShowSqlMigrateTaskResponse, error)
	UpsertCsvMigrateTask(context.Context, *UpsertCsvMigrateTaskRequest) (*UpsertCsvMigrateTaskResponse, error)
	DeleteCsvMigrateTask(context.Context, *DeleteCsvMigrateTaskRequest) (*DeleteCsvMigrateTaskResponse, error)
	ShowCsvMigrateTask(context.Context, *ShowCsvMigrateTaskRequest) (*ShowCsvMigrateTaskResponse, error)
	UpsertDataCompareTask(context.Context, *UpsertDataCompareTaskRequest) (*UpsertDataCompareTaskResponse, error)
	DeleteDataCompareTask(context.Context, *DeleteDataCompareTaskRequest) (*DeleteDataCompareTaskResponse, error)
	ShowDataCompareTask(context.Context, *ShowDataCompareTaskRequest) (*ShowDataCompareTaskResponse, error)
	OperateTask(context.Context, *OperateTaskRequest) (*OperateTaskResponse, error)
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) UpsertDatabase(context.Context, *UpsertDatabaseRequest) (*UpsertDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDatabase not implemented")
}
func (UnimplementedMasterServer) DeleteDatabase(context.Context, *DeleteDatabaseRequest) (*DeleteDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatabase not implemented")
}
func (UnimplementedMasterServer) ShowDatabase(context.Context, *ShowDatabaseRequest) (*ShowDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowDatabase not implemented")
}
func (UnimplementedMasterServer) UpsertDatasource(context.Context, *UpsertDatasourceRequest) (*UpsertDatasourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDatasource not implemented")
}
func (UnimplementedMasterServer) DeleteDatasource(context.Context, *DeleteDatasourceRequest) (*DeleteDatasourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatasource not implemented")
}
func (UnimplementedMasterServer) ShowDatasource(context.Context, *ShowDatasourceRequest) (*ShowDatasourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowDatasource not implemented")
}
func (UnimplementedMasterServer) UpsertStructMigrateTask(context.Context, *UpsertStructMigrateTaskRequest) (*UpsertStructMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertStructMigrateTask not implemented")
}
func (UnimplementedMasterServer) DeleteStructMigrateTask(context.Context, *DeleteStructMigrateTaskRequest) (*DeleteStructMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStructMigrateTask not implemented")
}
func (UnimplementedMasterServer) ShowStructMigrateTask(context.Context, *ShowStructMigrateTaskRequest) (*ShowStructMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowStructMigrateTask not implemented")
}
func (UnimplementedMasterServer) UpsertStmtMigrateTask(context.Context, *UpsertStmtMigrateTaskRequest) (*UpsertStmtMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertStmtMigrateTask not implemented")
}
func (UnimplementedMasterServer) DeleteStmtMigrateTask(context.Context, *DeleteStmtMigrateTaskRequest) (*DeleteStmtMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStmtMigrateTask not implemented")
}
func (UnimplementedMasterServer) ShowStmtMigrateTask(context.Context, *ShowStmtMigrateTaskRequest) (*ShowStmtMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowStmtMigrateTask not implemented")
}
func (UnimplementedMasterServer) UpsertSqlMigrateTask(context.Context, *UpsertSqlMigrateTaskRequest) (*UpsertSqlMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertSqlMigrateTask not implemented")
}
func (UnimplementedMasterServer) DeleteSqlMigrateTask(context.Context, *DeleteSqlMigrateTaskRequest) (*DeleteSqlMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSqlMigrateTask not implemented")
}
func (UnimplementedMasterServer) ShowSqlMigrateTask(context.Context, *ShowSqlMigrateTaskRequest) (*ShowSqlMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowSqlMigrateTask not implemented")
}
func (UnimplementedMasterServer) UpsertCsvMigrateTask(context.Context, *UpsertCsvMigrateTaskRequest) (*UpsertCsvMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCsvMigrateTask not implemented")
}
func (UnimplementedMasterServer) DeleteCsvMigrateTask(context.Context, *DeleteCsvMigrateTaskRequest) (*DeleteCsvMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCsvMigrateTask not implemented")
}
func (UnimplementedMasterServer) ShowCsvMigrateTask(context.Context, *ShowCsvMigrateTaskRequest) (*ShowCsvMigrateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCsvMigrateTask not implemented")
}
func (UnimplementedMasterServer) UpsertDataCompareTask(context.Context, *UpsertDataCompareTaskRequest) (*UpsertDataCompareTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDataCompareTask not implemented")
}
func (UnimplementedMasterServer) DeleteDataCompareTask(context.Context, *DeleteDataCompareTaskRequest) (*DeleteDataCompareTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataCompareTask not implemented")
}
func (UnimplementedMasterServer) ShowDataCompareTask(context.Context, *ShowDataCompareTaskRequest) (*ShowDataCompareTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowDataCompareTask not implemented")
}
func (UnimplementedMasterServer) OperateTask(context.Context, *OperateTaskRequest) (*OperateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateTask not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_UpsertDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertDatabase(ctx, req.(*UpsertDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteDatabase(ctx, req.(*DeleteDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowDatabase(ctx, req.(*ShowDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDatasourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertDatasource(ctx, req.(*UpsertDatasourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteDatasource(ctx, req.(*DeleteDatasourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowDatasourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowDatasource(ctx, req.(*ShowDatasourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertStructMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertStructMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertStructMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertStructMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertStructMigrateTask(ctx, req.(*UpsertStructMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteStructMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStructMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteStructMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteStructMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteStructMigrateTask(ctx, req.(*DeleteStructMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowStructMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowStructMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowStructMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowStructMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowStructMigrateTask(ctx, req.(*ShowStructMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertStmtMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertStmtMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertStmtMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertStmtMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertStmtMigrateTask(ctx, req.(*UpsertStmtMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteStmtMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStmtMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteStmtMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteStmtMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteStmtMigrateTask(ctx, req.(*DeleteStmtMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowStmtMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowStmtMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowStmtMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowStmtMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowStmtMigrateTask(ctx, req.(*ShowStmtMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertSqlMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertSqlMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertSqlMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertSqlMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertSqlMigrateTask(ctx, req.(*UpsertSqlMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteSqlMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSqlMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteSqlMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteSqlMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteSqlMigrateTask(ctx, req.(*DeleteSqlMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowSqlMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowSqlMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowSqlMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowSqlMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowSqlMigrateTask(ctx, req.(*ShowSqlMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertCsvMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCsvMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertCsvMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertCsvMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertCsvMigrateTask(ctx, req.(*UpsertCsvMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteCsvMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCsvMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteCsvMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteCsvMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteCsvMigrateTask(ctx, req.(*DeleteCsvMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowCsvMigrateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowCsvMigrateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowCsvMigrateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowCsvMigrateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowCsvMigrateTask(ctx, req.(*ShowCsvMigrateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpsertDataCompareTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDataCompareTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpsertDataCompareTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/UpsertDataCompareTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpsertDataCompareTask(ctx, req.(*UpsertDataCompareTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_DeleteDataCompareTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataCompareTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).DeleteDataCompareTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/DeleteDataCompareTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).DeleteDataCompareTask(ctx, req.(*DeleteDataCompareTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ShowDataCompareTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowDataCompareTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ShowDataCompareTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/ShowDataCompareTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ShowDataCompareTask(ctx, req.(*ShowDataCompareTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_OperateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).OperateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Master/OperateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).OperateTask(ctx, req.(*OperateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertDatabase",
			Handler:    _Master_UpsertDatabase_Handler,
		},
		{
			MethodName: "DeleteDatabase",
			Handler:    _Master_DeleteDatabase_Handler,
		},
		{
			MethodName: "ShowDatabase",
			Handler:    _Master_ShowDatabase_Handler,
		},
		{
			MethodName: "UpsertDatasource",
			Handler:    _Master_UpsertDatasource_Handler,
		},
		{
			MethodName: "DeleteDatasource",
			Handler:    _Master_DeleteDatasource_Handler,
		},
		{
			MethodName: "ShowDatasource",
			Handler:    _Master_ShowDatasource_Handler,
		},
		{
			MethodName: "UpsertStructMigrateTask",
			Handler:    _Master_UpsertStructMigrateTask_Handler,
		},
		{
			MethodName: "DeleteStructMigrateTask",
			Handler:    _Master_DeleteStructMigrateTask_Handler,
		},
		{
			MethodName: "ShowStructMigrateTask",
			Handler:    _Master_ShowStructMigrateTask_Handler,
		},
		{
			MethodName: "UpsertStmtMigrateTask",
			Handler:    _Master_UpsertStmtMigrateTask_Handler,
		},
		{
			MethodName: "DeleteStmtMigrateTask",
			Handler:    _Master_DeleteStmtMigrateTask_Handler,
		},
		{
			MethodName: "ShowStmtMigrateTask",
			Handler:    _Master_ShowStmtMigrateTask_Handler,
		},
		{
			MethodName: "UpsertSqlMigrateTask",
			Handler:    _Master_UpsertSqlMigrateTask_Handler,
		},
		{
			MethodName: "DeleteSqlMigrateTask",
			Handler:    _Master_DeleteSqlMigrateTask_Handler,
		},
		{
			MethodName: "ShowSqlMigrateTask",
			Handler:    _Master_ShowSqlMigrateTask_Handler,
		},
		{
			MethodName: "UpsertCsvMigrateTask",
			Handler:    _Master_UpsertCsvMigrateTask_Handler,
		},
		{
			MethodName: "DeleteCsvMigrateTask",
			Handler:    _Master_DeleteCsvMigrateTask_Handler,
		},
		{
			MethodName: "ShowCsvMigrateTask",
			Handler:    _Master_ShowCsvMigrateTask_Handler,
		},
		{
			MethodName: "UpsertDataCompareTask",
			Handler:    _Master_UpsertDataCompareTask_Handler,
		},
		{
			MethodName: "DeleteDataCompareTask",
			Handler:    _Master_DeleteDataCompareTask_Handler,
		},
		{
			MethodName: "ShowDataCompareTask",
			Handler:    _Master_ShowDataCompareTask_Handler,
		},
		{
			MethodName: "OperateTask",
			Handler:    _Master_OperateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbms_master.proto",
}
