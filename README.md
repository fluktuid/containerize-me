# containerize-me

This is some example go code, willing to get containerized.

## The Task
- Get this code running on kubernetes

### First Stage

Get this code in a container while using
- multistage build
- rootless user at runtime
- scratch container

Bonus:
- Implement healthcheck on container-level

### Second stage

- Create a deployment running this container
- Securely expose the container port to the world (tls)

Bonus:
- autoscale the container\
  e.g. based on currently Least Connection or Source IP Hash

## How to build

Building this app is very simple.
If a `go` binary is available, the command is:
```
go build .
```
or, if the executable is to be called 'bar':
```
go build -o bar .
```

## How use
The app will start a webserver running on port `:8080` accepting connections from all hosts.

There is the option configuring a delay by using the env var `DELAY`.
This will set a delay returing the site in milliseconds.


