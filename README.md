# simple-oauth2-server

# Setup (Docker)
Clone the repository and get the .env file from here: https://send.bitwarden.com/#s-oHiNHGmEi1PLKuAP8KTg/M4Qj9u-O5jUUtR7DkmoKhQ (Note: Link expires by 04.04.2025)

In your terminal create the private key:

    openssl genrsa -out private.pem 2048

Then run:

    docker-compose up --build

# (Alternative) Run without Docker
In the .env file adjust the PRIVATE_KEY_PATH, e.g. to "/private.pem". Then run:

    . run.sh

