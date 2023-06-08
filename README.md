# REST API with Clean Architecture

This project is a personal learning project aimed at building a URL Shortener API using Go with a clean architecture approach. The technology stack used includes Echo, PostgreSQL, and Redis.

## Features

- User authentication and authorization
- Create and redirect URL
- JSON Web Token (JWT) based authentication
- PostgreSQL database integration
- Redis caching for improved performance

## Installation

To run this project locally, follow these steps:

1. Clone the repository: `git clone https://github.com/ilhamnyto/url-shortener-go.git`
2. Install the required dependencies: `go mod download`
3. Copy the env files: `cp .env.example .env`.
4. Set up the Server host, port, PostgreSQL database, Redis and configure the connection details in `.env`.
5. Run the database migrations by running the server: `go run cmd/api/main.go`

## API Documentation

For detailed information on the API endpoints and their usage, refer to the [API Documentation](https://documenter.getpostman.com/view/13820554/2s93saaDRa).

## Configuration

The project's configuration is stored in the `.env` file. Update this file to adjust the server port, database connection details, Redis configuration, and other settings as needed.


## License

This project is licensed under the [MIT License](./LICENSE).

## Acknowledgments

This project was made possible by the following open-source libraries:

- [Echo](https://github.com/labstack/echo)
- [PostgreSQL](https://www.postgresql.org)
- [Redis](https://redis.io)

