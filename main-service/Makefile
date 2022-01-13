PHONY: gen_auth gen_todo

gen_auth:
	python -m grpc_tools.protoc -I. --python_out=./auth --grpc_python_out=./auth ./auth.proto

gen_todo:
	python -m grpc_tools.protoc -I. --python_out=./todo --grpc_python_out=./todo ./todo.proto