start: gen
	@go run cmd/app/main.go

build: gen-templ gen-frontend
	@go build -o ./.build/app -v ./cmd/app

watch:
	@air -c .air.toml

gen-templ:
	@templ generate

gen-frontend:
	cd frontend && vite build

install-tools:
	@go install github.com/a-h/templ/cmd/templ@latest
	