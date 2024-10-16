FROM golang:1.23-alpine AS builder

WORKDIR /builder

COPY . /builder

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o transform-response main.go


FROM kong
USER root

COPY --from=builder  /builder/transform-response  ./kong/

# reset back the defaults
USER kong
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 8000 8443 8001 8444
STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health
CMD ["kong", "docker-start"]