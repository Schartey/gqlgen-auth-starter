# gqlgen Starter with Authentication and Authorization using Keycloak

This project gives an example of how to use [99designs/gqlgen][1] to create a GraphQL - Webservice with implented authentication and authorization using [Keycloak][2].
I have also integrated concepts used in [OscarYuen/go-graphql-starter][3], which was designed for use with [graph-gophers/graphql-go][4]. Unfortunately that library missed 
a few features that I needed which is why I switched to gqlgen.

The purpose of this repository is to help others get started and receive valuable feedback to improve this starter pack. Pull Requests are welcome as well.

## Roadmap

- [x] Docker Support
- [ ] Docker Compose File
- [x] Use reflex for hot-reload
- [x] Integrate gqlgen
- [ ] Use go-bindata to merge schema files
- [ ] Add Authentication using Keycloak
- [ ] Add Authorization using Keycloak
- [ ] Add Registration
- [ ] Use GORM with Postgres to save additional data
- [ ] Use Redis to save session data
- [ ] Add unit tests
- [ ] Show example usage of subscription

## Structure

```
gqlgen-auth-starter/
├── docker/
│   ├── docker-entrypoint.sh
│   └── README.md
├── gqlgen/
│   ├── ... generated gqlgen-files
│   └── README.md
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE.md
├── README.md
└── reflex.conf
```

## Requirements

- Golang
- Docker

## Usage

### docker run

You can either use the prepared gqlgen - image I made or build it on your own. To edit gqlgen-files
use docker volumes and bind /app. 
```
docker run -v PROJECT_PATH:/app -p 8080:8080 schartey/gqlgen
# or
docker build . --tag=YOUR-TAG
docker run -v PROJECT_PATH:/app -p 8080:8080 YOUR-TAG
```

You can also bind your go cache for faster startup time.
```
docker run -v PROJECT_PATH:/app -v $GOPATH/pkg/mod/cache:/go/pkg/mod/cache -p 8080:8080 YOUR-TAG
```

The image will generate the gql-files automatically at first start.
Afterwards it will update generated files according to your changes and run the server.
I use [reflex][5] for hot-reload. It tracks changes on all .go files except those generated by gqlgen. 

### Development

To start developing just follow the guide on the gqlgen github page. There shouldn't be any difference.
Here is a short starter:

- Create your own schema.graphql in gqlgen folder
- Write your resolvers
- Adjust your gqlgen.yml
- Write your own packages

Each of these steps should cause reflex to hot-reload and by that regenerating needed gqlgen data.

### Basic Image

As mentioned before there is a prepared image which can be used to develop with gqlgen. This image will be tagged in
git with gqlgen-starter. Further development of the actual gqlgen-auth-starter project will continiously be merged to master.

## Known Issues

- The current version of gqlgen supports directives, but is not yet bug free. If you encounter the error:
```
unable to parse config: yaml: unmarshal errors:
    line 18: key "deprecated" already set in map
    line 20: key "include" already set in map
    line 22: key "skip" already set in map
```
you have to remove (comment) the directives settings in gqlgen.yaml.

- Reflex does not support Windows, therefor hot-reload will not work on this platform.

- As mentioned in this [Thread](https://github.com/99designs/gqlgen/blob/master/docs/content/getting-started.md#write-the-resolvers) 
  gqlgen is currently not able to update resolver.go. There are two viable solutions to this problem:
    - Copy the contents of resolver.go, delete it, run gqlgen and paste into the newly generated resolver.go. (Becomes tedious and hard to manage)
    - Set gqlgen to not generate resolver.go in the config file and just create the resolvers for yourself. This way you can also manage multiple files with multiple
      resolvers. Resolvers are generally not difficult to write, so I would suggest this option. I also recommend looking at [git-bug](https://github.com/MichaelMure/git-bug)
      for more examples, since gqlgen-auth-starter works the same way.
      
      
[1]: https://github.com/99designs/gqlgen
[2]: https://www.keycloak.org/
[3]: https://github.com/OscarYuen/go-graphql-starter
[4]: https://github.com/graph-gophers/graphql-go
[5]: https://github.com/cespare/reflex