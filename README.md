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
git clone https://github.com/ss_loves_AC/go-tube.git
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

