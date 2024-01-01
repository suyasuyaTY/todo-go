CMD := docker-compose

up:
	@$(CMD) up -d

run:
	@$(CMD) exec app go run main.go

down:
	@$(CMD) down

clean:
	@$(CMD) down --rmi all --volumes --remove-orphans
migrate-up:
	@$(CMD) exec app migrate -path "./db/migrations" -database "mysql://test:test@tcp(db:3306)/my_todolist" -verbose up
migrate-down:
	@$(CMD) exec app migrate -path "./db/migrations" -database "mysql://test:test@tcp(db:3306)/my_todolist" down