FROM golang


ADD hello-world.go /

RUN go build -o /updateapp /hello-world.go

ENTRYPOINT ["/updateapp"]
