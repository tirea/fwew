SOURCES = fwew.go word.go affixes.go affixes_test.go config.go numbers.go lib.go txt.go version.go

fwew: format all

format:
	gofmt -w $(SOURCES)

all:
	go build -o bin/fwew

install: fwew
	sudo cp bin/fwew /usr/local/bin/
	cp -r .fwew ~/

test: fwew
	go test -v

uninstall:
	sudo rm /usr/local/bin/fwew
	rm -rf ~/.fwew

clean:
	rm -f bin/fwew
