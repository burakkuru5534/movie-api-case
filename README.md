# Movie API Service

## Introduction

In this project, we will be building a movie api service.
There will be a register and login api for users to register and login.
There will be a api for users to list movies. Users could reach this api without registration and authentication.
There will be apis for users to create,update, delete and get movies.

### Languages and frameworks

Technologies used in this project:

Golang,
postgresql

Test Environments:

postman,
golang testing library

### Database

Postgresql was used as the database language.

Tables created:

```
table name:sysusr
columns:
id: serial primary key
code: text
upass: text
is_active: boolean
full_name: text
email: text (unique)

table name:movie
columns:
id: serial primary key
name: text (unique)
description: text
type: text
```

## Problem solution

We should be able to create, read, update, and delete movies with jwt authentication.
We should be able to list movies without authentication.
We should prevent movies from creating duplicate names.

### Register


Register request url example:

Method: POST

http://localhost:8080/api/register

request Body Example:

 ```json
{
  "FirstName":"Burak",
  "MiddleName":"",
  "LastName":"Kuru",
  "Email":"brkkr5534@gmail.com",
  "Password":"Test123456"
}
 ```

response example:

for 200:

 ```
User Burak.Kuru created:
 ```

for 400:

```json
{"error": "Bad request"}
```

for 403:
```json
{"error": "User with that email already exists"}
```

for 500:
```json
{"error": "server error"}
```



### Create Movie

Create movie request url example:

Method: POST

Note: this api for authenticated users only.

http://localhost:8080/api/movie

request Body Example:

 ```json
{
  "name": "The Godfather",
  "description": "description",
  "typ": "mafia"
}
 ```

response example:

for 200:

 ```json
{
  "id": 1,
  "name": "The Godfather",
  "description": "description",
  "typ": "mafia"
}
 ```

for 400:

```json
{"error": "Bad request"}
```

for 401:
```
no token found
```

for 403:
```json
{"error": "movie with that name already exists"}
```

for 500:
```json
{"error": "server error"}
```

### Get Movie

Get movie request url example:

Method: GET

Note: this api for authenticated users only.

http://localhost:8080/api/movie?id=1

id: this id should be one of the movie's ids.

request Body:

response example:

for 200:
 ```json
{
  "id": 1,
  "name": "The Godfather",
  "description": "description",
  "typ": "mafia"
}
 ```

for 400:

```json
{"error": "Bad request"}
```

for 401:
```
no token found
```

for 404:
```json
{"error": "movie with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Update movie

Update movie request url example:
Method: PATCH

Note: this api for authenticated users only.

http://localhost:8080/api/movie?id=1

id: this id should be one of the movie's ids.
request Body Example:

 ```json
{
  "name":"UpdatedMovieName",
  "description":"updatedDescription",
  "typ":"updatedTyp"
}
 ```

response example:

for 200:
 ```json
{
  "id": 1,
  "name":"UpdatedMovieName",
  "description":"updatedDescription",
  "typ":"updatedTyp"
}
 ```

for 403:
```json
{"error": "movie with that name already exists"}
```

for 400:

```json
{"error": "Bad request"}
```

for 401:
```
no token found
```

for 404:
```json
{"error": "movie with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Delete movie

Delete movie request url example:
Method: DELETE

Note: this api for authenticated users only.

http://localhost:8080/api/movie?id=1

id: this id should be one of the movie's ids.

response example:

for 200: "Movie deleted."


for 400:

```json
{"error": "Bad request"}
```

for 401:
```
no token found
```

for 404:
```json
{"error": "movie with that id does not exist"}
```

for 500:
```json
{"error": "server error"}
```

### Movie List

Movie List request url example:
Method: GET
Note: This api is for all users.

http://localhost:8080/api/movies

response example:

for 200:
 ```json
[{
  "id": 1,
  "name":"UpdatedMovieName",
  "description":"updatedDescription",
  "typ":"updatedTyp"
}]
 ```
for 400:

```json
{"error": "Bad request"}
```

for 500:
```json
{"error": "server error"}
```

### Test

I used postman and also golang testing libary to test these rest APIs

you can run test by typing:

go test -v

## Conclusion

We have successfully implemented the movie api service.
Used log library to log errors.
Used jwt library to authenticate users.
Used postgresql as the database language.
Used golang testing library to test the apis.
Used postman to test the apis.
Used golang testing library also to test the apis.


