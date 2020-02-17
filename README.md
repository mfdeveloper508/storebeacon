## FUPP CMS & API


#### Setup
Install fresh
```
go get github.com/pilu/fresh
```

Install dependencies
```
dep ensure
```

Database
```
store:store@localhost/store
```


#### Run
Run project:
```
fresh
```


### Development
On new imports or just for fun - https://golang.github.io/dep/docs/daily-dep.html
```
dep ensure
```


### Running tests
Run tests on any api update or just for fun:
```
make build
make test
```


### Kill Dockers
```
docker rm -f $(docker ps -a -q)
```


### Deployment
Make a new merge request into master and merge.


#### API Doc
Install Swagger
```
go get github.com/yvasiyarov/swagger
```

Generate doc
```
swagger -apiPackage="github.com/store" -mainApiFile=main.go -output=./API.md -format=markdown
```