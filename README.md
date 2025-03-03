# Gin API Project

This is a Gin API project.

## Prerequisites
- Go 1.24+
- Gin Framework

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/gin-api-project.git
    cd gin-api-project
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the server:
    ```sh
    go run main.go
    ```

2. The server will start at `http://localhost:8080`.

## Endpoints

### GET /ping
- Description: Health check endpoint.
- Response: `{"message": "pong"}`

### POST /create
- Description: Create a new resource.
- Request Body: 
    ```json
    {
        "name": "example",
        "value": "example_value"
    }
    ```
- Response: `{"id": "resource_id", "name": "example", "value": "example_value"}`

## Contributing

1. Fork the repository.
2. Create a new branch for your feature: `git checkout -b feature-name`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
