# GoLang course

Following the course at https://www.udemy.com/cursodego

## Requirements

1. Git
2. Go 1.21+
3. Docker
4. MySQL client
5. MongoDB client
6. Makefile (Optional)

## Clone

```shell
mkdir -p ~/go/src/github.com/humbertodias && cd $_
git clone https://github.com/humbertodias/go-course
cd go-course
```

Resolve dependencies

```shell
go get -d ./...
```

## Run

1. Variables

```shell
go run fundamental/variables/main.go
```

2. GoBuild

```shell
GOOS=linux go build -o myapp intermediate/gobuild/main.go
./myapp
```
eg.

| GOOS | Platform |
|------|----------|
| Mac  | darwin   |
| Win  | windows  |


3. Array

```shell
go run intermediate/array/main.go
```

4. Slice

```shell
go run intermediate/slice/main.go
```

5. File - Reading

```shell
go run intermediate/files/reading/main.go
```

5.1 File - SwfHeader

```shell
go run intermediate/files/swf/main.go intermediate/files/swf/race.swf
```

6. File - Writing

```shell
go run intermediate/files/writing/main.go
```

7. GoRoutine

```shell
go run intermediate/routines/main.go
```

8. Interfaces

```shell
go run intermediate/interfaces/main.go
```

9. Channel

```shell
go run intermediate/channel/main.go
```

10. Channel Select

```shell
go run intermediate/select/main.go
```

11. Web GET

```shell
go run advanced/web-get/main.go
```

12. Web POST

```shell
go run advanced/web-post/main.go
```

13. UnMarshall

```shell
go run advanced/unmarshall/main.go
```

14. Web Server

```shell
go run advanced/web-server/main.go
```
Then Open

http://localhost:8081/

http://localhost:8081/hello

http://localhost:8081/funcao

15. DataBase (SQL/NoSQL)

Start Server

```shell
docker run \
-e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
-e MYSQL_ROOT_PASSWORD= \
-e MYSQL_DATABASE=go-course \
-p 3306:3306 \
-d mysql
```
and

```shell
docker run \
-e MONGO_INITDB_ROOT_USERNAME=go \
-e MONGO_INITDB_ROOT_PASSWORD=go \
-e MONGODB_DATABASE=go-course \
-p 27017:27017 \
-d mongo
```

Load

```shell
mysql --protocol=TCP -h localhost -P 3306 -D go-course -u root < advanced/database/resources/ddl.sql
mysql --protocol=TCP -h localhost -P 3306 -D go-course -u root < advanced/database/resources/dml.sql
```

```shell
mongoimport --authenticationDatabase=admin --uri 'mongodb://go:go@localhost:27017/go-course' \
--collection local \
--jsonArray --file advanced/database/resources/places.json
```

Run

```shell
go run advanced/database/main.go
```

Then Open

http://localhost:8181/sql/45

http://localhost:8181/nosql/99


16: Web ASM

```shell
cd advanced/web-asm
GOOS=js GOARCH=wasm go build -o main.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

go get -u github.com/shurcooL/goexec
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```
Then open the browser

http://localhost:8080

And check out the console.log

![](doc/web-asm-console.png)


# References

[GoLang](https://golang.org)

[Udemy GoLang Course](https://www.udemy.com/cursodego)

[Go Build](https://golang.org/pkg/go/build)

[Go WebASM](https://github.com/golang/go/wiki/WebAssembly)
