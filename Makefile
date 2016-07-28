EXECUTABLE=godiffexporter
GOFMT=gofmt -w
GODEPS=go get

GOFILES=\
	main.go\

build:
	go build -o ${EXECUTABLE}

install:
	go install

format:
	${GOFMT} main.go

test:

deps:
	${GODEPS} github.com/waigani/diffparser
	${GODEPS} github.com/prsolucoes/gofpdf

stop:
	pkill -f ${EXECUTABLE}

update:
	git pull origin master
	make install

assets:
	go-bindata -o fonts/bindata.go -pkg fonts -ignore=.gitignore -ignore fonts/bindata.go -ignore .DS_Store fonts/

sample:
	go run main.go -f=pdf -d=/tmp/diff.txt -o=/tmp/out.pdf && open /tmp/out.pdf