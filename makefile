dku:
	docker-compose up -d
dkd:
	docker-compose down
build:
	mkdir  ./.build
	cp config.yml ./.build/
	go build -o ./.build/2ndbrand 
	rm -f  ./.build/config.yml
run:
	./.build/2ndbrand
run-d:
	nohup ./.build/2ndbrand > 2ndbrand.log 2>&1 &
.app: build run
.docker: docker-d docker-u