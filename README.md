# Go Hexagonal Neo4j API

This is a sample project that demonstrates how to build a RESTful API with Go, Neo4j, and the Hexagonal architecture. The project uses Gin Gonic as the web framework and includes integration tests.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository: `git clone https://github.com/sembh1998/go-hexagonal-neo4j-api.git`
2. Install the dependencies: `go mod tidy`
3. Set up your environment variables: create a `.env` file at the root of the project and add the following variables:
* NEO4J_URI=bolt://localhost:7687
* NEO4J_USER=neo4j
* NEO4J_PASSWORD=password
4. Run the tests: `go test -v ./...`
5. Start the server: `go run cmd/server/main.go`

You should now be able to access the API at `http://localhost:8080`.

## API Endpoints

The API exposes the following endpoints:

* `GET /products`: returns a list of all products
* `GET /products/:id`: returns the details of a specific product
* `POST /products`: creates a new product
* `PUT /products/:id`: updates an existing product
* `DELETE /products/:id`: deletes an existing product

## Project Structure

The project follows the Hexagonal architecture, with the following structure:
<br>
├── adapters # adapters to external dependencies (e.g. Neo4j)<br>
├── cmd # command-line interfaces (e.g. server)<br>
├── internal # internal packages<br>
│ ├── domain # domain logic (e.g. entities, repositories)<br>
│ ├── handlers # HTTP handlers (e.g. controllers)<br>
│ └── usecases # use cases (e.g. business logic)<br>
├── tests # integration tests<br>
└── .env # environment variables (not included in the repo)<br>

## Technologies Used

* Go 1.20
* Neo4j 5.6.0
* Gin Gonic 1.9.0
* Testify 1.8.1 (for testing)
* Faker 1.2.3 (for testing)

## Contributing

Contributions to the project are welcome. To contribute:

1. Fork the project
2. Create a feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push the branch: `git push origin my-new-feature`
5. Submit a pull request

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
