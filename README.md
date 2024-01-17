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

## Build

```bash
go build -o bin/registryctl main.go
```