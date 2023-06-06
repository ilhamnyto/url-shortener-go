# url-shortener-go
Build a URL Shortener with Go

Stack that i use:
- Echo
- PostgreSQL

## Create User Table
```
CREATE TABLE IF NOT EXISTS users (
            id serial primary key,
            username varchar(100) NOT NULL,
            email varchar(100) NOT NULL,
            password varchar(100) NOT NULL,
            salt varchar(100) NOT NULL,
            created_at timestamp,
            updated_at timestamp
        )
```
## Create Posts Table
```
 CREATE TABLE IF NOT EXISTS urls (
            id serial primary key,
            user_id int NOT NULL,
            long_url text NOT NULL,
            short_url text NOT NULL,
            created_at timestamp,
            CONSTRAINT fk_urls
            FOREIGN KEY(user_id)
            REFERENCES users(id)
        )
```
Or you can run this script to create both tables but you need to edit the env files.
```
go run cmd/api/main.go
```