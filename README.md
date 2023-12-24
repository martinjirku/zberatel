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

## Web layer

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

### Translations - i18n

- i18n is handled by the library [go-playground/universal-translator](https://github.com/go-playground/universal-translator)
- locales: [go-playground/locales](https://github.com/go-playground/locales)

The goal is support at least two language mutations - english and slovak. 

### Forms

Form handling is very important.
 - csfr protection - [justinas/nosurf](https://github.com/justinas/nosurf)
 - to bind form values to structs - [go-playground/validator](https://github.com/go-playground/form)
 - to validate structs - [go-playground/validator](https://github.com/go-playground/validator/v10)

 ### env and flags

 - env file is read by the library [joho/godotenv](https://github.com/joho/godotenv)
 - flags are read by the standard library

 ### db

 For db we are using the postgres.
 - postgre driver [lib/pq](github.com/lib/pq)
 - schema migration is done by dbmate (install it by brew). To run migration

## Usecases

- As collector, I want to share some information, such name/nick, preferences.
- As collector, I want to be able to create a Collection, which is list of collectibles owned, or wished by me. A collection has name, description and type of the collectibles.
- As collector, I want to share the collection with my family, friends.
- As family member, I want to identify, which collectible I can buy as a gift
- As collector, I want to see, which collectibles are missing from the collectible set in my collection
- As collector, I want to explore different collectible sets.
- As collecter, I want to track basic properties of the collectible (name, serial number, acquisition date, its value, condition, link,...)

## Domain

The goal is to support usercases and its entities:

__Collector__: Represents the users who are collectors. This entity can store information about the collector, like their name, contact information, and preferences.

__Collection__: This entity represents the entire collection owned or wished for by a collector. It can include attributes like collection name, description, and the type of items (e.g., coins, stamps).

__Collectible__: Each individual item in a collection. Attributes might include item name, description, acquisition date, value, and condition.

__CollectibleSet__: Represents a set or series within a collection. For example, a set of coins from a specific era or a series of Hot Wheels cars. This entity could have attributes like set name, theme, and total items in the set.

__Wishlist__: An entity for items that the collector wishes to acquire. Attributes can include desired item names, preferred conditions, and target acquisition dates.

__ItemCondition__: A supplementary entity that details the condition of each CollectibleItem, with attributes like condition rating, details, and date assessed.

__Category__: Represents different categories of collectibles, such as coins, stamps, etc. This helps in classifying CollectibleItems and sets.