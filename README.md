# Blog Agreggator

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=plastic&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=plastic&logo=postgresql&logoColor=white)

## Overview
This project is a backend server implemented in Go, using the Chi router, PostgreSQL for data storage, SQLc for SQL queries, and goose for database migrations. The main functionality of the project is to aggregate and manage RSS feeds. Users can add their own RSS feeds, follow and unfollow feeds from other users, and retrieve the latest posts from the feeds they follow.

## Features
- Management: Users can create accounts and manage their profile.
- RSS Feed Management: Users can add their own RSS feeds and follow/unfollow feeds from other users.
- Post Retrieval: Users can fetch the latest posts from the RSS feeds they follow.
- API Endpoints: The project exposes API endpoints for various actions.

## Tech Used
- Go: Main programming language.
- Chi: Lightweight router for building RESTful APIs in Go.
- PostgreSQL: Database for storing user data and RSS feed information.
- SQLc: SQL compiler for Go that helps manage SQL queries.
- Goose: Database migration tool for managing changes to the database schema.

## Installation
 
1. Clone repository:
```bash
  git clone https://github.com/nmowens95/Blog-Aggregator
```
2. Install dependencies:
```bash
    go get -d ./...
```
3. Set up the PostgreSQL database:
```bash
  createdb rss_aggregator
```
4. Create a .env file, you'll need to store your port number & your database URL:
```bash
 PORT=8000
```
This will also serve as your CONN for your goose migration:
```bash
  DB_URL=protocol://username:password@host:port/database
```
5. Run goose database migration:
```bash
  goose CONN up
```
6. Build and run server:
```bash
  go build -o out && ./out
```
## Testing
This project currently runs on the backend so to test you will want to use something like:
- [Postman](https://www.postman.com/)
- [Thunderclient](https://www.thunderclient.com/)

#### User Endpoints
- `Post /v1Router/users`: Create user
- `Get /v1Router/users`: Get user

#### Feed Endpoints
- `Post /v1Router/users`: Create feed
- `Get /v1Router/users`: Get feed

#### Feed Follows Endpoints
- `Post /v1Router/feed_follows`: Create feed follow
- `Get /v1Router/feed_follows`: Get all feed follows
- `Delete /v1Router/feed_follows/{feedFollowID}`: Delete a feed

#### Posts Endpoints
- `Get /v1Router/posts`: Get posts
    
## Future Improvements
- Support pagination of the endpoints that can return many items
- Support different options for sorting and filtering posts using query parameters
- Classify different types of feeds and posts (e.g. blog, podcast, video, etc.)
- Add a CLI client that uses the API to fetch and display posts, maybe it even allows you to read them in your terminal
- Scrape lists of feeds themselves from a third-party site that aggregates feed URLs
- Add support for other types of feeds (e.g. Atom, JSON, etc.)
- Add integration tests that use the API to create, read, update, and delete feeds and posts
- Add bookmarking or "liking" to posts
- Create a simple web UI that uses your backend API
