This is a simple Golang-based API-server for 
[this test assignment](https://github.com/avito-tech/adv-backend-trainee-assignment) 

# Deployment

In order to deploy stack with application and database, download 
source and run `docker-compose -f ./deployments/docker-compose.yml up -d` in the source root directory

# Configuration

You can configure a database with adding environment variables in file
`configs/dotenv/postgres.env`, for additional info read the "Environemnt variables"
section of [this instruction](https://hub.docker.com/_/postgres). By default,
there is one variable, `POSTGRES_PASSWORD`, so the other variables have default 
values.

You also can configure API server settings by editing the `configs/json/config.json`
file. In this file, you can change connection string of the database
and a bind address of the server. For more information about the connection string,
read [this instruction](https://github.com/go-gorm/postgres)