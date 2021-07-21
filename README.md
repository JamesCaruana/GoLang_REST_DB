# GoLang REST API


This is a GoLang REST Api which can able to save data to a MongoDB database. This application was tested by hosting a docker container of a MongoDB database locally listening on port 28017.

## Prerequisites
As specified above I had a MongoDB container and also using Postman ([Go_Api.postman_collection.json]( https://github.com/JamesCaruana/GoLang_REST_DB/blob/main/Go_Api.postman_collection.json )) to test/run my code. Below are the following commands used to setup the container and to run my project.
```bash
- docker pull mongo:latest
- docker run --name mongo-db-container -p 27017:27017 -p 28017:28017 -d mongo --port=28017
- go mod download
- go build
- ./main
```

## Libraries used
* [github.com/gorilla/mux]( https://pkg.go.dev/github.com/gorilla/mux )	<- Used to implement a request router and dispatcher
* [go.mongodb.org/mongo-driver]( https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.7.0 )	<- Used to provide a MongoDB Driver API for Go.

## Future Enchancements
<ul>
	<li>Dockerise the whole project and within the docker image create the MongoDB database (using multi-stage containers)</li>
	<li>Test cases</li>
	<li>Using Mapbox API to fetch location data based off the user's address</li>
</u>
