FROM golang:1.18-alpine

# Setup the working directory
WORKDIR /app

# Copy the source code
ADD . /app

# Install the dependencies
RUN go mod download

# Build the application
RUN go build -o /main

# Change the user and group
RUN adduser -u 1001 -D -s /bin/sh -G ping 1001
RUN chown 1001:1001 /main

RUN chmod +x /main

USER 1001

# Opening ports
EXPOSE 3000

# Run the application
CMD ./main