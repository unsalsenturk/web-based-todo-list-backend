FROM golang:1.17-alpine3.15 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o /web-based-todolist

FROM alpine

COPY --from=build /web-based-todolist ./app

CMD [ "./app" ]