# generate protobuf и grpc.pb for task-service
.PHONY: gentaskpb
gentaskpb:
	protoc -I api/proto/v1/task task.proto --go_out=./api/pb/v1/task --go_opt=paths=source_relative --go-grpc_out=./api/pb/v1/task/ --go-grpc_opt=paths=source_relative

# generate protobuf и grpc.pb for auth-service
.PHONY: genauthpb
genauthpb:
		protoc -I api/proto/v1/auth auth.proto --go_out=./api/pb/v1/auth --go_opt=paths=source_relative --go-grpc_out=./api/pb/v1/auth/ --go-grpc_opt=paths=source_relative

# generate protobuf for user created event
.PHONY: genusercreatedpb
genusercreatedpb:
	protoc -I api/proto/v1/events/user created.proto --go_out=./api/pb/v1/events/user --go_opt=paths=source_relative

# generate protobuf for user deleted event
.PHONY: genuserdeletedpb
genuserdeletedpb:
	protoc -I api/proto/v1/events/user deleted.proto --go_out=./api/pb/v1/events/user --go_opt=paths=source_relative