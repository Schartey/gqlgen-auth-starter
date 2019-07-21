# gqlgen Starter with Authentication and Authorization using Keycloak (WIP)

This project gives an example of how to use [99designs/gqlgen][1] to create a GraphQL - Webservice with implented authentication and authorization using [Keycloak][2].
It also integrates concepts used in [OscarYuen/go-graphql-starter][3], which was designed for use with [graph-gophers/graphql-go][4]. Unfortunately that library missed 
a few features that were needed which is why gqlgen is now being used. 

The purpose of this repository is to help others get started and receive valuable feedback to improve this starter pack. Pull Requests are welcome as well.

## Roadmap

- [x] Docker Support
- [x] Docker Compose File
- [x] Use reflex for hot-reload
- [x] Integrate gqlgen
- [x] Split resolvers
- [x] Schema generation with [mattdamon108/gqlmerge][7]
- [ ] Logging
- [ ] Proper context usage
- [ ] Cursor based pagination (Relay Cursor Connections Specification)
- [ ] Add Authentication using Keycloak
- [ ] Add Authorization using Keycloak
- [ ] Add Registration
- [ ] Use GORM with Postgres to save additional data
- [ ] Use Redis to save session data
- [ ] Caching GraphQL Requests
- [ ] Caching Keycloak Data
- [ ] Add unit tests
- [ ] Show example usage of subscription

## Structure

```
gqlgen-auth-starter/
├── docker/
│   ├── docker-entrypoint.sh
│   └── README.md
├── keycloak/
│   └── realm-export.json
├── gqlgen/
│   ├── resolvers/
│   │   ├── rootMutationResolver.go
│   │   ├── rootQueryResolver.go
│   │   └── rootResolver.go
│   ├── generated.go
│   ├── gqlgen.yml
│   ├── models_gen.go
│   └── README.md
├── schema/
│   ├── directives/
│   ├── input/
│   │   └── user.input.graphql
│   ├── mutation/
│   │   └── user.mutation.graphql
│   ├── query/
│   │   └── user.query.graphql
│   ├── subscription/
│   ├── types/
│   │   ├── time.type.graphql
│   │   └── user.type.graphql
│   └── schema.graphql
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE.md
├── README.md
├── reflex.conf
└── server.go
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
[reflex][5] is being used for hot-reload. It tracks changes on all .go and .graphql files except those generated by gqlgen and gqlmerge. 

### docker compose

Running the project with docker-compose is faster and easier. Just run the following command to start the project.
This will also build the docker image locally and in future start up needed databases. The project runs on port 3000 with this config, since
keycloak already uses 8080.

```
docker-compose up
```

To start keycloak without the project use the docker-compose.keycloak.yml. 

```
docker-compose -f docker-compose.keycloak.yml up
```

To start the project with keycloak use both compose files. On Windows hot-reload is not working as intended (see Known Issues),
therefore it's recommended to start keycloak independent as stated above, else you will have to restart both the keycloak and
project container continuously.

```
docker-compose -f docker-compose.yml -f docker-compose.keycloak.yml up
```

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

### Keycloak

This project uses Keycloak as Authentication Provider. A demo realm with preconfigured client and client roles is imported
with the docker-compose file. You can adjust the configuration file in the keycloak folder or the web-interface. 

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
      resolvers. Resolvers are generally not difficult to write, so I would suggest this option. I also recommend looking at [git-bug][6]
      for more examples, since gqlgen-auth-starter works the same way.
      
      
[1]: https://github.com/99designs/gqlgen
[2]: https://www.keycloak.org/
[3]: https://github.com/OscarYuen/go-graphql-starter
[4]: https://github.com/graph-gophers/graphql-go
[5]: https://github.com/cespare/reflex
[6]: https://github.com/MichaelMure/git-bug
[7]: https://github.com/mattdamon108/gqlmerge