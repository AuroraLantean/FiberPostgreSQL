# Golang Backend: Fiber + GORM + PostgreSQL + Websocket

TODO: https://github.com/gofiber/fiber
Coding with Verrol
https://www.youtube.com/watch?v=6IU_XdZngv4&list=PL0aDKsruoiW1H8Q1bZPL5SOrXQcibcjrz&index=3

Hitesh Choudhary
https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa


Fiber Framework: https://docs.gofiber.io/
As a rule of thumb, you must only use context values within the handler and must not keep any references. 

GORM, The fantastic ORM library: https://gorm.io/docs/

## Install
Install Go v1.23.4: https://go.dev/doc/install

Install Just command runner: https://github.com/casey/just

Setup you GOROOT and GOPATH
Make a repo folder
```
go mod init backend
go get github.com/gofiber/fiber/v2
go get github.com/charmbracelet/log@latest
just run
```