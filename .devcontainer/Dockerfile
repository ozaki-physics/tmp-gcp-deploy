# Use an official Go runtime as a parent image
FROM golang:1.22 AS golang

# Use an official Node.js runtime as a parent image
FROM node:20 AS nodejs

# Create a new image from the above images
FROM ubuntu:22.04

# Install dependencies and tools
RUN apt-get update && \
    # なんか非推奨らしい?
    # apt-get upgrade -y && \
    apt-get install -y \
    # 最小限でいいときは以下にする
    # apt-get install -y --no-install-recommends
    build-essential \
    curl \
    git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Copy Go and Node.js from the respective images
COPY --from=golang /usr/local/go /usr/local/go
COPY --from=nodejs /usr/local/bin/node /usr/local/bin/
COPY --from=nodejs /usr/local/lib/node_modules /usr/local/lib/node_modules
# COPY --from=nodejs /usr/local/bin/npm /usr/local/bin/
# COPY --from=nodejs /usr/local/bin/npx /usr/local/bin/

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"

# Set Node.js environment variables
ENV PATH="/usr/local/bin:${PATH}"

RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm
RUN ln -s /usr/local/lib/node_modules/npm/bin/npx-cli.js /usr/local/bin/npx

# Google CLI のインストール
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
    gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg && \
    apt-get update -y && \
    apt-get install google-cloud-sdk -y

# # Create and set the working directory
# WORKDIR /test_app

# # Install Go and Node.js dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# COPY package.json package-lock.json ./
# RUN npm install

# # Copy the project files
# COPY . .

# # Expose the application ports
# EXPOSE 8080

# # Start the application (update the command as needed for your application)
# CMD ["go", "run", "main.go"]
