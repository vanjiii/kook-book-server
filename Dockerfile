FROM golang:1.20-bookworm AS builder

WORKDIR "/opt/go/src/github.com/vanjiii/kook-book-server"
RUN mkdir -p /opt/go/bin

RUN apt-get update && apt-get upgrade -y

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /opt/go/bin/ ./...

FROM ubuntu:latest AS final

RUN useradd -u 5000 vnj
USER vnj:vnj

WORKDIR /opt/kook-book-server

COPY --from=builder /opt/go/bin/kook-book-server /opt/kook-book-server

EXPOSE 42069

CMD ["./kook-book-server", "serve"]

ENTRYPOINT [""]
