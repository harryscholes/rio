# Rio

# Run

```sh
go run server/main.go &
go run client/main.go "my friend"
```

Output:

```sh
❯ go run client/main.go "my friend"
{"level":"info","message":"Hello my friend","time":"2021-04-29T15:06:42+01:00","message":"Greeting: Hello my friend"}
```