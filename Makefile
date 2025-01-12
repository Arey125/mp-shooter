BIN = mp-shooter

all:
	@go build -o bin/$(BIN) cmd/main.go

run: all
	@templ generate
	@./bin/$(BIN) -f ./files/document.docx

help: all
	@./bin/$(BIN) -h
