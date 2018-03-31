lagerauth main repo

This repo contains the server/api for the oauth implementation for lagersoft.


### Project Structure:

- config: code that reads and parses conf.json
- handlers: general http handlers.
- handlers/api: api resource handlers.
- handlers/middleware: helpful middleware for http
- models: structs for requests and responses.
- logic: business logic, this should be the only access to `database` package
- database: code for querying, should use models and primitives for params and returns.
- logger: leveled logger package, logs to db table `logs` and to `os.Stdout` (at the same time)

### Routes for Pemissions Api:

- `can` (POST): Input: JSON -> `{Controller, Action, Method}`, Response -> HTTP STATUS: `200` for OK, `401` for Unathorized (no token present), `403` for when user is not authorized (but is authenticated) to the resource.

### Routes for Internal Api:

All routes to api are in the form: `/api/<resource>`
Every request to `/api` gets authenticated with the `Authentication` Header, right now it only supports type `Bearer` if you dont have that header for the requests it will return an `HTTP 401 Unauthorized` error, if you dont have access to that `resource` it will return `HTTP 403 Forbidden`.
All interactions with the api are trough `JSON`

Sopported Routes for `resource`:
- `GET` -> `/api/resource` : List of all `resource`.
- `GET` -> `/api/resource/id`: Data for `resource/id`.
- `POST` -> `/api/resource`: Create `resource` from body data.
- `PUT` -> `/api/resource/id`: Updates `resource/id` with the values from the body data.
- `DELETE` -> `/api/resource/id`: Deletes `resource/id`.

### Deploy:
- Build the site project, copy it into `wwwroot` folder.
- Configure `conf.json` with correct parameters, database should be mysql but its probably not hard to change into any other database since we use gorm for querying.
- `go get`
- `go build`
- Run it, remember to copy the `conf.json` file and `wwwroot/` folder to where the binary is located.
