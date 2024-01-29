start: gen
	@go run cmd/app/main.go

build: gen
	@go build -o ./.build/app -v ./cmd/app

watch:
	@air -c .air.toml

fmt:
	gofumpt -l -w .

gen: gen-templ gen-frontend

gen-templ:
	@templ generate

gen-frontend:
	cd frontend && vite build

install-tools:
	@go install github.com/a-h/templ/cmd/templ@latest
