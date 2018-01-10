# bear dude event API V2
golang version of beardude event api server

**instruction**
* install golang, mongodb
* git clone repository to $GOPATH/src/
* enter "make init" to setup and install dependencies

**development**
```
[terminal 1]
make db
---
[terminal 2]
make dev
```

**production**
```
make build
```

**racer API**
- Content-Type: application/json
* sign up racer
```
POST   localhost:8080/signup       {username, password, nickname}
```
* login & get token
```
POST   localhost:8080/login        {username, password}
```
* racer READ, UPDATE, DELETE
```
POST   localhost:8080/racer/get    {username}
PUT    localhost:8080/racer/update {username, password, nickname}
DELETE localhost:8080/racer/delete {username}
```
