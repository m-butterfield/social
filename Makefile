cloudrunbasecommand := gcloud run deploy --region=us-central1

terraformbasecommand := cd infra && terraform
terraformvarsarg := -var-file=secrets.tfvars

export DB_SOCKET=host=localhost dbname=social

build:
	go build -o bin/server cmd/server/main.go

deploy: docker-build docker-push
	$(cloudrunbasecommand) --image=gcr.io/mattbutterfield/social social

docker-build:
	docker-compose build

docker-push:
	docker-compose push

reset-db:
	dropdb --if-exists social
	createdb social
	go run cmd/migrate/main.go

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix
	cd infra/ && terraform fmt

run-server: export USE_LOCAL_FS=true
run-server:
	go run cmd/server/main.go

test: export DB_SOCKET=host=localhost dbname=social_test
test:
	dropdb --if-exists social_test && createdb social_test && go run cmd/migrate/main.go
	go test -v ./app/...

tf-plan:
	$(terraformbasecommand) plan $(terraformvarsarg)

tf-apply:
	$(terraformbasecommand) apply $(terraformvarsarg)

update-deps:
	go get -u ./...
	go mod tidy
	npm upgrade
	cd infra && terraform init -upgrade && cd -
