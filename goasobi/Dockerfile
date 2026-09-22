# syntax=docker/dockerfile:1

# ---- builder ----
FROM golang:1.27-trixie AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /out/hello ./goasobi

# ---- runner ----
FROM debian:trixie-slim AS runner

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && useradd -r -u 1000 -m hello

COPY --from=builder /out/hello /usr/local/bin/hello

USER hello

EXPOSE 8080

ENTRYPOINT ["hello"]
CMD ["-p", "8080"]
