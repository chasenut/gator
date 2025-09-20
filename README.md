# RSS Feed AggreGATOR

RSS Feed Aggregator (Gator) is a cli tool to collect rss feeds.

### Config

In order to work, you may have to create `.gatorconfig` file in your `$HOME` directory.
The file should include at least the `db_url`, which is your **connection string** with 
additional `sslmode=disable` query added to it - application code needs to know to not try to use SSL locally.
```
{
    "db_url":"<YOUR_DB_CONNECTION_STRING>?sslmode=disable",
}
```

# PostgreSQL

Gator uses **PostgreSQL** database to store its information about 
users, feeds, and other surely important stuff.

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

TODO: 
- make some script for it?
- add for ubuntu (maybe)

### Migrations

For migrations I can recommend using [Goose](https://github.com/pressly/goose). 
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
