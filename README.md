# Go RSS Aggregator

A fast and lightweight RSS aggregator built with Go. This project allows you to fetch, parse, and display RSS feeds from multiple sources, making it easy to keep up with your favorite blogs, news sites, and podcasts—all in one place.

## Features

- **User Management**: Create users with API key authentication
- **Feed Management**: Add and manage RSS feeds
- **Feed Following**: Users can follow/unfollow feeds
- **Automatic Scraping**: Background service that periodically fetches RSS feeds
- **Post Aggregation**: Collect and store posts from followed feeds
- **RESTful API**: Clean HTTP API with JSON responses
- **Docker Support**: Containerized deployment with Docker Compose

## Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Docker (optional)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/kpriyanshu2003/go-rss-aggregator.git
cd go-rss-aggregator
```

2. Install dependencies (if any):

```sh
make deps
```

3. Set up environment variables:

```sh
cp .env.example .env
```

4. Set up the database:

```sh
# Create database and run migrations
make migrate-up
```

5. Run the application:

```sh
make run
```

### Development

For development with hot reload:

```sh
go install github.com/cosmtrek/air@latest
make dev
```

### Running the Aggregator

Build and run the application:

```sh
go run main.go
```

Or build a binary:

```sh
go build -o rss-aggregator
./rss-aggregator
```

## Configuration

Configure the application using environment variables:

| Variable | Description                  | Default  |
| -------- | ---------------------------- | -------- |
| `PORT`   | Server port                  | `8080`   |
| `DB_URL` | PostgreSQL connection string | Required |

## Usage

- Add feed URLs via the web UI, API, or configuration file.
- Access your aggregated feed at the provided endpoint or in the console.
- Schedule feed refreshes as needed.

## API Docs

- [Postman Collection](https://www.postman.com/infinitybridge/workspace/public-projects/collection/29585525-361a9a16-70cf-444b-87d2-db486c20f767?action=share&source=copy-link&creator=29585525&active-environment=fac7a35e-2db6-45cc-97a6-81e8d0a93501)
- [api doc](./api-doc.md)

## Contributing

Pull requests are welcome! For significant changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Acknowledgments

- Built with [Gin](https://github.com/gin-gonic/gin) web framework
- Database queries generated with [sqlc](https://github.com/kyleconroy/sqlc)
- Uses [UUID](https://github.com/google/uuid) for unique identifiers# README.md

Happy aggregating! 🚀
