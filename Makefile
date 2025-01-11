generate:
	rm -rf gen/
	git clone https://github.com/babaunba/proto.git
	
	protoc \
		-I=./proto \
		-I=vendor.proto \
		--go_out . \
		--go-grpc_out . \
		--grpc-gateway_out . \
		--openapi_out . \
		$$(find ./proto/ -name '*.proto')
	mv openapi.yaml gen/proto/labels/v1
	
	rm -rf proto/

vendor:
	git clone --filter=blob:none --sparse https://github.com/googleapis/googleapis.git
	cd googleapis && git sparse-checkout set google/api
	
	mv googleapis/google .
	mkdir vendor.proto || true
	mv google vendor.proto
	rm -rf googleapis
