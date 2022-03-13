build:
	echo "Windows, Linux ve MacOS için Derleme İşlemi"
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows64.exe main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/main-macos64 main.go
run:
	go run main.go
hepsi: build run