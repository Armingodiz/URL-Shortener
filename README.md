# URl Shortener 
 simple url-shortener written in Go .
 
 Used DataBase : [redis](https://redis.io/)
 
## Features
1. signing up & logging in
2. shortenning inputted url 
3. showing shortenned urls for each user 

## Dependencies
name     | repo
------------- | -------------
  gorilla/mux | https://github.com/gorilla/mux
  go redis    | https://github.com/go-redis/redis
  go.uuid     | https://github.com/satori/go.uuid 
 
 Also :
 * [Docker](https://www.docker.com/)
## Usage

use `docker run --name redis-usdb -p 8282:6379 -d redis` to connect redis to port 8282

then just build and run **main.go** file
     default port for service is `:8080 `
