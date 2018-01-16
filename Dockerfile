# build stage
FROM golang:1.8-alpine3.6 AS build-env
ADD . /src/github.com/t-tran/lucas
ENV GOPATH /:/src/github.com/t-tran/lucas/vendor
WORKDIR /src/github.com/t-tran/lucas
RUN go build -o app


# test stage
#FROM golang:1.8-alpine3.6
#WORKDIR /src/github.com/t-tran/lucas
#RUN go test


# release stage
FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=build-env /src/github.com/t-tran/lucas/app /app/
COPY --from=build-env /src/github.com/t-tran/lucas/frontend/static/ /app/static/
CMD ["./app","run"]
