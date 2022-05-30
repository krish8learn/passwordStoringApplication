postgres start:
    sudo docker run --name postgres -p 5434:5434 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	  sudo docker exec -it postgres createdb --username=root --owner=root password_store_application

dropdb:
	  sudo docker exec -it postgres dropdb password_store_application

migrateCreate:
		migrate create -ext sql -dir dao/migration -seq init_schema

migrateup:
		migrate -path dao/migration -database "postgresql://root:krish@knight8@localhost:5432/password_store_application?sslmode=disable" -verbose up

migratedown:
		migrate -path dao/migration -database "postgresql://root:krish@knight8@localhost:5432/password_store_application?sslmode=disable" -verbose down

genGOStruct:
		tables-to-go -v -t pg -d password_store_application -h 127.0.0.1 -s public -u root -p krish@knight8 -of ../passwordStoringApplication/dto
		
startDB:
		sudo docker start -a postgres

startApp:
		sudo docker start --attach passwordapp
		
.PHONY: createdb dropdb migrateCreate migrateup migratedown genGOStruct
