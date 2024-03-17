.PHONY: install_deps up_db

install_deps:
	go mod tidy

up_db:
	docker compose up -d
