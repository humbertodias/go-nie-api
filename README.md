# API using GO

Simple API that crawles the website sede.administracionespublicas.gob.es and publish the offices and procedures that each one can deal with.

# Clone

    mkdir -p ~/go/src/github.com/humbertodias && cd $_
    git clone https://github.com/humbertodias/go-rest-api
    cd go-rest-api

# Run

    make run

# Api

    http://localhost:8080/oficinas

![](doc/oficinas.png)


    http://localhost:8080/oficinas

![](doc/tramites.png)


# Ref

[building-a-web-http-server-with-go](https://itnext.io/building-a-web-http-server-with-go-6554029b4079)