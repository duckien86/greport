# declare 
PROJECT_NAME = greport-api
MAIN_PATH = ./cmd/api
BUILD_PATH = ./.build
CONFIG_PATH = ./config

dku:
	docker-compose up -d

dkd:
	docker-compose down

build:
	rm -rf  $(BUILD_PATH)
	mkdir  $(BUILD_PATH)
	cp $(CONFIG_PATH)/* $(BUILD_PATH)
	go build -o $(BUILD_PATH)/$(PROJECT_NAME) $(MAIN_PATH)
	rm -f  $(BUILD_PATH)/config.*

run:
    # EXPORT 
	$(BUILD_PATH)/$(PROJECT_NAME)

run-d:
	nohup $(BUILD_PATH)/$(PROJECT_NAME) > $(PROJECT_NAME).log 2>&1 &

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