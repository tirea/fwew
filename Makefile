SOURCES = affixes.go affixes_test.go completer.go config.go fwew.go fwew_test.go lib.go lib_test.go numbers.go numbers_test.go txt.go version.go word.go
TAG = build
OS = nix
ifeq ($(OS),nix)
CP = sudo cp
RM = sudo rm
BINDEST = /usr/local/bin
else ifeq ($(OS),termux)
CP = cp
RM = rm
BINDEST = /data/data/com.termux/files/usr/bin
endif

fwew: format compile

all: test docker

format:
	gofmt -w $(SOURCES)

compile:
	go build -o bin/fwew

test: install
	go test -v

docker:
	docker build -t tirea/fwew:$(TAG) .
	docker run -it --rm tirea/fwew:$(TAG) -v -r test

install: fwew
	$(CP) bin/fwew $(BINDEST)/
	cp -r .fwew ~/

uninstall:
	$(RM) $(BINDEST)/fwew
	rm -rf ~/.fwew

clean:
	rm -f bin/fwew
