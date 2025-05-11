# Specifies a parent image
FROM golang:1.23.1-bookworm

# Creates an app directory to hold your app’s source code
WORKDIR /usr/src/app

RUN apt-get update

# Installs Node.js and npm
RUN apt-get install -y npm 
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
RUN . ~/.nvm/nvm.sh && nvm install 20.0.0

# Installs nodemon
RUN npm install -g nodemon

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o ./app ./cmd/server/main.go  

# Tells Docker which network port your container listens on
EXPOSE 3001 

# Specifies the executable command that runs when the container starts
CMD ["go", "run", "./main/init/app.go"]
