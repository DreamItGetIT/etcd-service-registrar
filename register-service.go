package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

const (
	dockerIP = "172.17.42.1"
)

func commandExecuter(c *cli.Context) {
	if !c.IsSet("container") {
		fmt.Println("--container argument is required")
		return
	}

	if c.IsSet("delete-all") {
		if err := deregister(c.GlobalString("container")); err != nil {
			fmt.Fprintf(os.Stderr, "deregistration for container %s, failed: %s\n", c.GlobalString("container"), err)
			os.Exit(1)
		}

		fmt.Printf("%s container deregistered\n", c.GlobalString("container"))
		os.Exit(0)
	}

	if !c.IsSet("port") {
		fmt.Println("--port argument is required")
		return
	}

	var mappedPort string
	if c.IsSet("mapped-port") {
		mappedPort = c.GlobalString("mapped-port")
	} else {
		mappedPort = c.GlobalString("port")
	}

	if err := register(c.GlobalString("container"), c.GlobalString("ip"), c.GlobalString("port"), mappedPort, uint64(c.GlobalInt("ttl"))); err != nil {
		fmt.Fprintf(os.Stderr, "registration for container %s, failed: %s\n", c.GlobalString("container"), err)
		os.Exit(1)
	}

	fmt.Printf("host's ip and port for container %s registered\n", c.GlobalString("container"))
	os.Exit(0)
}

func register(container string, ip string, port string, mappedPort string, ttl uint64) error {
	etcdClient := etcd.NewClient([]string{fmt.Sprintf("http://%s:4001", dockerIP)})

	// Create directory if it already exist it report an error, but it is ignore
	etcdClient.SetDir(fmt.Sprint("containers/", container), ttl)

	if _, err := etcdClient.Set(fmt.Sprintf("containers/%s/ports/%s/host/", container, port), ip, ttl); err != nil {
		return err
	}

	if _, err := etcdClient.Set(fmt.Sprintf("containers/%s/ports/%s/port/", container, port), mappedPort, ttl); err != nil {
		return err
	}

	return nil
}

func deregister(container string) error {
	etcdClient := etcd.NewClient([]string{fmt.Sprintf("http://%s:4001", dockerIP)})

	if _, err := etcdClient.Delete(fmt.Sprint("containers/", container), true); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "register"
	app.Usage = "Register and deregister in local Etcd container the host and port of the specified service name"
	app.Action = commandExecuter
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "container, c", Usage: "The container name or id"},
		cli.StringFlag{Name: "ip, i", Value: "127.0.0.1", Usage: "The host's ip name or ip where the service is running"},
		cli.StringFlag{Name: "port, p", Usage: "The original port of the service"},
		cli.StringFlag{Name: "mapped-port, mp", Usage: "The real port which the service is listening; by default port"},
		cli.IntFlag{Name: "ttl", Value: 0, Usage: "The time to live for the registration"},
		cli.BoolFlag{Name: "delete-all, da", Usage: "Remove all the entries for this container on etcd"},
	}

	app.Run(os.Args)
}
