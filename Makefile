EXEC = bin/goteach


all: $(EXEC)

$(EXEC):
	go build -o $(EXEC) cmd/goteach/goteach.go
