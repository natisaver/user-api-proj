# Go API Server and CURL Commands

## Run the Go API Server

Open a terminal and execute the following commands to run the Go API server:

```bash
cd /Users/pncy1926/Desktop/goapi
go run .
```

## Making Curl Requests
- After the API server is running, you can interact with it using curl commands. Open another terminal window and use the following commands:

## Create a New User
```bash
curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4", "name": "testname", "age": 55}'
```

## Retrieve All Users
```bash
curl http://localhost:8080/users
```

## Retrieve a Specific User by ID
```bash
curl http://localhost:8080/users/2
```

## Update a User by ID
```bash
curl -X PUT \
    http://localhost:8080/users/1 \
    --include \
    --header "Content-Type: application/json" \
    --data '{"id": "1", "name": "Updated Name", "age": 35}'
```