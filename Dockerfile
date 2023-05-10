# 1 шаг - сборки
FROM golang:1.19-alpine AS build_stage
COPY . /go/src/VK_BOT_GO
WORKDIR /go/src/VK_BOT_GO/cmd
RUN go env -w GO111MODULE=auto
RUN go install .

# 2 шаг
FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/bot /app_binary/
RUN chmod +x ./bot
ENTRYPOINT ./bot