# Command Line Message App

This is a simple command line chat application written in Go.

## Screenshots

### *Client:*

![Chat App Client](images/chat-app-client.png)

### *Server:*

![Chat App Server](images/chat-app-server.png)

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22.0 or higher

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/m-d-nabeel/go-cli-chat.git
   ```

2. **Change to the project directory:**

   ```sh
   cd go-cli-chat
   ```

3. **Build the server:**

   ```sh
   go build -o server cmd/server/main.go
   ```

   **Run the server:**

   ```sh
   # Run the server:
   ./server/main -port <port_number>

   # Replace <port_number> with the desired port number.
   # If not specified, the default port is 8000.
   ```

4. **Running the Client:**

   Open a new terminal window.

   Navigate to the project directory.

   **Build the client:**

   ```sh
   go build -o client cmd/client/main.go
   ```

   **Run the client:**

   ```sh
   # Run the client:
   ./client/main -port <port_number>

   # Replace <port_number> with the desired port number.
   # If not specified, the default port is 8000.
   ```

   Repeat the steps in "Running the Client" for each additional client you want to connect to the server.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
