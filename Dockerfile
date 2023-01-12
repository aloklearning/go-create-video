# Using the official Golang image as the base image
FROM golang:latest

# Copying all the project assets to the the /app in the container
ADD . /app

# Setting the working directory in the container
WORKDIR /app

# Copying everything from the project to the app working directory
COPY go.mod ./
COPY go.sum ./

# Installing the dependencies from the project
RUN go mod download

# Copy everything now
COPY *.go ./

# Compiling our application now
RUN go build -o /go-create-video

EXPOSE 8080

CMD ["/go-create-video"]

