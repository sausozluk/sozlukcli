EXECUTABLE := sozluk
TARGET := cmd/main/main.go
CONFIG := cmd/main/sozluk.ini

make:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE) $(TARGET)
	cp $(CONFIG) bin/sozluk.ini

cross:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-linux $(TARGET)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-windows.exe $(TARGET)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-darwin $(TARGET)
	cp $(CONFIG) bin/sozluk.ini
	cp bin/$(EXECUTABLE)-linux bin/$(EXECUTABLE)