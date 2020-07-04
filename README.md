# Golang CRUD Sample
Docker + Golang(Echo + Gorm) + MySQLでCRUDを行うサンプルプログラム
## Usage
Dockerを使用してGoサーバーとMySQLサーバーを起動する

    $ make build
    $ make up
## CRUD
レスポンスはJSON形式で対象のUser情報を返す
### Create
    $ curl -X POST -H "Content-Type: application/json" -d '{"Name":"Sample太郎", "Age":18, "Sex":1, "Email":"taro@example.com"}' localhost:8080/users
    $ curl -X POST -H "Content-Type: application/json" -d '{"Name":"Sample次郎", "Age":16, "Sex":1, "Email":"joro@example.com"}' localhost:8080/users
    $ curl -X POST -H "Content-Type: application/json" -d '{"Name":"Sample花子", "Age":20, "Sex":2, "Email":"hanako@example.com"}' localhost:8080/users
### Read
- Index
  - `$ curl localhost:8080/users`
- Show
  - `$ curl localhost:8080/users/1`
  - `$ curl localhost:8080/users/2`
  - `$ curl localhost:8080/users/3`
### Update
    $ curl -X PATCH -H "Content-Type: application/json" -d '{"Name":"Sample次郎+1", "Age":16, "Sex":1, "Email":"joro+1@example.com"}' localhost:8080/users/2
### Delete
    $ curl -X DELETE localhost:8080/users/3