# Deployment Manager

## Overview
This project aims at simplifying the deployment 
of docker-compose stacks.
By specifying the URL of a repository containing a docker-compose file,
the deployment manager will clone the repository, build the images and start the services.

## Usage
The deployment manager provides a REST API to manage projects.
It is planned to support https and ssh URLs for the repository, 
as the program will use the users the git command line tool to clone the repository.