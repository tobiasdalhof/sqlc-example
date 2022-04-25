# sqlc-example

This repository contains example code for [sqlc](https://github.com/kyleconroy/sqlc) with Go.

## Todo

- [ ] Add example that demonstrates the usage of sqlc generated code
- [ ] Add database test example that uses [dockertest](https://github.com/ory/dockertest)

## Generate Go code from SQL

```bash
make sqlc
```

## Commands for DB migration

```bash
# create new migration
make migration name=<name>

# run migrations
make migrate_up

# rollback migrations
make migrate_down
```

## Links

- https://github.com/johanbrandhorst/grpc-postgres
- https://github.com/techschool/simplebank
