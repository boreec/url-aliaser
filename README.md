# url-aliaser by Cyprien Borée

This application serves as a user-friendly URL aliasing server. Users can 
provide long URLs, and the server generates shortened aliases that redirect to 
the original URLs. The aliases act as convenient shortcuts for accessing the 
original links.

# Setup

## Building the server

To get started, follow these steps to build the url-aliaser server:

1. Clone the repository and navigate to the cloned directory.
2. Build the application using the `go` command:

```bash
go build ./...
```

Upon successful compilation, you'll find a new executable named `url-aliaser` 
in the current directory.

## Running Unit Tests (Optional)

You can run unit tests to ensure the application functions correctly. Use the 
following command to execute the tests:

```bash
go test ./...
```

# Usage

The URLAliaser server runs by default on  `http://localhost:8080`.

To create a shortened alias, make a POST request to the `/alias`` endpoint with 
a JSON payload. The payload must adhere to the following structure:

```json
{
    "url" : "https://example.com",
    "length" : 10
}
```

- `url`: The original link you want to shorten (required).
- `length`: The desired maximum length of the shortened alias (required).

The server will store the provided URL, generate a new alias, and respond with 
the following JSON payload:

```json
{
    "url" : "http://localhost:8080/xxxx"
}
```
- `url`: The shortened alias, where `xxxx` is a randomly generated string 
matching the requested length.

## example with Curl

In the following command, replace `https://example.com` with the url you wish to alias.

```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com", "length":10}' http://localhost:8080/alias
{"url":"http://localhost:8080/100680ad54"}
```

After this command, you can use `http://localhost:8080/100680ad54` as a 
redirection to `https://example.com`.
