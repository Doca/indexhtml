FROM golang:1.11 AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-filechecker


FROM alpine:3.9
MAINTAINER 	Dorian Cantzen <cantzen@extrument.com>
ADD . /mountready
COPY --from=build /src/go-filechecker /usr/local/bin/go-filechecker
RUN chmod +x /usr/local/bin/go-filechecker
CMD ["go-filechecker"]
RUN apk add curl
EXPOSE 3001

