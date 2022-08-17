
1. Slide Deck
2. Talk About Connect and Connect Protocol
3. Setup Simple Servers
4. Show demo with grpcurl and the first server

```shell
go run ./grpc/simple/server
```

```shell
grpcurl --plaintext localhost:4444 list
```

```shell
grpcurl --plaintext -d '{"idea":{"name":"connors idea", "category":"daily thoughts", "description":"hello go guild"}}' localhost:4444 simple.v1.SimpleService/CreateIdea
```

```shell
grpcurl --plaintext -d '{"name":"connors idea"}' localhost:4444 simple.v1.SimpleService/GetIdea
```

```shell
grpcurl --plaintext -d '{"name":"not found"}' localhost:4444 simple.v1.SimpleService/GetIdea
```


```shell
curl  --header 'Content-Type: application/json' --data '{"name":"connors idea"}' http://localhost:4444/simple.v1.SimpleService/GetIdea
```

```shell
go run ./connect/simple/server
```

```shell
grpcurl --plaintext localhost:5555 list
```

```shell
grpcurl --plaintext -d '{"idea":{"name":"connors idea", "category":"daily thoughts", "description":"hello go guild"}}' localhost:5555 simple.v1.SimpleService/CreateIdea
```

```shell
grpcurl --plaintext -d '{"name":"connors idea"}' localhost:5555 simple.v1.SimpleService/GetIdea
```

```shell
grpcurl --plaintext -d '{"name":"not found"}' localhost:5555 simple.v1.SimpleService/GetIdea
```

```shell
curl  --header 'Content-Type: application/json' --data '{"name":"connors idea"}' http://localhost:5555/simple.v1.SimpleService/GetIdea
```
