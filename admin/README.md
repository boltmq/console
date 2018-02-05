# Console Admin

This is a background program - console, used to manage nodes of boltmq cluster.

### Features

* GraphQL API
* Separate frontend and backend
* Built-in file server in console admin
* use Third-party file proxy server, etc nginx.


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
