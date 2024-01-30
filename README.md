# goLangExampleService

Commands to run file:

Change the gopath to the xxx\gitlab.com\saadaminj
In each package use "go build" to compile& create executable file in windows
run ".\executableName" to run the executable file in windows

docker file is in the root of the service catFactService

build docker image using this cmd:
docker build -t mygoservice-image .

to run the docker image use this cmd:
docker run -p 8080:8080 mygoservice-image

to test use the url in postman or thunderbolt
http://localhost:8080/getCatFact
