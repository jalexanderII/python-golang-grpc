PHONY: gen_auth gen_todo

gen_auth:
	python -m grpc_tools.protoc -I./protos --python_out=./auth --grpc_python_out=./auth ./protos/auth.proto

gen_todo:
	python -m grpc_tools.protoc -I./protos --python_out=./todo --grpc_python_out=./todo ./protos/todo.proto