FILES=server.go serverEvents.go facebookCalls.go videoContent.go photoContent.go statusContent.go postContent.go util.go

all: server

server: *.go
	go build $(FILES)

clean: *.go
	go clean $(FILES)
