FROM node:24-alpine AS builder

WORKDIR /app

COPY ./frontend/package*.json ./
RUN npm install

COPY ./frontend/ .

RUN npm run build

FROM golang:1.24

COPY ./backend/ /app/
COPY --from=builder /app/dist /app/static/

WORKDIR /app

RUN go build ./

# ENTRYPOINT [ "/app/sub", "-c", "/app/config.yml" ]
ENTRYPOINT [ "tail", "-f", "/dev/null" ]
