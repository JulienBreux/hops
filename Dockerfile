# Stage 1: Build the Svelte SPA
FROM node:20-alpine AS frontend-builder
WORKDIR /app/web
COPY web/package.json web/package-lock.json* ./
RUN npm ci || npm install
COPY web/ .
RUN npm run build

# Stage 2: Build the Go binary
FROM golang:1.22-alpine AS backend-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o hops main.go

# Stage 3: Assemble the final minimal image
FROM alpine:3.19
WORKDIR /app
# Install CA certificates for external API calls if needed
RUN apk --no-cache add ca-certificates tzdata

# Copy the compiled Go binary
COPY --from=backend-builder /app/hops .
# Copy the compiled Svelte SPA assets (if they need to be served by Go, we place them in a known directory)
# Note: For SvelteKit with static adapter, it outputs to 'build'. If using auto adapter, it might output differently.
# Assuming standard static files serving.
COPY --from=frontend-builder /app/web/build ./web/build

# Set permissions for the binary
RUN chmod +x ./hops

# Expose the port the Go application runs on
EXPOSE 8080

# Command to run the application
CMD ["./hops"]
