# Server-Sent Events (SSE) Example with Golang

This repository contains a simple Golang application demonstrating how to implement Server-Sent Events (SSE) without using any web frameworks. SSE enables real-time updates from the server to the client without continuous polling.

## Features

- Provides a basic HTML page with a button that triggers SSE events when clicked.
- Sends server-generated events to connected clients when the button is clicked.
- Demonstrates SSE communication using vanilla Go HTTP.

## Usage

1. Clone the repository: `git clone https://github.com/phuongdanh/golang-sse.git`
2. Navigate to the repository directory: `cd golang-sse`
3. Run the application: `go run main.go`
4. Access the application in your browser at `http://localhost:8080`.

## How It Works

1. When you access the application in your browser, you'll see a button labeled "Send Event".
2. Clicking the button triggers a POST request to the server, which sends an SSE event with a timestamp to all connected clients.
3. The client-side JavaScript establishes an SSE connection to the server and displays the received events in real-time.

## Requirements

- Go (Golang) must be installed on your system.

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to use this example to learn about SSE implementation or as a starting point for your own SSE-based projects.

---

If you find this example useful, consider giving it a ⭐️ on GitHub!
