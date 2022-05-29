# Back-end Builder
FROM golang:1.18.2-alpine3.15 as golang_builder
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./backend/cmd/main.go

#===============================

# Front-end builder
FROM node:latest AS react_builder
WORKDIR /app
COPY ./frontend/package.json ./
COPY ./frontend/package-lock.json ./
RUN npm i
COPY ./frontend ./
RUN npm run build

#===============================

FROM alpine:3.15
WORKDIR /application/
RUN apk --no-cache add ca-certificates
COPY --from=golang_builder /main .
COPY --from=react_builder /app/build/ ./backend/build/
RUN chmod +x ./main
ENV ENV=Production \
    PORT=5000
EXPOSE ${PORT}
CMD ["./main"]