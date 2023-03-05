FROM golang:1.18.2-alpine3.15

# AGS
ARG PERSONAL_ACCESS_TOKEN

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update
RUN apk --update add git make

# package for delve
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

# copy source
# WORKDIR /go/src/app
WORKDIR /go/src/app

COPY . .

# module install
# RUN go mod tidy -go=1.18
RUN go get -u github.com/cosmtrek/air
RUN go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u github.com/kataras/iris/v12
RUN go get -u github.com/syndbg/goenv
RUN go get -u github.com/go-gorp/gorp
RUN go get -u github.com/go-sql-driver/mysql

CMD ["air", "-c", ".air.toml"]
