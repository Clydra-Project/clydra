## Structur Directory
 - cmd : command line tools for this project
 - service : service for this project, service app or service domain
   
## Getting Started
#### Generate wire
`cd cmd/clydra-server/command && wire`

### Config app with env
`cp .env.example .env`

#### Run server
`go run cmd/clydra-server/main.go`