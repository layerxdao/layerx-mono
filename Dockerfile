FROM golang:1.19.3 as builder

ARG PACKAGE=eventindexer

RUN apt install git curl

RUN mkdir /layerx-mono

WORKDIR /layerx-mono

COPY . .

RUN go mod download

WORKDIR /layerx-mono/packages/$PACKAGE

RUN CGO_ENABLED=0 GOOS=linux go build -o /layerx-mono/packages/$PACKAGE/bin/${PACKAGE} /layerx-mono/packages/$PACKAGE/cmd/main.go

FROM alpine:latest

ARG PACKAGE

RUN apk add --no-cache ca-certificates

COPY --from=builder /layerx-mono/packages/$PACKAGE/bin/$PACKAGE /usr/local/bin/

ENTRYPOINT ["$PACKAGE"]