FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

RUN go build -mod=readonly -v -o server


RUN  set 0x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
     --no-insatll-recommends \
     ca-cartificates && \
     rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /appp/server
COPY app/config.env /app

ENTRYPOINT [ "app/server" ]


