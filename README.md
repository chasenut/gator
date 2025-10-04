# Gator

RSS Feed Aggregator (Gator) is a cli tool to collect, store, browse and manage online rss feeds.

## Features

Gator features following commands:

- login <username> - log in using your username / switch to user
- register <username> - create new user with your username
- reset - resets everything
- users - prints all users
- agg <duration> - fetches and updates followed feeds in a continuous loop (with a little break between)
- addfeed <name> <url> - add new feed to the collection (auto follows it)
- feeds - prints all feeds in a collection
- follow <url> - allows user to follow a feed
- following - prints all feeds followed by the user
- unfollow <url> - allows user to unfollow a feed
- browse [limit] - prints all posts from the feeds followed by the user


# Setup and build

Here is how to get your Gator working.

### Config

In order to work, you may have to create `.gatorconfig` file in your `$HOME` directory.
The file should include at least the `db_url`, which is your **connection string** with 
additional `sslmode=disable` query added to it - application code needs to know to not try to use SSL locally.
```
{
    "db_url":"<YOUR_DB_CONNECTION_STRING>?sslmode=disable",
}
```


Gator uses **PostgreSQL** database to store its information about 
users, feeds, and other surely important stuff. Database is set locally, 
but you can probably also use a remote server.

## PostgreSQL

Here is a setup walk-through of PosgreSQL on local machine for **Linux**, 
if you use MacOS, it might be similiar, if you use Windows, 
grow some hair on your chest.

### Arch

Install PostgreSQL:
```
sudo pacman -S postgresql
```
Configuration, set your password (follow output instructions):
```
sudo passwd postgres
```
Start service:
```
sudo systemctl start postgresql
```

### Connection string

It might be helpful to know what the **connection string** looks like. 
The format is: `protocol://username:password@host:port/database`.

On **Linux**, given the username `postgres`, password `postgres` and database `gator`,
it looks like this (most likely): `postgres://postgres:postgres@localhost:5432/gator`.

Test your connection by running `psql`, for example:
```
psql "postgres://postgres:postgres@localhost:5432/gator"
```

# Development

### Migrations

For migrations I use [Goose](https://github.com/pressly/goose). 
You can install it following the install process at source, or use:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

For any migrations (sql files inside `sql/schema`) use the following command:
```
goose postgres <connection_url> <up/down>
```

### SQLC

[SQLC](https://sqlc.dev/) is used to compile SQL code to GO.
Install it with the following command:
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

To use, in the root of your repo type the following command:
```
sqlc generate```


## Credit

Project created as a part of [boot.dev](https://boot.dev) course.
