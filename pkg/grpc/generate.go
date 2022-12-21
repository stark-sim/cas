package grpc

// 手动更新 proto，不使用 entproto
//go:generate protoc --proto_path=./pb --go_out=plugins=grpc:./pb ./pb/cas.proto
