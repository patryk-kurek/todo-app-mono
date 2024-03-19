dev: 
	../bin/air

dependencies: go.mod
	go get github.com/go-sql-driver/mysql
	go get github.com/gorilla/mux
	go get github.com/joho/godotenv

start_db:
	docker compose -f ./docker/docker-compose.yml up -d

stop_db: 
	docker compose -f ./docker/docker-compose.yml down

connect_term_to_db:
	mysql -h localhost --protocol=TCP -u root -p