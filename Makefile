env:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest
	go get -u github.com/swaggo/swag
	mkdir -p build
	docker buildx create --name ct_builder --use --bootstrap --driver-opt image=portus.in.chaitin.net/library/buildkit:buildx-stable-1

init-build-dir:
	mkdir -p build/config/api \
	&& cp cmd/api/config.yaml build/config/api/config.yaml \
	&& mkdir -p build/config/consumer \
	&& cp cmd/consumer/config.yaml build/config/consumer/config.yaml \
	&& mkdir -p build/config/cron \
	&& cp cmd/cron/config.yaml build/config/cron/config.yaml \
	&& mkdir -p build/config/admin \
	&& cp cmd/admin/config.yaml build/config/admin/config.yaml \
	&& cp docker-compose.yml build/docker-compose.yml

swag: init-build-dir
	swag fmt && swag init -g ./cmd/api/main.go -pd
	wire ./cmd/api/wire.go && wire ./cmd/consumer/wire.go && wire ./cmd/cron/wire.go && wire ./cmd/admin/wire.go

checkenv:
ifndef SECRET_TOKEN
	$(error SECRET_TOKEN is undefined)
endif

PLATFORM=local
OUTPUT=type=docker,dest=build/intelligence-backend_image.tar
IMAGE_NAME=portus.in.chaitin.net/safepoint/intelligence-backend:local
DOCKERFILE=Dockerfile.api

image: checkenv
	docker buildx build \
	  --platform ${PLATFORM} \
	  --tag ${IMAGE_NAME} \
	  --secret id=SECRET_TOKEN \
	  --output ${OUTPUT} \
	  --progress plain \
	  --file ${DOCKERFILE} \
	  .

dev: swag
	make image PLATFORM=linux/amd64 DOCKERFILE=Dockerfile.api IMAGE_NAME=portus.in.chaitin.net/safepoint/intelligence-backend:local OUTPUT=type=docker,dest=build/intelligence-backend_image.tar \
	&& make image PLATFORM=linux/amd64 DOCKERFILE=Dockerfile.consumer IMAGE_NAME=portus.in.chaitin.net/safepoint/intelligence-consumer:local OUTPUT=type=docker,dest=build/intelligence-consumer_image.tar \
	&& make image PLATFORM=linux/amd64 DOCKERFILE=Dockerfile.cron IMAGE_NAME=portus.in.chaitin.net/safepoint/intelligence-cron:local OUTPUT=type=docker,dest=build/intelligence-cron_image.tar \
	&& make image PLATFORM=linux/amd64 DOCKERFILE=Dockerfile.admin IMAGE_NAME=portus.in.chaitin.net/safepoint/intelligence-admin:local OUTPUT=type=docker,dest=build/intelligence-admin_image.tar \
	&& docker load < build/intelligence-backend_image.tar \
	&& docker load < build/intelligence-consumer_image.tar \
	&& docker load < build/intelligence-cron_image.tar \
	&& docker load < build/intelligence-admin_image.tar \
	&& cd build && docker compose up -d

prod:
	make image PLATFORM=linux/amd64
