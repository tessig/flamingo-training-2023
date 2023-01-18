# Getting started with Flamingo 2023

A solution to the coding exercises as seen in the workshop

## Keycloak

To run this project, you will need a local keycloak:

```shell
cd keycloak
docker run --rm -ti -p 8080:8080 -v $(pwd)/import:/opt/keycloak/data/import -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:20.0.2 start-dev --import-realm
```

## Other local services

### Jaeger

```shell
docker run --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.17.1
```
