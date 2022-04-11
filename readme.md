# Golang Web Server With Fiber

<center>
<img align="center" width="150" height="150" alt="Logo" src="https://go.dev/images/gophers/motorcycle.svg">
</center>

## fresh Package

fresh tools run as like as pm2 

## Docker Compose Rebuild run

docker-compose up --build --force-recreate

## ORM Config

Here i use gorm and mysql with docker compose. Here create a docker-compose yml, Here the mysql container name is the database URL and PORT
Now create a user and password. Give your permission and the same username and password give to the gorm config. As config/database.go configer.

Now Run the docker-compose with build and recreate. Hope its working well.