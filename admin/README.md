# Console Admin

This is a background program - console, used to manage nodes of boltmq cluster.

### Features

* GraphQL API
* Separate front and rear ends
* Built-in file server in console admin
* file server use Third-party server, etc nginx.


### Get it

**Build it from source code**

Get source code from Github:
```Go
git clone https://github.com/boltmq/console.git
```


### Running Param

* -p listen port
* --root web root, use file server. default ./sources. 
* --perfix web perfix url, use with --root option.
* --index web home url, use with --root option. default index.html
* --debug open debug model & grapql debug api(http://localhost/debug), use with development env.

E.g
```
./admin --debug --root=./sources/ --perfix=/sources/
```
