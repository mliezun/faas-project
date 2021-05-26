faas: bindirs
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/faas main.go

bindirs:
	mkdir -p bin/

clean:
	rm -rf bin/