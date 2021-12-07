FROM golang:latest AS build
WORKDIR /go/src/github.com/loqutus/O-1
COPY . ./
RUN make

FROM alpine:latest
WORKDIR /
COPY --from=build /go/src/github.com/loqutus/O-1/bin/o1-linux /o1
CMD ["o1"]
EXPOSE 6969