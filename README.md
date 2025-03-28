# simple-oauth2-server

# Setup (Docker)
In your terminal create the private key:

    openssl genrsa -out private.pem 2048

Then run:

    docker-compose up --build

# (Alternative) Run without Docker
In the .env file adjust the PRIVATE_KEY_PATH, e.g. to "/private.pem". Then run:

    . run.sh

