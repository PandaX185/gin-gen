# Gin Gen
## Overview
This is a simple go package to help build backend systems faster with less code. 

It provides generators for jwt tokens, middleware, rest controllers and more.

It uses gin for http request handling and gorm for db management

## Installation
To start using this package, you need to install it

`go install github.com/PandaX185/gin-gen@latest`

Run this command to ensure the latest version of the package anytime

## Commands

- `gin-gen`: show help menu
- `gin-gen jwt <optional-name>`: generates a new file with token generation and middleware (file name is optional)
- `gin-gen model <name>`: creates a new gorm model file with the specified name
- `gin-gen repo <name>`: creates a new repository file with the specified name and creates the model if not created
- `gin-gen service <name>`: creates a new service file with the specified name and creates the model and repo if not created
- `gin-gen controller <name>`: creates a new controller file with the specified name and creates the model, repo and service if not created

Notice that all files have some CRUD samples
