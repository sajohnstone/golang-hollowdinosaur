# golang-hollowdinosaur

A simple example using Go and Svelte that generates an Avatar.

## Run the code:

```
./build.sh
```

To debug the frontend you can 
```
cd ./src/frontend
npm run start
cd ../../
./build.sh
```

## add dependacies

```
go get "github.com/gin-gonic/contrib/static"
go get "github.com/gin-gonic/gin"
```

Try inserting a value with
```
curl -XPOST localhost:8080/endpoint -v
```
Then fetching a value with
```
curl localhost:8080/endpoint/1 -v
```