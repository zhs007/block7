CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./xlsx2map/main.go
mv main.exe xlsx2map.exe
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ./checkmaps/main.go
mv main.exe checkmaps.exe