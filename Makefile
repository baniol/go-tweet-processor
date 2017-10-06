APP_BUILD_VER ?= %appVersion%

DEPEND=github.com/Masterminds/glide 

default: builddocker

depend:
	go get -v $(DEPEND)
	glide install

setup:
	# go get github.com/aws/aws-sdk-go/aws
	# go get gopkg.in/urfave/cli.v1
	go get -v $(DEPEND)
	glide install

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${APP_BUILD_VER} -s" -a -installsuffix cgo -o main ./go/src/go-tweet-processor/cmd/twitter/main.go

buildcmd:
	go build -ldflags "-X main.version=${VERSION} -s" -o daznapi ./cmd_main/main.go # for local build

builddocker:
	docker build -t twitter-processor-2 -f ./Dockerfile.build .
	docker run -t twitter-processor-2 /bin/true
	docker cp `docker ps -q -n=1`:/main .
	chmod 755 ./main
	docker build --rm=true --tag=twitter-processor-2 -f Dockerfile.static .

run: builddocker
	docker run -p 9090:9090 twitter-processor