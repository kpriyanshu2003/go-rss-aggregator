# Go RSS Aggregator

A fast and lightweight RSS aggregator built with Go. This project allows you to fetch, parse, and display RSS feeds from multiple sources, making it easy to keep up with your favorite blogs, news sites, and podcasts—all in one place.

## Features

- 🚀 **High Performance**: Built entirely in Go for speed and efficiency.
- 📡 **Multiple Feed Support**: Aggregate updates from several RSS feeds.
- 🗂️ **Simple Feed Management**: Add, remove, and organize your feeds easily.
- 🕒 **Scheduled Fetching**: Automatically refreshes feeds at configurable intervals.
- 🖥️ **Web Interface/API**: (Optional) Access your aggregated content via a web frontend or RESTful API.
- 📝 **Clean Codebase**: Modular, idiomatic Go code, easy to extend.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### Installation

Clone the repository:

```sh
git clone https://github.com/kpriyanshu2003/go-rss-aggregator.git
cd go-rss-aggregator
```

Install dependencies (if any):

```sh
go mod tidy
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

### Configuration

- Configure your feed sources and settings in the `config.yaml` (if provided) or via environment variables.
- Example feeds can be added in the config or via CLI/API (see usage instructions).

## Usage

- Add feed URLs via the web UI, API, or configuration file.
- Access your aggregated feed at the provided endpoint or in the console.
- Schedule feed refreshes as needed.

## API Docs

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

## Acknowledgements

- [Go](https://golang.org/)

---

Happy aggregating! 🚀
