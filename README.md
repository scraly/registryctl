# registryctl

The goal of this CLI is to manage your various OCI/Docker images registries (Docker Hub, OVHcloud Managed Private Registry, Artifactory, Google Container Registry, Harbor, Docker registries…).

Warning: this CLI have been created for a Programmez! magazine article so for the moment this CLI is not complete and production ready and it's only getting OVHcloud Managed Private Registries ;-).

## Get OVHcloud credentials/tokens

Create a `.env` file and define into it the environment variables:

```
OVH_ENDPOINT="ovh-eu"
OVH_APPLICATION_KEY="xxx"
OVH_APPLICATION_SECRET="yyy"
OVH_CONSUMER_KEY="zzz"
OVH_CLOUD_PROJECT_SERVICE="aaa"
```

More infos: https://docs.ovh.com/gb/en/customer/first-steps-with-ovh-api/

## Run

```bash
go run main.go ovhMpr get
```

Example of result:

```
NAME                    STATUS  REGION  URL                                                     VERSION CREATED AT                 
my-registry-01d0aa3     READY   GRA     https://xxxxxxxx.c1.gra9.container-registry.ovh.net     2.6.4   2023-12-13T15:10:26.72334Z 
my-registry-8bbe01f     READY   GRA     https://yyyyyy.c1.gra9.container-registry.ovh.net       2.6.4   2023-12-19T12:40:44.442604Z
my-new-registry         READY   GRA     https://zzzzzz.c1.gra9.container-registry.ovh.net       2.6.4   2023-11-22T14:11:53.250714Z
my-registry-b425251     READY   GRA     https://aaaaaaa.c1.gra9.container-registry.ovh.net      2.6.4   2024-01-05T13:02:59.536574Z
my-registry             READY   GRA     https://bbbbbb.gra7.container-registry.ovh.net          2.6.4   2022-04-12T09:36:36.559078Z
```

## Build

```bash
go build -o bin/registryctl main.go
```