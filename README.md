# Micro-services GoLang Example

## Folder structure

- **api**: Stores the versions of the APIs swagger files and also the proto and pb files for the gRPC protobuf interface.
- **cmd**: This will contain the entry point (main.go) files for all the services and also any other container images if any
- **docs**: This will contain the documentation for the project
- **config**: All the sample files or any specific configuration files should be stored here
- **deploy**: This directory will contain the deployment files used to deploy the application
- **internal**: This package is the conventional internal package identified by the Go compiler. It contains all the packages which need to be private and imported by its child directories and immediate parent directory. All the packages from this directory are common across the project
- **pkg**: This directory will have the complete executing code of all the services in separate packages.
- **tests**: It will have all the integration and E2E tests
- **vendor**: This directory stores all the third-party dependencies locally so that the version doesn’t mismatch later

## Articles

https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golang