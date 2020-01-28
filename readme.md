# Geralt


## Development

Best way to set up environment is to use `docker` along with `docker-compose`. 
The following commands will set up the database.

```shell
$ docker-compose up db
$ go run migrations/*.go up
```

To run Geralt server use command below. Remember to set all environment variables.
```shell
$ go run main.go
```

## Client

To run terminal client use the following command.
```shell
$ go run client/main.go http://hostname:port
```
