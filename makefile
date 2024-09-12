PROJECT_NAME = greport-api
dku:
	docker-compose up -d

dkd:
	docker-compose down

build:
	rm -rf  ./.build/
	mkdir  ./.build
	cp config.yml ./.build/
	go build -o ./.build/$(PROJECT_NAME) 
	rm -f  ./.build/config.yml

run:
	./.build/$(PROJECT_NAME)

run-d:
	nohup ./.build/$(PROJECT_NAME) > $(PROJECT_NAME).log 2>&1 &

.app: build run

.docker: docker-d docker-u

watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi