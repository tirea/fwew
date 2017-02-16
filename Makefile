SOURCES = fwew.go util/affix.go util/txt.go util/version.go

fwew: format all

format:
	gofmt -w $(SOURCES)

all:
	go build -o bin/fwew fwew.go
	
install: fwew
	sudo cp bin/fwew /usr/local/bin/
	cp -r .fwew ~/

uninstall:
	sudo rm /usr/local/bin/fwew
	rm -rf ~/.fwew

clean:
	rm bin/fwew