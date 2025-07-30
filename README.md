# gator-rss

A command-line RSS aggregator built in Go, featuring user authentication, feed management, and aggregation backed by PostgreSQL.

## Features

- **User Authentication**: Register and login users, manage session state.
- **PostgreSQL Database**: All data is persisted using a local PostgreSQL database. **PostgreSQL must be installed and running for this project to work.**
- **Feed Management**:
  - Add new feeds (`addfeed`) by providing a name and URL.
  - Follow/unfollow feeds by URL (`follow`, `unfollow`).
  - List followed feeds (`following`), browse posts (`browse`).
- **Aggregation**: Aggregate RSS feeds and store posts in the database (`agg`).
- **Goose & SQLC**:
  - Database migrations are placed in the `sql/schema` folder and managed with [Goose](https://github.com/pressly/goose).
  - Type-safe queries auto-generated with [SQLC](https://github.com/kyleconroy/sqlc).
- **Session & Config**: Uses a JSON config file named `.gatorconfig.json` stored in the user’s home directory to save the database URL and currently logged-in user.

## Database Schema

Main tables:

- `users`: Registered users.
- `feeds`: Available RSS feeds.
- `feed_follows`: User-follow relationships for feeds.
- `posts`: Aggregated posts from feeds.

## Getting Started

### Prerequisites

- Go 1.20+
- **PostgreSQL installed and running locally or remotely**
- [Goose](https://github.com/pressly/goose) CLI
- [SQLC](https://github.com/kyleconroy/sqlc) CLI

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/youruser/gator-rss.git
   cd gator-rss
   ```

2. Create the database using the Postgres CLI tool `psql`:

   ```bash
   # You may need to enter your Postgres password
   psql -U your_pg_user -c 'CREATE DATABASE gator;'
   ```

   Replace `your_pg_user` with your actual Postgres username.

3. Apply database migrations:

   ```bash
   goose -dir sql/schema postgres "postgres://user:pass@localhost:5432/gator?sslmode=disable" up
   ```

4. Generate Go code from queries:

   ```bash
   sqlc generate
   ```

### Configuration

- The application uses a JSON config file named `.gatorconfig.json` stored in your home directory.

- This config file stores:

  - `db_url`: The PostgreSQL database connection URL.
  - `current_user`: The username currently logged in.

- Example `.gatorconfig.json` content:

  ```json
  {
    "db_url": "postgres://user:pass@localhost:5432/gator?sslmode=disable",
    "current_user": "alice"
  }
  ```

- The config file is created and updated automatically by the application when you log in or register.

- **Ensure PostgreSQL is installed and running** since the `db_url` must point to a valid database instance.

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

*Sets the current user in the `.gatorconfig.json` config file.*

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

*Fetches new posts from followed feeds and stores them in the `posts` table.*  
*The `` argument specifies how long to wait between fetching each feed (e.g., `1m` for 1 minute, `1s` for 1 second, `1h` for 1 hour). This helps avoid overwhelming the feed servers.*

### Browse posts

```bash
gator-rss browse 10
```

*Shows up to 10 recent posts from feeds the current user is following.*

## Development

- Migrations:  
  ```bash
  goose -dir sql/schema postgres "$DB_URL" up
  ```
- Database models and queries:  
  ```bash
  sqlc generate
  ```
- Config package manages the `.gatorconfig.json` file located in the user’s home directory.

## Design Notes

- Configuration is stored in a JSON file `.gatorconfig.json` located in your home directory.
- This file keeps track of `db_url` and `current_user`.
- All user actions and state are managed through CLI commands.
- Posts are only visible in `browse` for feeds the current user follows.
- PostgreSQL must be installed and running; ensure the connection URL in the config file is valid.

## License

MIT — see [LICENSE](LICENSE).

Contributions welcome! Please open issues or submit PRs for improvement or bugfixes.