mysql:
	docker run --name books-mysql -p 9002:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql

createdb: 
	docker exec -it books-mysql mysql --user='root' --password='secret' --execute='CREATE DATABASE books_db'

dropdb:
	docker exec -it books-mysql mysql --user='root' --password='secret' --execute='DROP DATABASE books_db'

migrateup:
	migrate -path migration/ -database "mysql://root:secret@tcp(localhost:9002)/books_db" -verbose up

migratedown: 
	migrate -path migration/ -database "mysql://root:secret@tcp(localhost:9002)/books_db" -verbose down

server:
	go run cmd/main.go

.PHONY: mysql createdb dropdb	migrateup migratedown server