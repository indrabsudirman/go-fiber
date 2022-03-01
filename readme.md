
#### Learn Golang as Backend using Gofiber and GORM


# API Specs

##### RestAPI : https://learn-gofiber-gorm.herokuapp.com

&nbsp;
&nbsp;

## Sign Up User


- Method : POST
- Endpoint : `/user`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "id": "bigint, auto_increment ",
    "name": "string",
    "email": "string",
    "address": "string",
    "phone": "string",
    "password": "string, bcrypt hashing"
}
```
- Response :
```json
{
    {
    "data": {
        "id": 1,
        "name": "Indra Sudirman",
        "email": "indrasudirman@gmail.com",
        "address": "Tangerang Selatan",
        "phone": "089636002345",
        "created_at": "2022-03-01T07:03:16.600610182Z",
        "updated_at": "2022-03-01T07:03:16.600610182Z"
    },
    "message": "success"
}
}
```

## Login


- Method : POST
- Endpoint : `/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    
    "email": "string",
    "password": "string, bcrypt hashing"
}
```
- Response :
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiVGFuZ2VyYW5nIFNlbGF0YW4iLCJlbWFpbCI6ImluZHJhc3VkaXJtYW5AZ21haWwuY29tIiwiZXhwIjoxNjQ2MTIyMjE3LCJuYW1lIjoiSW5kcmEgU3VkaXJtYW4iLCJwaG9uZSI6IjA4OTYzNjAwMjM0NSIsInJvbGUiOiJhZG1pbiJ9.lR2y2aHHh9bgXPttdjPSSkEMbltr0BzSbCDv1WHniWA"
}
```


## Get All User


- Method : GET
- Endpoint : `/user`
- Header :
    - Content-Type: application/json
    - Accept: application/json 
    - Key: JWT 

- Response :
```json
[
    {
        "id": 1,
        "name": "Indra Sudirman",
        "email": "indrasudirman@gmail.com",
        "address": "Tangerang Selatan",
        "phone": "089636002345",
        "created_at": "2022-03-01T07:03:16.60061Z",
        "updated_at": "2022-03-01T07:03:16.60061Z"
    }
]
```
