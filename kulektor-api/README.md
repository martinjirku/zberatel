# Zberatel

## Prepare environment

- atlas - database migration tools: `brew install ariga/tap/atlas`
- Task - running tasks: `go install github.com/go-task/task/v3/cmd/task@latest`
- Air - watch&reload server `go install github.com/cosmtrek/air@latest`
- Templ - install templ for template rendering `go install github.com/a-h/templ/cmd/templ@latest` 

- copy .env.example to .env `cp ./.env.example ./.env`
    - DB_DOCKER_NAME -> needed for development, used by `scripts/db-start-docker.sh` to create dockername for db
    - DB_NAME -> database name
    - DB_USER -> database user
    - DB_PWD -> database password

## Development

- Start server and watch it. `air`
- Start docker DatabaseStart - postgre by running docker. `task db`

- Generate migration file after change:
  ```sh
  atlas migrate diff ksuid_type \
      --dev-url "docker://postgres/15/dev?search_path=public" \
      --to "file://schema.sql" \
      --format '{{ sql . "  "}}'
  ```
- Run migrations
  ```sh
  atlas migrate apply \
    --dir "file://db/migrations" \
    --url "postgres://$DB_USER:$DB_PWD@$DB_ADDR:$DB_PORT/$DB_NAME?sslmode=disable"
  ```

## Initialize graphl

- `go run github.com/99designs/gqlgen init`
