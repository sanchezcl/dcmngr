# dcmngr
dcmgr it's a wrapper command for docker-compose with most common day by
day actions, to optimize time writing less commands to build, up, down, stop, 
get a terminal inside a container and others functions.

### Dependencies
- A proper installation of docker and docker-compose
- A project with a docker-compose.yml

### Installation
- Download the binaries from [releases section](https://github.com/sanchezcl/dcmngr/releases)
- Unpack from the compressed file
- Copy to /bin directory or a directory in your path environment variable
- Give execution permissions if is necessary
- Go to your docker-compose project directory
- Execute:
```bash
dcmngr genyml //this will make a .dcmngr.yml with your services 
              //in the docker-compose.yml file
```
- Edit .dcmngr.yml as you need

### Config file (.dcmngr.yml)
```yaml
sh_default_service: ""    //the default container where to run a terminal
sh_always_admin: true     //get the terminal always as an admin
build_default_containers: //container to build
  - 
up_default_containers: //containers to start
  - 
watch_configs:
  service: "" //service to go in in the compilation
  command: "" //command to execute
  args: []    //args for the compilation commnad
```

This is an example for golang (go) project with [cosmtrek/air](https://github.com/cosmtrek/air):
```yaml
sh_default_service: "golang"
sh_always_admin: true
build_default_containers:
  - mongo
  - mongo-webui
  - golang
  - redis
  - redis-webui
up_default_containers:
  - mongo
  - golang
  - redis
watch_configs:
  service: golang
  command: air
  args:
    - "-c"
    - "/etc/air.conf"
```

### Available Commands

```bash
build       Build or rebuild services
down        Stop and remove containers, networks, images, and volumes
genyml      Generate .dcmngr.yml file
help        Help about any command
logs        View output from containers
ps          List containers
sh          Get a shell inside a container.
stop        Stop running containers without removing them.
up          Create and start containers
watch       Runs live reload/build
```