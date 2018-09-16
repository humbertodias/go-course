fmt:
	go fmt ./...

dep:
	go get -d ./...

# MySQL
mysql_start:
	docker run \
	-e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
	-e MYSQL_ROOT_PASSWORD= \
	-e MYSQL_DATABASE=go-course \
	-p 3306:3306 \
	-d mysql

mysql_client_install:
	sudo apt install mysql-client

mysql_import:
	mysql -h localhost -P 3306 -D go-course -u root < advanced/database/resources/ddl.sql
	mysql -h localhost -P 3306 -D go-course -u root < advanced/database/resources/dml.sql

# MongoDB
mongodb_start:
	docker run \
	-e MONGODB_USERNAME=go \
	-e MONGODB_PASSWORD=go \
	-e MONGODB_DATABASE=go-course \
	-p 27017:27017 \
	-d bitnami/mongodb

mongodb_client_install:
	sudo apt install mongo-tools

mongodb_import:
	mongoimport --uri mongodb://go:go@localhost:27017/go-course \
	--collection local \
	--jsonArray --file advanced/database/resources/places.json

# PostGres
postgres_start:
	docker run \
	-e POSTGRES_PASSWORD= \
	-e POSTGRES_USER=root \
	-e POSTGRES_DB=go-course \
	-p 5432:5432 \
	-d postgres

postgres_client_install:
	sudo apt install postgresql-client

postgres_import:
	psql -h localhost -p 5432 go-course root < advanced/database/resources/ddl.sql
	psql -h localhost -p 5432 go-course root < advanced/database/resources/dml.sql

wait_containers:
	sleep 10

start_mysql_and_mongodb:	mysql_start	mongodb_start	wait_containers	mongodb_import	mysql_import

docker_rm_all:
	docker ps -qa | xargs -r docker rm -f

clean:
	rm -f myapp