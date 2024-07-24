<div align=center>create a json crud api in go,gin-gonic/gin</div>
<div align=center>coding with robby</div>

#### This project repository gives the idea of how we can use gin-gonic/gin package

#### We will understand the package by creating a _CRUD BACKEND API PROJECT_

#### following list of articles and youtube video lectures were consumed to create this repository

1. [Creating a JSON CRUD API in Go (Gin/GORM)](https://youtu.be/lf_kiH_NPvM?list=PL9KOhO7TtG6YazaOAKlU_elPlRDQ2ht3n) by _Coding with Robby_

> a heartiest thanks to all the creators of the video lectures and articles for the content ♥

##### the application is running on port 3000

#### we are using the following package for this project

1. CompileDaemon --> go get github.com/githubnemo/CompileDaemon
2. goDotEnv --> go get github.com/joho/godotenv
3. gin-gonic/gin --> go get -u github.com/gin-gonic/gin
4. gorm --> go get -u gorm.io/gorm
5. mysql-driver --> go get -u gorm.io/driver/mysql

#### what does CompileDaemon do ? and how to run it

CompileDaemon provides hot reloading on change
syntax to run it is

> CompileDaemon -command="./git-mod-file-name"
> ex: CompileDaemon -command="./go-gin-crud-api-by-robby"

#### creating a .env file and using goDotEnv package

PORT=
USER=
PASSWORD=
ADDRESS=
DB_NAME=
loading env variables in project

```go
err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}
```
