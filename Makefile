APP_BUILD_VER ?= %appVersion%
APP_NAME=go-tweet-processor
DEPEND=github.com/Masterminds/glide
IMAGE_WEB=go-tweet-processor-web

# setup:
# 	go get -v $(DEPEND)
# 	glide install -force

preparedockerfile:
	sed "s@{{buildgo}}@buildgoweb@g" "Dockerfile.build.template" > "Dockerfile.build"

buildgostream:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${APP_BUILD_VER} -s" -a -installsuffix cgo -o main ./go/src/github.com/baniol/${APP_NAME}/cmd/twitter/main.go

buildgoweb:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${APP_BUILD_VER} -s" -a -installsuffix cgo -o main ./go/src/github.com/baniol/${APP_NAME}/cmd/web/main.go

builddockerstream:
	docker build -t ${IMAGE_WEB} -f ./Dockerfile.build .
	docker run -t ${IMAGE_WEB} /bin/true
	docker cp `docker ps -q -n=1`:/main .
	chmod 755 ./main
	docker build --rm=true --tag=${IMAGE_WEB} -f Dockerfile.static .

builddockerweb:
	docker build -t ${IMAGE_WEB} -f ./Dockerfile.build .
	docker run -t ${IMAGE_WEB} /bin/true
	docker cp `docker ps -q -n=1`:/main .
	chmod 755 ./main
	docker build --rm=true --tag=${IMAGE_WEB} -f Dockerfile.static .