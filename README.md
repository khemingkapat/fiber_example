this would be my note of my learning
- **dont forget**, exported component need to be start with capital letter!!
- use `go get -u <module_name>` to get and update module in `go.mod`
- most code come from [mikelopster's tutorial](https://docs.mikelopster.dev/c/goapi-essential/chapter-3/crud). I'll try to extend it for my own variant(for learning)
- to connect dockerized postgres to desktop pgAdmin, use hostname as `127.0.0.1`
- to connect it, use port `5430` since the port `5432` is taken by postgres installed by my `home-manager`(at this time)


# ORM Approach
using ORM is to represent relation(table) as a object in Go. Then the library will help convert it to be relation and deal with db for us.

now we don't need to create table by our own. We will be using **Migration** to create table from object

## Migration
in this example, we would use `db.AutoMigration` from `gorm`, but I suggest(myself) to do more research on **Migration** since there a lot behind the AutoMigration

this will do the job for now but not forever
