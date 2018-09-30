#Golang REST API boilerplate with TDD

##TODO:
 - Logger (logrus)
 - ~~Graceful Shutdown~~
 - Middleware (eg. Authentication)
 - ~~gRPC transport~~
    should use same model for performance
 - Resilient Pattern
 - Better Error
 - Timeout context
 - Dockerize
 - More test case (TDD)

##Protoc
```
protoc  todo/transport/gRPC/todo.proto --go_out=todo/transport/gRPC --proto_path=/usr/local/include --proto_path=todo/transport/gRPC
```