# basic-api-backend

This is a demo application for skaffold usage.

# API definition

Build an application that serves the following HTTP-based APIs :

Description: Saves/updates the given user's name and date of birth in the database
Request: PUT /hello/Morty { "dateOfBirth": "2000-01-01" }
Response: 204 No Content


Description: Return a hello/birthday message for the given user
Request: GET /hello/Morty
Response: 200 OK
* when Morty’s birthday is in 5 days:
{ "message": "Hello, Morty! Your birthday is in 5 days" }
* when Morty's birthday is today:
{ "message": "Hello, Morty! Happy birthday" }

## Package, Build App 

To build the app, from the root of the repo

make sure have go installed (https://golang.org/doc/install)
```
go build
```

To package the application ad docker image

make sure have docker installed (https://docs.docker.com/install/)

```
docker build -t yahyaozturk/basic-api .
```

OR
Install skaffold from (https://skaffold.dev/docs/install/)

```
skaffold build
```

## Deploy / Run Locally

To run, from the root of the repo

make sure have go installed (https://golang.org/doc/install)
make sure have MongoDB instance and set following ENV

MONGOHOST  = mongodb hostname
MONGOUSER  = mongodb username
MONGOPASSWORD  = mongodb password

if not have mongodb, you can provision one by docker 

```
docker run -itd --name mongodb -p 27017:27017 --env MONGO_INITDB_ROOT_USERNAME=admin --env MONGO_INITDB_ROOT_PASSWORD=password --env MONGO_INITDB_DATABASE=testdb mongo
```

```
go run .
```

OR

```
go build
./basic-api-backend
```
OR

make sure have docker installed (https://docs.docker.com/install/)
```
docker run -itd -p 8080:8080 --env MONGOHOST=<replace> --env MONGOUSER=<replace> --env MONGOPASSWORD=<replace> yahyaozturk/basic-api
```


## Access the app 

The App has a one Endpoint

Api endpoint is prefixed with `/api/v1`

To reach the endpoint use `baseurl:8080/api/v1/hello{name}`

Request: PUT /hello/{name} { "dateOfBirth": "2000-01-01" }

Request: GET /hello/{name}

## Test the app 

Use basic Jmeter script to test put and get operation
`basic-loadtest.jmx`

OR

GET - `curl -v http://{baseurl}:8080/api/v1/hello/{name}`

PUT - `curl -X PUT -H "Content-Type: application/json" -d '{ "dateOfBirth": "2000-02-17" }' http://{baseurl}:8080/api/v1/hello/{name}`

## Deploy app to cloud

skaffold is used for local development `skaffold.yml`

Install skaffold from (https://skaffold.dev/docs/install/)

Make sure have K8 cluster to deploy application, select your context or go to under setup folder and find all required file to setup Azure Kubernetes Cluster for your test

`aks-preparation.sh`
`create-ingress.sh`
`install-mongo-by-helm.sh`
`create-mongo-secret.sh` -- this is mandatory step to save the mongodb info as secret in your cluster

In order to DEV mode `skaffold dev`
In order to RUN mode `skaffold run`

skaffold takes care of baking docker image (`dockerfile`) and deploy to K8 cluster automatically (`k8s-deployment.yml`)

OR 
`create-mongo-secret.sh` -- this is mandatory step to save the mongodb info as secret in your cluster
`kubectl apply -f k8s-deployment.yml` -- installs deployment + service
`kubectl apply -f ingress.yml` -- install ingress, makesure to chance ingress IP in YML file
`kubectl apply -f hpa.yml` -- install horizantal pod auto-scaler to make the app scalable


