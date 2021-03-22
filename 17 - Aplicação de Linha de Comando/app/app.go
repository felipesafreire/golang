package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {

	host := ""

}
