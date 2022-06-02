FROM golang:alpine AS masterclientbuilds
WORKDIR /appbuilds
COPY . .
RUN go mod tidy
RUN go build -o binary

FROM alpine:latest as masterclientrelease
WORKDIR /app
RUN apk add tzdata
COPY --from=masterclientbuilds /appbuilds/binary .
COPY --from=masterclientbuilds /appbuilds/.env /app/.env
ENV PORT=5000
ENV PATH_API="https://api.isbpanel.com/"
ENV TZ=Asia/Jakarta

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT [ "./binary" ]