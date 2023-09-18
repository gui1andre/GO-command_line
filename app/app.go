package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação linha de comando"
	app.Usage = "Busca IP e nome do servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "https://www.google.com/",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca endereço de ip na internet",
			Flags:  flags,
			Action: buscarIp,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome do servidor na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIp(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
