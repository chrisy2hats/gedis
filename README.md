# gedis  
A redis-like program supporting a subset of redis commands and features with roughly the same syntax.  
Just a project for learning Golang, Redis and Sockets. Not for serious/production use  

# Building
```github
export GO111MODULE=off
export GOPATH=$PWD
go run gedis.go
```
# To use 
There is a Python client in-tree but a socket from any language should work      

```github
python3 client.py
> set a 5
> get a
5
>
```

# Supported Commands
The commands and arguments are much the same as redis-cli   

| Key    | Usage        | Description                                                      |
|--------|--------------|------------------------------------------------------------------|
| SET    | SET foobar 5 | Set a KEY to a provided value                                    |
| GET    | GET foobar   | Fetch the value of a key                                         |
| KEYS   | KEYS foo*    | Fetch the keys of all those matching the regex                   |
| LPUSH  | LPUSH foo 5  | Initialise or append to a list                                   |
| LINDEX | LINDEX foo 0 | Fetch the value at a given index in a list                       |
| LRANGE | LRANGE 0 5   | Fetch the values between the indices provided values in the list |


# To run unit tests run  
```github
go test -v ./...
```
