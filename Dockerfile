FROM golang:1.13.4-alpine AS build_deps
RUN apk add --no-cache git
WORKDIR /workspace
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

FROM build_deps AS build
COPY . .
RUN CGO_ENABLED=0 go build -o webhook -ldflags '-w -extldflags "-static"' .

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build /workspace/webhook /usr/local/bin/webhook
ENTRYPOINT ["webhook"]