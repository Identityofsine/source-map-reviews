# Fofx's Go GinAPI Template
This is a template for streamlining the development of RESTful APIs using Go and Gin. It includes a pre-configured project structure, essential middleware, and a basic setup for database interactions.

Somethings are a bit primitive, such as the DTOs, but they can be easily extended to suit your needs.

## Features
- RESTful API structure using **Gin**
- Middleware for stored logging, error handling, and CORS
- Basic database setup using Docker, Postgres, and Goose Migrations.
- Generic endpoints for server health checks and versioning
- Simple Permission system for user roles and permissions
- Basic DTOs for request and response handling, with a simple Database interface
- Notification system for sending messages to users, which also includes a simple interface to extend from.
- Authentication and Authorization using JWT and Google OAuth2; with a simple interface to extend from.
- In-depth configuration utilities for managing environment variables and application settings using both `.env`s and various `.yaml` files. (located in `./config`)

## Getting Started
--- 

### Prerequisites
- Docker and Docker Compose

### Clone the repository
```bash
git clone https://github.com/identityofsine/fofx-go-gin-api-template.git
```

### Set up the environment
Take a look at the `docker-compose.yml` file and tweak some of the environment variables to suit your needs. You can also create a `.env` file in the root directory, to share environment variables with the Docker containers. The `.env` file is included in this repo, but is not included in the `.gitignore` file, so you can modify it as needed. 

Also take a look at the `config` directory, which contains various `.yaml` files for configurations ranging from server version to authentication settings to cors settings. You can modify these files to suit your needs, and you may be able to add more files and use the Configurable interface found in `src/pkg/config`.

### Build and run the application

This application uses Docker to run the application and the database. You can use the following command to get iit up and running:
```bash
docker compose up --build
```

#### Development Environment
To run the application in a development environment, which means that the application will be reloaded on every change using nodemon, you can use the following command:
```bash
./start-dev.sh
#or
docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build
```


### Access the application
By default, the application will be running on port `3030`. You can access the API at `http://localhost:3030/api/v1/`.

### Run Migrations
If you want to extend the database schema, you can use Goose Migrations. The migrations are located in the `migrations` directory. Migrations are run automatically when the application starts.

To add a new migration, you can make a new file in the `migrations` directory with the current version format, which has to be consistent with the other migrations in other folders. Or you can use the following command to create a new migration:
```bash
goose create -d migrations/path your_migration_script sql
```
or you could use the script `./create-migration.sh` to create a new migration file in the `migrations` directory. This script will automatically create a new migration file with the current timestamp and the name you provide. For example, to create an initial migration, you can run: 

```bash
./create-migration.sh init create_users_table 
```


