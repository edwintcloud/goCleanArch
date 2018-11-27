# goCleanArch
An API in Go following Clean Architecture principles as outlined [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and further simplified in [this](https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f) Medium article.

## Structure
### entities
Also know as models, here lives our data structures and functions that define how data is stored in our database(s).
### repositories
Here is where our database handlers go. Each database handler should be defined seperately with it's own interface functions. This layer acts for CRUD operations to specified repository. Repositories are not limited to databases, any data inputs can be included such as pulling data from microservices and sanitizing the data. No business logic here, that will be handled in usecases layer.
### usecases
Here is where we process our business logic. This layer will decide which repository to use and handle sanitized data flowing from the controllers/delivery layer to the specified repository and vice-versa. This layer can also be called services.
### controllers
Also know as the delivery layer, this layer will act as the presenter. In the case of an API, then means sending back JSON data and receiving JSON, sanitizing that JSON and sending to the proper usecase/service. 

## Setup
1. Clone this repo in your $GOPATH
2. Use `govendor sync` to fetch project dependencies
3. To run server use `go run server.go`