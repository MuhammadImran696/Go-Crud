# Go-Crud

## How to run this service locally?
The first step is to clone the microservices. It can be done through the following command:
`git@github.com:MuhammadImran696/Go-Crud.git`

### Databases
you have to craete databases for both microservices which are named below <br/>
1) user
2) company

After this step, We can run the microservices by the following steps: <br/>
Step 1: Starting the gateway Server <br/>

run these commands <br/>
`cd ApiGateway`
`go run cmd/main.go` <br/>

Step 2: Starting the Auth service <br/>

run these commands <br/>
`cd authService`
`go run cmd/main.go` <br/>

Step 3: Starting the Company service <br/>

run these commands <br/>
`cd companyService`
`go run cmd/main.go` <br/>

#### EndPoints
POST    http://localhost:3000/auth/register <br/>
POST    http://localhost:3000/auth/login <br/>
POST    http://localhost:3000/product <br/>
DELETE    http://localhost:3000/product/id <br/>
GET    http://localhost:3000/product/id <br/>
PATCH    http://localhost:3000/product/id <br/>



##### Example
1- Login <br/>
### Request
curl --location --request POST 'http://localhost:3000/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"imran@gmail.com",
    "password":"imran"
}'
### Response
`{
    "status": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg0MTAxNDEsImlzcyI6ImF1dGhzZXJ2aWNlIiwiSWQiOjEsIkVtYWlsIjoic2FtZWVyQGdtYWlsLmNvbSJ9.nsl2f1Td1OHYlerDbVc8goHEKXZ-p4Z1XB6z2W3TG2Q"
}`

2- Create Company <br/>
### Request
curl --location --request DELETE 'http://localhost:3000/product' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgzOTg1MzksImlzcyI6ImF1dGhzZXJ2aWNlIiwiSWQiOjEsIkVtYWlsIjoic2FtZWVyQGdtYWlsLmNvbSJ9.JfcRU6W8awD0iZUyL_p0HNOqEL53rPVuk2rmE6mMFcY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"abc",
    "description":"abc",
    "employees":67,
    "registered":false,
    "type":"private"
}'
### Response
`{
    "status": 201,
    "id": "f691a4e6-6ada-427c-98e8-83a1292634d3"
}`
