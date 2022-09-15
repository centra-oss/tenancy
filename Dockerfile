FROM golang:1.19-alpine AS build

ENV CGO_ENABLED=0

RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor vendor

COPY Taskfile.yaml Taskfile.yaml

COPY cmd cmd
COPY util util

RUN go build -ldflags="-w -s" -o /tenancy ./cmd/tenancy/tenancy.go

################################################################################

FROM scratch

COPY --from=build /tenancy /tenancy

ENTRYPOINT [ "/tenancy" ]

