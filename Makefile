build-proto:
	protoc \
	--go_out=client_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=client_proto \
	--go-grpc_opt=paths=source_relative \
	client.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;
