#### This is a repo for my major final assignment for my Computer Science education

# Cool, what is it?
Showcasing creating a quiz system using Golang serving REST API endpoints and static frontend content such as HTML, CSS and VanillaJS.
The system consists of an administration frontend interface and a user frontend interface.
Each quiz has an unique quiz code, which enables participants to join a quiz quickly.
See TODO.md for plans ahead for the project.

## Prerequisites
- Golang installed (I used version 1.22)
- MariaDB/MySQL database installed
- Setup database structure from `"migrations"` folder
- Adjust `.env` file for SQL credentials and HTTP(S) ports

## How do I run it?
- `cd` to `\cmd\server\` directory
- execute `go run main.go`
- Go to https://localhost - if hosted locally

## Fly.io CD
[![Fly Deploy](https://github.com/mickeygeecom/HOP/actions/workflows/deploy.yml/badge.svg)](https://github.com/mickeygeecom/HOP/actions/workflows/deploy.yml)
