## go_test is a simple http_server

### Requirement
1. go 1.10.2 or later
2. postgresql 10.4 or later

### Dependencies
1. Golang PostgreSQL driver [Github][1]
2. Gorilla Request Router [Github][2]

install dependencies follow instructions on github pages

### Installation
1. ensure you have `$GOPATH` define in your environment
2. `cd $GOPATH/src && git clone https://github.com/youthy/go_test.git`
3. import sql data by 
`psql -d [database_name]  -f homework.sql [username]` 
replace [database_name], [username] with real database and user name

4. write `config.json` according to your own environment **be aware of that you should use sudo if `listen_port` is less 1024, otherwise error of permission denied will occurred**
5. run `go build` then `./go_test` to start the server

  [1]: https://github.com/go-pg/pg
  [2]: https://github.com/gorilla/mux

