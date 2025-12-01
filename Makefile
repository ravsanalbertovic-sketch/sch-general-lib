.PHONY: gentaskpb
gentaskpb:
	protoc -I api/proto/v1/task task.proto --go_out=./api/pb/v1/task --go_opt=paths=source_relative --go-grpc_out=./api/pb/v1/task/ --go-grpc_opt=paths=source_relative

.PHONY: genauthpb
genauthpb:
	protoc -I api/proto/v1/auth auth.proto --go_out=./api/pb/v1/auth --go_opt=paths=source_relative --go-grpc_out=./api/pb/v1/auth/ --go-grpc_opt=paths=source_relative