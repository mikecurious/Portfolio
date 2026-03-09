# Stage 1: Build React frontend
FROM node:20-alpine AS frontend
RUN apk add --no-cache git
RUN git clone https://github.com/mikecurious/mikkoh-bytes /frontend
WORKDIR /frontend

# Replace Lovable favicon with custom MB favicon
COPY favicon.svg /frontend/public/favicon.svg
RUN rm -f /frontend/public/favicon.ico && \
    sed -i 's|<meta name="viewport"|<link rel="icon" href="/favicon.svg" type="image/svg+xml" />\n    <meta name="viewport"|' /frontend/index.html

RUN npm install
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.23-alpine AS backend
WORKDIR /build
COPY go.mod ./
RUN go mod tidy
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o portfolio-server .

# Stage 3: Final image
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=backend /build/portfolio-server .
COPY --from=frontend /frontend/dist ./dist
RUN addgroup -S portfolio && adduser -S portfolio -G portfolio
RUN chown -R portfolio:portfolio /app
USER portfolio
EXPOSE 8083
HEALTHCHECK --interval=30s --timeout=10s --start-period=10s --retries=3 \
  CMD wget -qO- http://localhost:8083/ || exit 1
CMD ["./portfolio-server"]
