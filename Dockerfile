FROM golang:latest AS builder
LABEL stage=builder

ENV GO111MODULE "on"

RUN mkdir /task-manager

WORKDIR /task-manager
COPY . /task-manager

RUN cd /task-manager && go mod download; \
  CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-s -w" -o /go/bin/task-manager ./main.go; \
  CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-s -w" -o /go/bin/task-manager-db ./cmd/database/main.go

FROM golang:latest
LABEL name=task-manager

WORKDIR /task-manager

COPY --from=builder /go/bin/task-manager /task-manager/service
COPY --from=builder /go/bin/task-manager-db /task-manager/database
COPY --from=builder /task-manager/config /task-manager/config

COPY --from=builder /task-manager/init.sh /task-manager/init.sh

RUN apt-get update && apt-get install -y netcat

RUN chmod +x /task-manager/init.sh

EXPOSE 3000

ENTRYPOINT ["/task-manager/init.sh"]

CMD ["/task-manager/init.sh"]
