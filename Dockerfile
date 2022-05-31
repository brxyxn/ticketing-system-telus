# Back-end Builder
FROM golang:1.18.2-alpine3.15 as golang_builder
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY ./backend ./backend
RUN ls
RUN go build -o /main ./backend/cmd/main.go

#===============================

# Front-end builder
FROM node:latest AS react_builder
WORKDIR /app
COPY ./frontend/package.json ./
COPY ./frontend/package-lock.json ./
RUN npm i
COPY ./frontend ./
RUN npm run build
RUN ls

#===============================

FROM alpine:3.15
WORKDIR /application/
RUN apk --no-cache add ca-certificates
COPY --from=golang_builder /main .
COPY --from=react_builder /build/ ./build/
RUN chmod +x ./main
ENV ENV=Production \
    PORT=5000
EXPOSE ${PORT}
CMD ["./main"]