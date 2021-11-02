FROM golang:1.16 as build-env
WORKDIR /go/src/app
ADD . /go/src/app
RUN go mod tidy
RUN go build -o /go/bin/library

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/library /
CMD ["/library"]
