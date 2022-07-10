# Go todo

### Simple todo list REST JSON API written in go

---

## Setup

Before Running this app you need to setup environment variables (.env file)

> Your .env file shuld look simlar to `.env.example`
> 
> You can even copy contents of this file and project will work

If you want to run only database on docker your `.env` file should look like this

```bash
DB_USER=user
DB_PASSWORD=password
DB_ROOT_PASSWORD=root_password
DB_HOST=localhost
```

If you want to run all components of this app on docker your `.env` should look like this

```bash
DB_USER=user
DB_PASSWORD=password
DB_ROOT_PASSWORD=root_password
DB_HOST=database
```

> Notice the difference in the `DB_HOST` variable

## Running

#### Running with only database on docker (or fully without docker)

To run this app without docker you need to install external modules first with this command

```bash
go mod download
```

Then you need need to have database server running 

You can do it using docker with this command

```bash
docker-compose up --build 
```

*You don't need to do this if you have external mysql or mariadb server running and changed configuration in `.env` file*

Then you run go project using command

```bash
go run main.go
```

#### Running fully on docker

When running fully on focker all you have to do is run this command

```bash
docker-compose -f docker-compose-prod.yml up --build
```

## Testing

When testing the app you need to run the app first in order for integration tests to work

See **Running with only database on docker (or fully wihout docker)** paragraph

---

© @piotreknow02 2022
