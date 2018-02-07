# Console Admin

This is a background program - console, used to manage nodes of boltmq cluster.

### Features

* GraphQL API
* Separate frontend and backend
* Built-in file server in console admin
* use Third-party file proxy server, etc nginx.


### Running Param

* -p listen port
* -f run as foreground
* --pid file, default console.pid.
* --root web root, use file server. not config is disable.
* --prefix web prefix url, use with --root option.
* --index web home url, use with --root option. default index.html
* --debug open debug model & grapql debug api(http://localhost/debug), use with development env.

E.g
```
# enable Built-in file server, set web root and url prefix.
./admin --root=./sources/ --prefix=/sources/
```

or

```
# open grapiql debug, use development debug.
./admin -f --debug --noauth
```

