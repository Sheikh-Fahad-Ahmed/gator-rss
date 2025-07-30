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
  - Database migrations managed with [Goose](https://github.com/pressly/goose).
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

git clone https://github.com/youruser/gator-rss.git
cd gator-rss
text

2. Set up the database and apply migrations:

Create the database (change credentials as needed)

createdb gator_rss
goose -dir db/migrations postgres "postgres://user:pass@localhost:5432/gator_rss?sslmode=disable" up
text

3. Generate Go code from queries:

sqlc generate
text

### Configuration

- The application uses a JSON config file named `.gatorconfig.json` stored in your home directory.

- This config file stores:

- `db_url`: The PostgreSQL database connection URL.
- `current_user`: The username currently logged in.

- Example `.gatorconfig.json` content:

{
"db_url": "postgres://user:pass@localhost:5432/gator_rss?sslmode=disable",
"current_user": "alice"
}
text

- The config file is created and updated automatically by the application when you log in or register.

- **Ensure PostgreSQL is installed and running** since the `db_url` must point to a valid database instance.

## Usage

All commands are invoked through the CLI.

### Register a new user

gator-rss register alice
text

### Login

gator-rss login alice
text

*Sets the current user in the `.gatorconfig.json` config file.*

### Add a new feed

gator-rss addfeed "Go Blog" https://blog.golang.org/feed.atom
text

### Follow a feed

gator-rss follow https://blog.golang.org/feed.atom
text

### List followed feeds

gator-rss following
text

### Unfollow a feed

gator-rss unfollow https://blog.golang.org/feed.atom
text

### Aggregate feeds

gator-rss agg <time_between_request>
text

*Fetches new posts from followed feeds and stores them in the `posts` table.*  
*The `<time_between_request>` argument specifies how long to wait between fetching each feed (e.g., `1m` for 1 minute, `1s` for 1 second, `1h` for 1 hour). This helps avoid overwhelming the feed servers.*

### Browse posts

gator-rss browse 10
text

*Shows up to 10 recent posts from feeds the current user is following.*

## Development

- Migrations: `goose -dir db/migrations postgres "$DB_URL" up`
- Database models and queries: `sqlc generate`
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