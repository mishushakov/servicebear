GOOS=darwin go build -o ./../../../bin/servicebear-darwin -ldflags="-s -w" main.go 
GOARCH=amd64 GOOS=linux go build -o ./../../../bin/servicebear-linux -ldflags="-s -w" main.go 
GOOS=windows go build -o ./../../../bin/servicebear-win -ldflags="-s -w" main.go 