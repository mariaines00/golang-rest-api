# Simple golang REST API

- it's a robot friend app 👏 🤖 
 - RESTful service that stores and displays information about robots and their friends.

Using Go, PostgreSQL and Docker.

## How to run it locally
0. have docker and docker-compose installed
1. clone the repo
2. navigate to the directory containing the docker-compose.yml file
3. run docker-compose up -d --build
4. Use [postman](https://www.getpostman.com/) to reach the endpoints


## API Endpoints 

- GET /

- GET /robots => show all robots in the database

- POST /robots => make new bot
  - using json body
  - `{"name": "wall-e", "model": 12345}`
---
- GET /robots/{id} => single robot

- PUT /robots/{id} => update robot data
  - using json body
  - `{"name": "wall-e", "model": 12345}`

- DELETE /robots/{id} => remove robot

---
- GET /robots/{id}/buddies => list friends

- PUT /robots/{id}/buddies => add friend
  - using single form data
  - `key` name `value` wall-e

- DELETE /robots/{id}/buddies => remove friend from list


---
### Author: Maria Inês Serra