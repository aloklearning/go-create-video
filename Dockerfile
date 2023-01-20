# Using the official Golang image as the base image
FROM golang:latest

# Copying all the project items to the the /app in the container
# Makes a new layer everytime something changes in the prooject
# Not an efficient solution but a hack to make the docker build 
# ADD . /app

# Setting the working directory in the container
WORKDIR /app

# Copying dependencies files to the app working directory
COPY go.mod go.sum ./

# Installing the dependencies from the project
RUN go mod download && go mod verify

# Copy everything now
COPY . .

# As an alternative, you can only copy the selected items from the project
# COPY /cmd/ ./cmd/
# COPY /pkg/db/ ./pkg/db/
# COPY /pkg/handlers/*.go ./pkg/handlers/
# COPY /pkg/src/*.go ./pkg/src/

# Main file is in the cmd folder, hence making the setting the 
# New working directory now for all the executables
WORKDIR /app/cmd

# Compiling our application now
RUN go build -o /go-create-video

EXPOSE 8080

# Command which is being used to run after the container has been started
CMD ["/go-create-video"]

