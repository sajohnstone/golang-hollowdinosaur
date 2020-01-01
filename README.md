# golang-hollowdinosaur

Run the code:

```
go install
go run ./src/main.go
```

Try inserting a value with
```
curl -XPOST localhost:8080/endpoint -v
```
Then fetching a value with
```
curl localhost:8080/endpoint/1 -v
```