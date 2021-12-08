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

