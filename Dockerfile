# Use the official Golang image as the base image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Set the Go proxy
ENV GOPROXY=https://proxy.golang.org

# Initialize Go modules (if not already initialized)
COPY go.mod go.sum ./
RUN go mod download

# Install make, wget, unzip, default-jre, and docker-compose
RUN apt-get update && \
    apt-get install -y make wget unzip default-jre && \
    apt-get install -y docker-compose

# Install Ginkgo and Gomega
RUN go install github.com/onsi/ginkgo/v2/ginkgo
RUN go get github.com/onsi/gomega/...

# Install Allure command-line tool
RUN wget https://github.com/allure-framework/allure2/releases/download/2.14.0/allure-2.14.0.zip && \
    unzip allure-2.14.0.zip -d /opt/ && \
    ln -s /opt/allure-2.14.0/bin/allure /usr/bin/allure

# Copy the entire source code into the container
COPY . .

# Default command starts a bash shell for interactive use
CMD ["/bin/bash"]
