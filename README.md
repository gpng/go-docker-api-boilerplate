# Go Docker API Boilerplate

> Boilerplate for Go based REST API with PostgreSQL, with live reload using CompileDaemon.

Builds a docker container for PostgreSQL Database.
Builds a docker container with live reload for Go REST API and links to Postgres container

## Table of Contents

- [Usage](#usage)
- [Maintainers](#maintainers)
- [License](#license)

## Usage

1. Create .env using .env.sample as example

1. Start docker containers

    ```
    make
    ```

1. View logs

    ```
    make logs
    ```

1. Visit `localhost:4000/` to check if API is responding

1. Generate docs from swagger comments

    ```
    make gen-docs
    ```

1. Visit `localhost:4000/docs` for documentation

1. Stop docker containers

    ```
    make down
    ```

## Migrations

[Install goose](https://github.com/pressly/goose)

Create new migrations

```
goose -dir sqlc/schemas create <migration_name> sql
```

Run migrations

```
env $(cat .env) make migrate
```

Rollback migrations

```
env $(cat .env) make rollback
```

## Generating models

[Install sqlc](https://dl.equinox.io/sqlc/sqlc/devel)

```
make gen-models
```


## Test

```
make test
```

## Maintainers

[@gpng](https://github.com/gpng)

## License

MIT © 2020 Gerald Png
