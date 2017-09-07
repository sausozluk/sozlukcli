EXECUTABLE := sozlukcli
TARGET := cmd/main/main.go
CONFIG := cmd/main/config.json

make:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-linux $(TARGET)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE).exe $(TARGET)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-darwin $(TARGET)
	cp $(CONFIG) bin/config.json