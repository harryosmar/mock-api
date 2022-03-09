## How to run

- Mac : https://github.com/harryosmar/mock-api/blob/master/build/mock-api-darwin-64
- Windows 64 Bit : https://github.com/harryosmar/mock-api/blob/master/build/mock-api-windows-64
- Windows 32 Bit : https://github.com/harryosmar/mock-api/blob/master/build/mock-api-windows-32

```shell script
git clone git@github.com:harryosmar/mock-api.git

cd mock-api

# for mac
./build/mock-api-darwin-64
```

## Endpoints

### GET "/new"

Request

```shell script
curl --location --request GET 'http://localhost:3001/new'
```

Response
- Status Code : `200` OK
- Content-Type : application/json; charset=UTF-8
- Body :
```json
{
    "text": "welcome to the new page"
}
```

### GET "/"

Request

```shell script
curl --location --request GET 'http://localhost:3001'
```

Response
- Status Code : `301` Moved Permanently
- Content-Type : text/plain; charset=UTF-8

### GET "/nonexisting"

Request

```shell script
curl --location --request GET 'http://localhost:3001/nonexisting'
```

Response
- Status Code : `404` Not Found
- Content-Type : text/plain; charset=UTF-8

### GET "/users"

Request

```shell script
curl --location --request GET 'http://localhost:3001/users'
```

Response
- Status Code : `200` OK
- Content-Type : application/json; charset=UTF-8
- Body :
```json
[
    {
        "name": "john doe",
        "id": 1
    },
    {
        "name": "anna boe",
        "id": 2
    }
]
```

### POST "/users"

Request

```shell script
curl --location --request POST 'http://localhost:3001/users' \
--header 'Content-Type: application/json' \
--data-raw '{"name": "kevin"}'
```

Response
- Status Code : `200` OK
- Content-Type : text/plain; charset=UTF-8
- Body : `kevin has been added to the users list`