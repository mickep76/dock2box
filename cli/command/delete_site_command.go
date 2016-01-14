package command

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/imc-trading/dock2box/cli/prompt"
	"github.com/imc-trading/dock2box/client"
)

func NewDeleteSiteCommand() cli.Command {
	return cli.Command{
		Name:  "site",
		Usage: "Delete site",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) {
			deleteSiteCommandFunc(c)
		},
	}
}

func deleteSiteCommandFunc(c *cli.Context) {
	var site string
	if len(c.Args()) == 0 {
		log.Fatal("You need to specify a site")
	} else {
		site = c.Args()[0]
	}

	clnt := client.New(c.GlobalString("server"))
	if c.GlobalBool("debug") {
		clnt.SetDebug()
	}

	if !prompt.Bool("Are you sure you wan't to remove "+site, true) {
		os.Exit(1)
	}

	h, err := clnt.Site.Delete(site)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v\n", string(h.JSON()))
}
