PHONY: gen_auth

gen_auth:
	protoc -I=./ --go_opt=paths=source_relative --go_out=plugins=grpc:./auth ./auth.proto