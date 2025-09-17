# RSS Feed AggreGATOR

RSS Feed Aggregator (Gator) is a cli tool to collect rss feeds.

TODO: Add more details to readme

# Database

Gator uses **PostgreSQL** server database to store its information about 
users, feeds, and other surely important stuff.

## PostgreSQL setup

Setup walk-through of PosgreSQL on local machine for **Linux distros**, 
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

TODO: 
- make some script for it?
- add for ubuntu (maybe)
