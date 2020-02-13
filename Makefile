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

all: test docker cross-compile

format:
	gofmt -w $(SOURCES)

compile:
	go build -o bin/fwew ./...

cross-compile:
	GOOS=darwin go build -o bin/mac/fwew ./...
	GOOS=linux go build -o bin/linux/fwew ./...
	GOOS=windows go build -o bin/windows/fwew.exe ./...

test: install
	go test -v -cover

docker:
	docker build -t tirea/fwew:$(TAG) .
	docker run -it --rm tirea/fwew:$(TAG) -v -r test

copy:
	@test -n "$(BIN)" || (echo "Error: BIN variable not set. BIN must be set to one of the following:" ; ls bin | grep -v fwew ; exit 1)
	$(CP) bin/$(BIN)/fwew $(BINDEST)/
	cp -r .fwew ~/

install: fwew
	$(CP) bin/fwew $(BINDEST)/
	cp -r .fwew ~/

uninstall:
	$(RM) $(BINDEST)/fwew
	rm -rf ~/.fwew

release:
	cd bin
	tar -czvf fwew-linux-4.0.0-dev.tar.gz linux/fwew
	zip fwew-macos-4.0.0-dev.zip mac/fwew
	zip fwew-windows-4.0.0-dev.zip windows/fwew.exe
	cd -

clean:
	rm -f bin/fwew*
