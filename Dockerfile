# Select the Go runtime version
FROM golang:1.18-alpine

# Setup the working directory
WORKDIR /app

# Copy the source code
ADD . /app

# Install the dependencies
RUN go mod download

# Build the application
RUN go build -o /main

# Add the user
RUN adduser -u 1001 -D -s /bin/sh -G ping 1001

# Change the user and group of the executable file
RUN chown 1001:1001 /main

# Change the executable file permission
RUN chmod +x /main

# Login as the user 1001
USER 1001

# Opening port 3000
EXPOSE 3000

# Run the application
CMD ./main