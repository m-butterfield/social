build:
	go build -o bin/server cmd/server/main.go

docker-build:
	docker-compose build

docker-push:
	docker-compose push

db:
	createdb social

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix
	terraform fmt infra/

run-server:
	DB_SOCKET="host=localhost dbname=social" USE_LOCAL_FS=true go run cmd/server/main.go

test:
	dropdb --if-exists social_test && createdb social_test && psql -d social_test -f schema.sql
	DB_SOCKET="host=localhost dbname=social_test" go test -v ./app/...

update-deps:
	go get -u ./...
	go mod tidy
	npm upgrade
	cd infra && terraform init -upgrade && cd -
