SOURCES = fwew.go util/affix.go util/txt.go util/version.go

fwew: format all

format:
	gofmt -w $(SOURCES)

all:
	go build -ldflags "-X github.com/tirea/fwew/util.Build=`git rev-parse HEAD`" -o fwew fwew.go
	mv fwew bin/

install: fwew
	sudo cp bin/fwew /usr/local/bin/
	cp -r .fwew ~/

uninstall:
	sudo rm /usr/local/bin/fwew
	rm -rf ~/.fwew

clean:
	rm bin/fwew
