# sap-filter-trails

Home assignment for SAP

## Project

This is a web API made with gin.
When launched, it will first parse the Boulder trail CSV file and then listen for incoming requests on port 3001.
At this point you can send some GET requests to filter out the trails.

Here are the available endpoints:

| Method | URL                      | Arguments                                   | Description                               | Error codes                             |
| ------ | ------------------------ | ------------------------------------------- | ----------------------------------------- | --------------------------------------- |
| GET    | /trails                  |                                             | Returns all trails                        |
| GET    | /trails/with-grills      |                                             | Returns all trails with grills            |
| GET    | /bike-trails/            |                                             | Returns only bike trails                  |
| GET    | /bike-trails/with-picinc |                                             | Returns bike trails with picnic available |
| GET    | /trails/by-name/:name    | `:name`: string, name of the required trail | Returns the trail named `name`            | `404` if no trail with `name` was found |

## Building and Running with docker

First build the docker image with:

`docker build -t <TAG> .`

Then run it with

`docker run -it --rm -p 3001:3001 --name=<NAME> <TAG>`

## Building and Running locally for debugging

First get the dependencies with

`go get`

Then run the project:

`go run .`

## Things I would change

This was my first "real" golang project. I had toyed with it before but have very "surface-level" knowledge of it.

At first I wanted to use the geojson file provided by the Boulder county, but I quickly realised it contained way less data than the csv that was provided in the SAP repository, which is why I went with it.

Splitting the project in modules came way too late in the development process. I would start with that next time, that would allow a better separation of concerns and would allow me to design the project with the MVC design pattern.

I would also add unit testing.

I'd like to add some environment variables to use a customized port, to launch the Docker container in debug mode, and so on.
