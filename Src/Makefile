FILES = .\todo\main.go .\todo.go .\colors.go 
BINS = .\bin\todo.exe

all: $(FILES)
	mkdir bin
	go build -o $(BINS) .\todo\main.go

clean:
	rm -r bin