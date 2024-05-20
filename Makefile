build-proto:
	rm -f ./voucher_proto/voucher.pb.go ./voucher_proto/voucher_grpc.pb.go;
	protoc \
	--go_out=voucher_proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=voucher_proto \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go=/home/lucas94/go/bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=/home/lucas94/go/bin/protoc-gen-go-grpc \
	voucher.proto

run-terraform:
	terraform -chdir=infrastructure/terraform init;
	terraform -chdir=infrastructure/terraform plan;
	terraform -chdir=infrastructure/terraform apply;

run-bdd:
	docker build -f ./infrastructure/docker/Dockerfile.go_app_bdd -t hf-voucher-bdd:latest .;
	docker run --name hf-voucher-bdd hf-voucher-bdd:latest
	@docker rmi -f hf-voucher-bdd >/dev/null 2>&1
	@docker rm $$(docker ps -a -f status=exited -q) -f >/dev/null 2>&1
