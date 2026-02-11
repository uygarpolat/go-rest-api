# Go REST API

A basic REST API built with Go, Gin, and SQLite.

## Features

- User signup and login
- CRUD endpoints for events
- SQLite database auto-initialization on startup
- Password hashing with bcrypt

## Tech Stack

- Go
- Gin (`github.com/gin-gonic/gin`)
- SQLite (`github.com/mattn/go-sqlite3`)

## Project Structure (WIP)

- `main.go` - app entry point
- `db/` - database setup and table creation
- `models/` - data models and DB operations
- `routes/` - HTTP route handlers
- `utils/` - helpers (password hashing, JWT utility)
- `api-test/` - sample HTTP request files

## Setup

1. Clone the repository.
2. Create a `.env` file in the project root:

```env
JWT_SECRET=your_secret_here
```

3. Download dependencies:

```bash
go mod tidy
```

## Run

```bash
go run .
```

Server starts on: `http://localhost:8080`

On first run, the app creates:

- `users` table
- `events` table
- local SQLite database file: `api.db`

## API Endpoints

### Auth

- `POST /signup` - create a user
- `POST /login` - login with email/password

### Events

- `GET /events` - list all events
- `GET /events/:id` - get event by ID
- `POST /events` - create event
- `PUT /events/:id` - update event
- `DELETE /events/:id` - delete event

## Quick Testing

You can use the request files in `api-test/`:

- `create-user.http`
- `login.http`
- `create-event.http`
- `get-events.http`
- `get-single-event.http`
- `update.event.http`
- `delete-event.http`

If your editor supports `.http` files (e.g. VS Code / Cursor REST Client extensions), run those requests directly.

