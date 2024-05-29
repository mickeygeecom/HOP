#### This is a repo for my major final assignment for my Computer Science education
Online demo can be found here: [https://lootlocker-quiz.fly.dev/](https://lootlocker-quiz.fly.dev/)

# Cool, what is it?
Showcasing creating a quiz system using Golang serving REST API endpoints and static frontend content such as HTML, CSS and VanillaJS.
The system consists of an administration frontend interface and a user frontend interface.
Each quiz has an unique quiz code, which enables participants to join a quiz quickly.
See TODO.md for plans ahead for the project.

## Pre-requisites
- Golang installed (I used version 1.22)
- MariaDB/MySQL database available
- Setup database structure from `"migrations"` folder
- Adjust `.env` file for SQL credentials and HTTP(S) ports

## How do I run it?
- `cd` to `\cmd\server\` directory
- execute `go run main.go`
- Go to https://localhost - if hosted locally

## Fly.io CD
[![Fly Deploy](https://github.com/mickeygeecom/HOP/actions/workflows/deploy.yml/badge.svg)](https://github.com/mickeygeecom/HOP/actions/workflows/deploy.yml)
- Project is setup supporting Fly.io continous deployment, and secrets can be set with f.eg: `flyctl secrets set DB_USER=quiz` as an alternative to the `.env` file
