GOOS=windows GOARCH=amd64 go build -o kitwork-win.exe .
GOOS=linux GOARCH=arm64 go build -o kitwork-linux .
GOOS=darwin GOARCH=amd64 go build -o kitwork-mac .
