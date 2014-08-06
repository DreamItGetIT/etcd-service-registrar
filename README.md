# etcd service registrar

An opinionated command line tool, implemented in [Go](http://golang.org/), to register a service host a port in a local [etcd](http://coreos.com/using-coreos/etcd/) instance which runs into a [Docker](https://www.docker.com/) container.  


## Description

An opinionated command line tool, implemented in [Go](http://golang.org/), to register a service host a port in a __local [etcd](http://coreos.com/using-coreos/etcd/) instance__ which runs into a __[Docker](https://www.docker.com/) container__.  

The __local [etcd](http://coreos.com/using-coreos/etcd/) instance__ must run into a docker container, which must expose the [etcd](http://coreos.com/using-coreos/etcd/) client default port 4001 in all the network interfaces, IP `0.0.0.0`, therefore it is accessible under the IP `172.17.42.1` (default docker daemon IP).

The tool is bound to the data structure used by [etcd container presence](https://github.com/DreamItGetIT/etcd-container-presence) used to register the hosts and ports exposed by services, which may be [Docker](https://www.docker.com/) containers but it is not a must, as this tool has been backed to allows to publish the services running in your local machine to give the ability to other services which depend of them to run and access them meanwhile you are working (fixing/refactoring/implementing) on them.


## How to use

The binary has the classical brief help about the arguments that it must and may been provided, showed when `-h`, or it is extended version `--help`, argument is used.

You can download it from [here](https://github.com/DreamItGetIT/etcd-service-registrar/releases/download/v0.0.0/register-service.zip).

However you can also take a look them there:
```
   --container, -c      The container name or id [REQUIRE]
   --ip, -i '127.0.0.1' The ip or name of the host where the service is running
   --port, -p           The original port of the service [REQUIRED]
   --mapped-port, --mp  The real port which the service is listening; by default port
   --ttl '0'            The time to live for the registration
   --delete-all, --da   Remove all the entries for this container on etcd; 
                        when used only --container is required and used
   --help, -h           show help
   --version, -v        print the version
```


## Compile me

To compile it, you must have [gom](https://github.com/mattn/gom) installed and working; then you are ready to compile it, just executing `make`.

## Why we implemented it

We implemented it to have a simple command line tool that allows us to register services in [etcd](http://coreos.com/using-coreos/etcd/) obeying the same data structures use by [etcd container presence](https://github.com/DreamItGetIT/etcd-container-presence) without the constraint to run the services into [Docker](https://www.docker.com/) containers but in a host machine (the most of the times `localhost`), which it is a requirement during the development cycle, whenever we must run different services, which some depend of others.

Therefore the tool allows us to remove the hassle to create different configurations between development and testing/production environments, which it is appreciated because we remove the differences between them and we are less prone to introduce issues which would be detected, hopefully, in CI testing than production.

## License

Just MIT, Copyright (c) 2014 DreamItGetIT, read LICENSE file for more information.

