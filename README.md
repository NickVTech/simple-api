## Project: Creating REST API

GIN: Web framework that makes routing really easy

Viper: Working with enviroment variables to keep secrets mine

GORM: Allows object relations, so I can deal with users.GetById(1) instead of "SELECT id, name... FROM users WHERE id = 1"

JWT-GO: Go implementation of JSON Web Tokens

## App Structure

1. Run docker database container
2. Load enviroment variables    => loadEnv.go
    - Uses Viper to access app.env
3. Provides connection to database          => connectDB.go
    - Uses GORM to make DB object
    - Uses postgres.Open(dsn) driver to connect to db
4. main.go => init()
    - Connect to database
    - Start gin router
5. main.go => main()
    - Creates api group
    - Sets up routes
        - /healthchecker 
        - /GetUsers => Returns all users

## Starting / Stopping container
- docker-compose up -d
=> runs docker container in background

- docker-compose down
=> deletes container

## login to container:

- docker exec -it <container name> bash

- docker              - the cml tool

- exec                - subcommand to execute command inside container

- -it                 - interactive, pseudo-tty. they allocate a terminal and enable interactive shell session

- <container name>    - the name of the container

- bash                - starts bash shell


### steps taken:

1. setup project

docker-compose.yml
=> this will be used to create the docker container, that hosts the database

initializers

loadenv.go
=> I made a config structure, then make a func loadconfig that took in the path
to the location of the dir, and loaded in the app.env file as config object using viper

connectdb.go
=> using the config object, I created the data source name (dsn) to which
the open database connectivity (odbc) driver needs to connect to.
then I used gorm (go object relational mapping) to change the sql query
into a nice little object <3

2. data modeling / migration

models/user.model.go
=> I created a user struct with a uuid, name, email, createdat, and updatedat

I then installed the uuid (universally unique indentifier)
I needed the uuid oosp plugin, because I am using the uuid_generate_v4() function
as a default value on the id column

after, I added pgadmin to the docker-compose, so I can have
a nice webpage to interact with the database.

3. create golang server w/ gin

go get github.com/gin-gonic/gin
=> this is the gin router

go install github.com/cosmtrek/air@latest
=> the air package allows for hot-reloads on file changes

main.go
=> Initializes the database connection, and the Gin router
=> Groups and creates the API routes



Credit: https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/