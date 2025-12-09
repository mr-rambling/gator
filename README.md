## Requirements
1. Postgres v15+
2. Go 1.25+

## Installation
Gator can be installed by using the command `go install github.com/mr-rambling/gator` in the command line

## Use
A config file, `.gatorconfig.json` must be created at your home directory.
It should contain the following: 

```
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Where the "db_url" is the connection string to the database in the format: 
```protocol://username:password@host:port/database```

### Commands
The program is used by first registering a new user. The user can then add RSS feeds to follow, and to allow other users to follow. By using the browse command a user can see the latest posts by their followed feeds.

#### Reset 
Used in the format `./gator reset` to clear the database
#### Login 
Used to login an existing user in the format `./gator login <user>`
#### Register 
Used to register a new user in the format `./gator register <user>`
#### Users 
Used to list users in the format `./gator users`
#### Aggregate 
Used to aggregate registered RRS feeds in the database. Command format `./gator agg <refresh_time>`
#### Add Feed 
Used to add a new RSS feed to the database in the format `./gator addfeed <name> <url>`
#### Feeds 
Used to list registered RSS feeds in the format `./gator feeds`
#### Follow 
Used to follow a feed as the currently logged in user in the format `./gator follow <url>`
#### Following 
Used to list all feeds being followed by the currently logged in user in the format `./gator following`
#### Unfollow 
Used to unfollow a feed as the currently logged in user in the format `./gator unfollow <url>`
#### Browse 
Used to list the latest posts of feeds followed by the logged in user in the format `./gator browse <number_of_posts>`