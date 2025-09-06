##
## Build
##
ARG GOVERSION
FROM golang:${GOVERSION} AS build


ENV GOPROXY=https://goproxy.cn,direct

COPY kk-schedule-server /kk-schedule

WORKDIR /kk-schedule

RUN CGO_ENABLED=0 GOOS=linux go build -o /kk-schedule/output /kk-schedule/internal/main

##
## Deploy
##
FROM alpine:latest



COPY --from=build /kk-schedule/output /kk-schedule/output
COPY kk-schedule-ui/.output/public /kk-schedule/public

WORKDIR /kk-schedule

CMD ["./output"]

