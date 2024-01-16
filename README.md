# Blog Agreggator

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=plastic&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=plastic&logo=postgresql&logoColor=white)

#### RSS Feed 
- Allows you to add feeds to be collected from a Blog
- There is an ability to follow/unfollow a feed
- Use this project to keep up to date with your favorite blogs


## Installation
 
To clone repo:
```bash
  git clone https://github.com/nmowens95/Do-Go-To-Learn-Go
```
You will need to install dependencies:
```bash
    go get -u github.com/go-chi/chi/v5
    go get github.com/go-chi/cors
    go get github.com/joho/godotenv
```
This project currently runs on the backend so to test you will want to use something like:
- [Postman](https://www.postman.com/)
- [Thunderclient](https://www.thunderclient.com/)

    
## Lessons Learned


- Working with SQL and making data migrations 
- Working with XML to reed Blog feed
- Using ApiKeys as authentication for users 
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
