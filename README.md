# Zberatel

## Prepare environment

- dbmate - database migration tools: `brew install dbmate`
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


## Architecture

Architecutre of the rendering is mix of templ templates. Each templ should have its view-model struct. E.g. Login will have LoginVM. This is naming convetion, but it increase simplicity. The downside is, that each templ component needs to define its own view-model struct.

The templates folder contains 4 packages:
  - _components_ - this is place to store atomic components (such as buttons, inputs,...)
  - _layout_ - place for specifying different layouts. The basic and required layout is the page layout. It is usually the root layout.
  - _page_ - this is the place for high level components which will be mapped to endpoints
  - _partials_ - partials are higher level compoennts, which are composing more components, and contain more business logic.

The templ templates are just rendering layer, the tempaltes needs to be rendered by handlers. Handler should render the content section. The simple example, looks like this:

```go
func AuthRegisterHandler(w http.ResponseWriter, r *http.Request) {
	content := page.Register(page.NewRegisterVM())
	layout.Page(layout.NewPageVM("Login")).Render(templ.WithChildren(r.Context(), content), w)
}
```