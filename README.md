# gator-rss

A command-line RSS aggregator built in Go, featuring user authentication, feed management, and aggregation backed by PostgreSQL.

## Features

- **User Authentication**: Register and login users, manage session state.
- **PostgreSQL Database**: All data is persisted using a local PostgreSQL database.
- **Feed Management**:
  - Add new feeds (`addfeed`) by providing a name and URL.
  - Follow/unfollow feeds by URL (`follow`, `unfollow`).
  - List followed feeds (`following`), browse posts (`browse`).
- **Aggregation**: Aggregate RSS feeds and store posts in the database (`agg`).
- **Goose & SQLC**:
  - Database migrations managed with [Goose](https://github.com/pressly/goose).
  - Type-safe queries auto-generated with [SQLC](https://github.com/kyleconroy/sqlc).
- **Session & Config**: Stores `db_url` and current user in a simple JSON config file.

## Database Schema

Main tables:

- `users`: Registered users.
- `feeds`: Available RSS feeds.
- `feed_follows`: User-follow relationships for feeds.
- `posts`: Aggregated posts from feeds.

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL instance (local or remote)
- [Goose](https://github.com/pressly/goose) CLI
- [SQLC](https://github.com/kyleconroy/sqlc) CLI

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/youruser/gator-rss.git
   cd gator-rss
   ```

2. Set up the database and apply migrations:
   ```bash
   # Create the database (change credentials as needed)
   createdb gator_rss
   goose -dir db/migrations postgres "postgres://user:pass@localhost:5432/gator_rss?sslmode=disable" up
   ```

3. Generate Go code from queries:
   ```bash
   sqlc generate
   ```

### Configuration

- Place your config in `config.json`:
  ```json
  {
    "db_url": "postgres://user:pass@localhost:5432/gator_rss?sslmode=disable",
    "current_user": "alice"
  }
  ```

## Usage

All commands are invoked through the CLI.

### Register a new user

```bash
gator-rss register alice
```

### Login

```bash
gator-rss login alice
```
*Sets the current user in the config.*

### Add a new feed

```bash
gator-rss addfeed "Go Blog" https://blog.golang.org/feed.atom
```

### Follow a feed

```bash
gator-rss follow https://blog.golang.org/feed.atom
```

### List followed feeds

```bash
gator-rss following
```

### Unfollow a feed

```bash
gator-rss unfollow https://blog.golang.org/feed.atom
```

### Aggregate feeds

```bash
gator-rss agg
```
*Fetches new posts from followed feeds and stores them in `posts`.*

### Browse posts

```bash
gator-rss browse 10
```
*Shows up to 10 recent posts from feeds the current user is following.*

## Development

- Migrations: `goose -dir db/migrations postgres "$DB_URL" up`
- Database models and queries: `sqlc generate`

## Design Notes

- The application stores the `db_url` and current user in a JSON config, parsed into a Go struct.
- All user actions and state are managed through CLI commands.
- Posts are only visible in `browse` for feeds the current user follows.

## License

MIT — see [LICENSE](LICENSE).

Contributions welcome! Please open issues or submit PRs for improvement or bugfixes.