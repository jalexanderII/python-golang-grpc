PHONY: gen_todo

gen_todo:
	protoc -I=./ --go_opt=paths=source_relative --go_out=plugins=grpc:./todo ./todo.proto