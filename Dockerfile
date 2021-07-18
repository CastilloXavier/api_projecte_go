FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/api_project
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/api_project-mooc-api/API/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/api_project-mooc-api /go/bin/api_project-mooc-api
ENTRYPOINT ["/go/bin/api_project-mooc-api"]
