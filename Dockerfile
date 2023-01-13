# Using the official Golang image as the base image
FROM golang:latest

# Copying all the project items to the the /app in the container
ADD . /app

# Setting the working directory in the container
WORKDIR /app

# Copying dependencies files to the app working directory
COPY go.mod ./
COPY go.sum ./

# Installing the dependencies from the project
RUN go mod download

# Copy everything now
COPY *.go ./

# Compiling our application now
RUN go build -o /go-create-video

EXPOSE 8080

# Command which is being used to run after the container has been started
CMD ["/go-create-video"]

