connection string: postgres://postgres:postgres@localhost:5432/gator

<h2> Requirements </h2>
1. Postgres v15+
2. Go 1.25+

<h2> Installation </h2>
Gator can be installed by using the command ```go install github.com/mr-rambling/gator``` in the command line

<h2> Use </h2>
A config file, ```.gatorconfig.json``` must be created at your home directory.
It should contain the following: 
```{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}```

Where the "db_url" is the connection string to the database in the format: 
```protocol://username:password@host:port/database```

<h3> Commands </h3>
The program is used by first registering a new user. The user can then add RSS feeds to follow, and to allow other users to follow. By using the browse command a user can see the latest posts by their followed feeds.

<h4> Reset </h4>
	Used in the format ```go run . reset``` to clear the database
<h4> Login </h4>
    Used to login an existing user in the format ```go run . login <user>```
<h4> Register </h4>
    Used to register a new user in the format ```go run . register <user>```
<h4> Users </h4>
    Used to list users in the format ```go run . users```
<h4> Aggregate </h4>
	Used to aggregate registered RRS feeds in the database. Command format ```go run . agg <refresh_time>```
<h4> Add Feed </h4>
    Used to add a new RSS feed to the database in the format ```go run . addfeed <name> <url>```
<h4> Feeds </h4>
    Used to list registered RSS feeds in the format ```go run . feeds```
<h4> Follow </h4>
    Used to follow a feed as the currently logged in user in the format ```go run . follow <url>```
<h4> Following </h4>
    Used to list all feeds being followed by the currently logged in user in the format ```go run . following```
<h4> Unfollow </h4>
    Used to unfollow a feed as the currently logged in user in the format ```go run . unfollow <url>```
<h4> Browse </h4>
    Used to list the latest posts of feeds followed by the logged in user in the format ```go run . browse <number_of_posts>```