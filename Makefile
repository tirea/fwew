SOURCES = affixes.go affixes_test.go completer.go config.go fwew.go lib.go numbers.go txt.go version.go word.go

fwew: format compile

all: test docker

format:
	gofmt -w $(SOURCES)

compile:
	go build -o bin/fwew

test: install
	go test -v

docker:
	docker build -t tirea/fwew:build .
	docker run -it --rm tirea/fwew:build -v -r test

install: fwew
	sudo cp bin/fwew /usr/local/bin/
	cp -r .fwew ~/

uninstall:
	sudo rm /usr/local/bin/fwew
	rm -rf ~/.fwew

clean:
	rm -f bin/fwew
