# Go REST API

A simple REST API in Go, using Gin for routing and middleware, and SQLite for the database.

## What It Does

- User signup and login (JWT auth)
- Create, read, update, and delete events
- Register for an event and cancel a registration

## Endpoints

- `POST /signup`
- `POST /login`
- `GET /events`
- `GET /events/:id`
- `POST /events` (auth)
- `PUT /events/:id` (auth)
- `DELETE /events/:id` (auth)
- `POST /events/:id/register` (auth)
- `DELETE /events/:id/register` (auth)

## Run

Set `JWT_SECRET` in `.env`, then:

```bash
go run .
```

The API runs on `http://localhost:8080` and uses `api.db`.

