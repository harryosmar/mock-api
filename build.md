```shell script
env GOOS=darwin GOARCH=amd64 go build -o build/mock-api-darwin-64 .\
&& env GOOS=windows GOARCH=amd64 go build -o build/mock-api-windows-64 .\
&& env GOOS=windows GOARCH=386 go build -o build/mock-api-windows-386 . \
&& chmod +X build/
```