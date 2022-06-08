# Efishery Authentication Apps

This apps is for authentication user 

## Config

You can copy file `config.json.example` into `config.json` for each apps

## Application
This Apps consist of 2 parts

1. Authentication apps (mandatory)
2. Fetch apps (mandatory)

## Run the application locally

For Auth apps, you can go inside efishery-auth folder and run below command

```
$ cd efishery-auth

$ npm install

$ npm run start
```

For Fetch apps, you can go inside fetch-app folder and run below command

```
$ cd fetch-app

$ go run main.go fetchapp
```

## Open API Docs
For open API documentation both for efishery-auth and fetch-apps can be found at {base_application_url}/swagger

## Generating Swagger

### Authentication apps
For generating swagger Authentication apps you need to modify swagger.json files

You can try to edit and generate `swagger.json` files on this site `https://editor.swagger.io/`

### Fetch apps
For generating swagger Fetch apps you need to modify or add each Comments above router func

For generating swagger automaticly, you can use below command in main project folder

```
$ swag init -g controller/restapi/router.go
```

If having problem with `swag: command not found`

```
$ export PATH=$(go env GOPATH)/bin:$PATH
```

Please refer to this documentation `https://github.com/swaggo/gin-swagger`