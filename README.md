# Manga Release Tracker API

This is a simple API that scrapes manga release data from the Anime News Network website and provides an endpoint to access the scraped data. The project is built using Golang, Colly (a powerful web scraping framework), and standard HTTP for API handling.

## Features

- Scrapes the latest manga release information from [Anime News Network](https://www.animenewsnetwork.com/encyclopedia/releases.php?format=manga)
- Provides an API endpoint to access the scraped data in JSON format
- Can be containerized with Docker for easy deployment (in progress)
- Terraform scripts for infrastructure deployment (future)

## Tech Stack

- **Golang**: The main programming language used for building the API.
- **Colly**: A web scraping library for scraping manga release data.
- **Docker**: Used for containerizing the API.
- **Terraform**: Infrastructure as Code (IaC) tool for deployment (upcoming).

## Project Structure
```
manga-tracker-api/
│
├── internal/
│   └── manga/
│       └── scraper.go      # Contains the scraping logic
│
├── api/
│   └── handler.go          # API handler functions for the /manga endpoint
│
├── cmd/
│   └── main.go             # Entry point for the application
│
├── Dockerfile              # Dockerfile to containerize the API (not uploaded yet)
├── README.md               # Project documentation
└── go.mod                  # Dependency management
```

## Getting Started

### Prerequisites

- **Go 1.23** installed on your machine
- **Docker** (optional) for containerization

### Installing Dependencies

Clone the project to your local machine:

```bash
git clone https://github.com/yourusername/manga-release-tracker-api.git
cd manga-release-tracker-api
```
Download the required Go dependencies:
```
go mod tidy
```
Running the Application Locally
To run the application, you can use the following command:
```
go run cmd/main.go
```
The server will start on port 3333 by default. You can access the API at:

- Home: http://localhost:3333/
- Manga Releases: http://localhost:3333/manga

###Scraping Logic
The scraping logic is located in the internal/manga/scraper.go file. It scrapes manga release data from Anime News Network and returns it in a structured format.

###Endpoints
1. GET /
- Description: A basic home route that returns a welcome message.
- URL: /
- Method: GET
- Response:
```
{
  "message": "Welcome to the Manga Release Tracker API"
}
```
2. GET /manga
- Description: Fetches the latest manga releases by scraping data from Anime News Network.
- URL: /manga
- Method: GET
- Response: JSON array of manga release data.
```
[
  {
    "title": "Manga Title",
    "pages": "200",
    "distributor": "Distributor Name",
    "releaseDate": "2024-09-01",
    "srp": "$15.99"
  },
  ...
]
```
## Future Features
- Add a database to store scraped manga data
- Deploy the API using Terraform
- Implement unit and integration tests
- Add logging and error monitoring
## Contributing
Feel free to open issues or submit pull requests if you'd like to contribute to the project!



