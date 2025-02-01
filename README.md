# go-tube
### YouTube Video Fetcher API

This is a simple API that fetches the latest YouTube videos based on a predefined search query and stores them in a database. You can query the stored videos, search by title or description, and get them in paginated responses.

## Features
- Fetches the latest YouTube videos based on a predefined search query.
- Stores video metadata (title, description, thumbnail, url , published time) in a database.
- Supports searching videos by title or description.
- Provides paginated responses for video queries.

## Setup

### 1. Clone the Repo
```bash
git clone https://github.com/ss-loves-AC/go-tube.git
cd go-tube
```

### 2. Add Environment Variables
Create a .env file in the root directory with the following variables:
```bash
YOUTUBE_API_KEY=your_api_key_1,your_api_key_2
PREDEFINED_QUERY=cricket
```

### 3. Use Docker
Build and run the project from go-tube directory using
```bash
docker-compose build
docker-compose up

```


## API Endpoints

### 1. `GET /videos`
Fetches stored videos, sorted by publish date.

- **Query Params:**
  - `page` (optional): Page number (default is 1)
  - `limit` (optional): Number of results per page (default is 10)

#### Example Request:
```bash
curl "http://localhost:8080/videos?page=1&limit=10"
```
#### Example Response:
```bash
{
  "page": 1,
  "limit": 10,
  "data": [
    {
      "video_id": "123abc",
      "title": "Golang Tutorial",
      "description": "Learn Golang in 10 minutes",
      "published_at": "2025-02-01T12:00:00Z",
      "thumbnail": "https://img.youtube.com/vi/123abc/default.jpg",
      "url" : ""
    },
    ...
  ]
}
```

### 1. `GET /search`
Search for videos by title or description.

- **Query Params:**
  -  `query` (required): Search term (e.g., golang)
  - `page` (optional): Page number (default is 1)
  - `limit` (optional): Number of results per page (default is 10)

#### Example Request:
```bash
curl "http://localhost:8080/search?query=golang&page=1&limit=10"
```
#### Example Response:
```bash
{
  "page": 1,
  "limit": 10,
  "data": [
    {
      "video_id": "456def",
      "title": "Advanced Golang",
      "description": "Golang best practices",
      "published_at": "2025-02-01T14:00:00Z",
      "thumbnail": "https://img.youtube.com/vi/456def/default.jpg"
    },
    ...
  ]
}
```
