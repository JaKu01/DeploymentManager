# Deployment Manager

## Overview
This project aims at simplifying the deployment 
of docker-compose stacks.
By specifying the URL of a repository containing a docker-compose file,
the deployment manager will clone the repository, build the images and start the services.

## Usage
Build the deployment manager with the following command:
```bash
go build
```

After that, you can start the deployment manager with the following command:
```bash
./DeploymentManager
```
The deployment manager exposes a REST API to manage projects at port 8080.
## User Interface
The UI is available at `:8080/`

## API Documentation
The API documentation is available at `:8080/swagger`
