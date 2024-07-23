###############################################
# FRONTEND
FROM oven/bun:1.1.20-slim AS build_frontend

WORKDIR /app

# Install dependencies
COPY ./bun.lockb ./package.json ./
RUN bun install

# Copy files
COPY ./tsconfig.json ./vite.config.ts ./
COPY public ./public
COPY frontend ./frontend

# Execute build (output in dist)
RUN bun run build

###############################################
# BACKEND
FROM golang:1.22.5 AS build_go

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY *.go ./
COPY server ./server
COPY --from=build_frontend /app/dist ./dist 

# Compile
ENV CGO_ENABLED=0 \
  GOOS=linux
RUN mkdir bin && go  build -o bin/server /app/server/cmd/http/main.go


###############################################
# PROD
FROM scratch AS prod 

WORKDIR /app

ENV IS_DEV=false

EXPOSE 8000

COPY --from=build_go /app/bin/server /app/server

CMD ["/app/server", "0.0.0.0:8000"]


