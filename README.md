<p align="center"><img src="https://user-images.githubusercontent.com/1092882/60883564-20142380-a268-11e9-988a-d98fb639adc6.png" alt="webgo gopher" width="256px"/></p>

# go-boilerplate

Golang microservice boilerplate.

### Directory structure

```bash
|
|____build # This include the build & run commands which helps to run service on local machine using docker.
|    |
|    |
|____docker # Contains the dockerfile.
|    |
|    |
|____src # Contains the source code of project.
|    |
|    |
|____|____apis # Contains the APIs routes (REST, gRPC), middlewares and each route controllers.
|    |
|    |
|    |____cmd # Contains the main file.
|    |     |____main.go # main file.
|    |     
|    |____config # Contains the configuration values.
|    |
|    |
|    |____dal # This is data access layer, here we can define the db interaction code like: reads data from db.
|    |
|    |
|    |____init # Contains the intialisation of all dependecies.
|    |
|    |
|    |____models # Contains the APIs request & response models.
|    |
|    |
|    |____modules # Contains the business logic of all APIs.
|    |
|    |
|    |____pkg # Contains the external dependencies like: database, cache handling
|    |
|    |
|    |____server # Handling the server related code. Starts and apply the server related feature.
|    |
|    |
|    |____utils # Contains the all utilities func which is used in different module. like: read file, bind data, http call etc.
```