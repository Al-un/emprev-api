# EmpRev API <!-- omit in toc -->

- [Run the API](#run-the-api)
  - [Environment variables](#environment-variables)
- [Project structure](#project-structure)
  - [File and folder structure](#file-and-folder-structure)
  - [Service layer](#service-layer)
- [Deployment to GCP](#deployment-to-gcp)

## Run the API

```
git clone https://github.com/Al-un/emprev-api.git
cd emprev-api.git

go run main.go
```

The default configuration assumes that a local Mongo DB is up and running and is listening to port 27017 (default port).

### Environment variables

Available environment variables are:

- `PORT` : port number to have the API listening to
- `SECRET_PWD_SALT` : customize the password salt
- `SECRET_JWT_KEY` : customize the secret used to encrypt JWT
- `DB_URL` : MongoDB URL. The URL must include the database name: `mongodb://<username>:<password>@<host>/<dbname>`
- `DB_NAME` : the database name

## Project structure

### File and folder structure

This project follows the [Standard Go project layout](https://github.com/golang-standards/project-layout) as much as possible. The `main.go` should be in `cmd/api/` (for example) but is located at the root to faciliate GCP deployment. Within `internals/` each package represents a "module". Each module, when involving a business concept (_"users"_ and _"reviews"_) is meant to follow a _Service layer pattern_ approach

### Service layer

Each module is divided into specific layer with a dedicated responsibility per layer:

- _Models (Data model layer)_ defines the different entities involved in the module. Some entities might require entities from other modules
- _DAO (Persistence layer)_ communicates with the database. No ORM is used so the connection is done with a MongoDB driver
- _Handlers (Service layer)_ handles incoming requests and call the appropriate DAO methods
- _Routers_ maps each handler to an endpoint path and method. This layer also define the authentication and authorization requirements to protect some endpoints

## Deployment to GCP

Some notes about the deployment to GCP (with the help of Google docs and [this article](https://medium.com/google-cloud/cloud-build-golang-app-engine-36e27ba976cd)). The API is currently deployed on Google App Engine.

- Prepare a Google cloud account
- Install the GO SDK:
  ```sh
  sudo apt-get install google-cloud-sdk-app-engine-go
  ```
- Create a project:
  ```sh
  gcloud projects create emprev-api --name=emprev-api
  ```
- Create a `app.yaml`:

  ```yaml
  automatic_scaling:
    max_instances: 1

  env_variables:
    DB_URL: mongodb+srv://<username>:<password>@<host>/<dbname>
    DB_NAME: <dbname>
    SECRET_PWD_SALT: <some super salt>
    SECRET_JWT_KEY: <some suer key>

  instance_class: F1

  runtime: go113
  ```

- Deploy!

  ```sh
  gcloud app deploy --appyaml=app.yaml --project=emprev-api
  ```

  > The version field (`--version=1-0-0`) can be defined. Service field, cannot be overriden.

- If prompted, allow Billing and Cloud Build API
