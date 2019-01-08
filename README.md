# Go Docker API Boilerplate

> Boilerplate for Go based REST API with PostgreSQL, with live reload using CompileDaemon.

Builds a docker container for PostgreSQL Database.
Builds a docker container with live reload for Go REST API and links to Postgres container
Builds a zip package with binary for deployment to AWS ElasticBeanstalk

## Table of Contents

- [Usage](#usage)
- [Maintainers](#maintainers)
- [License](#license)

## Install

Check for dependancies

```
go get github.com/tools/godep
dep ensure
```

## Usage

1. Create your own .env using .env-sample

2. Start docker containers

```
make up
```

3. View logs

```
make logs
```

4. Stop docker containers

```
make down
```

## Deploy

1. Create .env.prod using .env.sample

2. Build binary and zip

```
make deploy-prod
```

3. Upload new .zip in deploy/ to AWS EB instane

## Maintainers

[@gpng](https://github.com/gpng)

## License

MIT © 2018 Gerald Png
