#FROM golang:1.14-alpine AS build
#
#WORKDIR /src/
#COPY main.go go.* /src/
#RUN go build -o /bin/demo
#
#FROM scratch
#COPY --from=build /bin/demo /bin/demo
#ENTRYPOINT ["/bin/demo"]


FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go build -o /bin/demo

CMD ["/bin/demo"]



# Take  from https://tutorialedge.net/golang/go-docker-tutorial/
#FROM golang:1.12.0-alpine3.9
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
### Add this go mod download command to pull in any dependencies
#RUN go mod download
### Our project will now successfully build with the necessary go libraries included.
#RUN go build -o main .
### Our start command which kicks off
### our newly created binary executable
#CMD ["/app/main"]