## Borsch Playground Service

This API allows to execute code written in Borsch programming language
without installing the interpreter on the local machine.

Build and set up the application:
```shell
mkdir "bin"
go build -o ./bin/borschplayground main.go
cp ./settings.json ./bin
cd ./bin
```

Migrate the database:
```shell
./borschplayground migrate
```

Run the server:
```shell
./borschplayground --address 127.0.0.1:8080
```

### API
Check out the [documentation](https://app.swaggerhub.com/apis-docs/YURALISOVSKIY98/BorschPlaygroundService/1.0.0).
