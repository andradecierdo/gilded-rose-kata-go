# GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test test/...
```

- Run tests and coverage :

```shell
go test test/... -coverprofile=coverage.out

go tool cover -html=coverage.out
```
