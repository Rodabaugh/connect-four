# About
connect-four is a web application built to explore HTMX, templ, and websockets. The game board is shared globally, updating in real-time with each move. Grab some friends and enjoy a game together!

This project was created as a personal exercise and is not intended for production use.

# Project Highlights
This project involved:
- Server-side management of the game state.
- Rendering the game state on the client-side using templ.
- Utilizing websockets to broadcast move updates to all connected clients.

# Hosting
This project is hosted in a Docker container on a VPS, with an Nginx reverse proxy providing public access. This setup allows for SSL via Let's Encrypt (though not critical for this application) and enables traffic monitoring through Nginx logs. If you'd like to host it yourself, here's how to get started:

## Get the repo
1. Clone the repository: `git clone https://github.com/Rodabaugh/connect-four/`
2. Navigate to the project directory: `cd connect-four`

## Configuration
Environment variables are used for configuration. You can set the server port using:```PORT=8080```

## Build and Run
Once your `.env` is configured, you can build and run the application.

Build the application: `make build`

Run the backend: `./connect-four`

After the application is running, you can configure your server to run it as a service. Alternatively, you can use the provided Dockerfile to build and run the application in a Docker container.