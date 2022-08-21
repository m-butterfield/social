cloudrunbasecommand := gcloud run deploy --region=us-central1
deployservercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/social social
deployworkercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/social-worker social-worker

terraformbasecommand := cd infra && terraform
terraformvarsarg := -var-file=secrets.tfvars

export DB_SOCKET=host=localhost dbname=social
export CGO_CFLAGS_ALLOW=-Xpreprocessor

build: build-server build-worker

build-server: run-webpack-prod
	go build -o bin/server cmd/server/main.go

build-worker:
	go build -o bin/worker cmd/worker/main.go

deploy: docker-build docker-push
	$(deployservercommand)
	$(deployworkercommand)

deploy-server: docker-build-server docker-push-server
	$(deployservercommand)

deploy-worker: docker-build-worker docker-push-worker
	$(deployworkercommand)

docker-build:
	docker-compose build

docker-build-server:
	docker-compose build server

docker-build-worker:
	docker-compose build worker

docker-push:
	docker-compose push

docker-push-server:
	docker-compose push server

docker-push-worker:
	docker-compose push worker

reset-db:
	dropdb --if-exists social
	createdb social
	go run cmd/migrate/main.go

migrate:
	go run cmd/migrate/main.go

generate:
	go run github.com/99designs/gqlgen generate

fmt:
	go fmt ./...
	npx eslint app/static/ts/ --fix
	cd infra/ && terraform fmt

run-server: export USE_LOCAL_FS=true
run-server: export SQL_LOGS=true
run-server: export GQL_PLAYGROUND=true
run-server: export WORKER_BASE_URL=http://localhost:8001/
run-server:
	go run cmd/server/main.go

run-worker: export SQL_LOGS=true
run-worker:
	go run cmd/worker/main.go

run-webpack:
	yarn run webpack --mode development --watch

run-webpack-prod:
	rm -rf app/static/js/dist
	yarn run webpack --mode production

test: export DB_SOCKET=host=localhost dbname=social_test
test:
	dropdb --if-exists social_test && createdb social_test && go run cmd/migrate/main.go
	go test -v ./app/...

tf-plan:
	$(terraformbasecommand) plan $(terraformvarsarg)

tf-apply:
	$(terraformbasecommand) apply $(terraformvarsarg)

tf-refresh:
	$(terraformbasecommand) apply $(terraformvarsarg) -refresh-only

update-deps:
	go get -u ./...
	go mod tidy
	yarn upgrade
	cd infra && terraform init -upgrade && cd -
