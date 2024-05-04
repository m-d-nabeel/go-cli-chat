# Command Line Message App

This is a simple command line chat application written in Go.

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
   cd cmd-line-message-app
   ```

3. **Build the server:**

   ```sh
   go build -o server cmd/server/main.go
   ```

   **Run the server:**

   ```sh
   ./server
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
   ./client
   ```

   Repeat the steps in "Running the Client" for each additional client you want to connect to the server.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.