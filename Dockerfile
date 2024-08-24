FROM golang:1.23.0-alpine AS build
ARG PORT=3000

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gems

FROM alpine:3
COPY --from=build /gems /gems

EXPOSE ${PORT}
CMD ["/gems"]