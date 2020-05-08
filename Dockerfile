FROM golang:1.9.4 as builder

# Set the working directory to the app directory
WORKDIR /go/src/happy-birthday

ENV GIT_TERMINAL_PROMPT=1

# Install godeps
RUN go get -u -v github.com/gorilla/mux
RUN go get -u -v go.mongodb.org/mongo-driver/mongo

# Copy the application files
COPY . .

# Build stage
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o happy-birthday .

## App stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/happy-birthday .

# Set ENV
ENV MONGOHOST=
ENV MONGOUSER=
ENV MONGOPASSWORD=

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["./happy-birthday", "run"]
