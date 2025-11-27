package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var tagName string

	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install qBittorrent",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set qBittorrent tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = installBinaryFile(tagName)
				if err != nil {
					return err
				}
				return installService()
			},
		},
		{
			Name:  "uninstall",
			Usage: "Remove config,cache and uninstall qBittorrent",
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = uninstallService()
				if err != nil {
					return err
				}
				return uninstallBinaryFile()
			},
		},
		{
			Name:  "update",
			Usage: "Update qBittorrent",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set qBittorrent tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				err = updateBinaryFile(tagName)
				if err != nil {
					return err
				}
				return updateService()
			},
		},
		{
			Name:  "reload",
			Usage: "Reload service",
			Action: func(ctx context.Context, cmd *cli.Command) (err error) {
				return reloadService()
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n", cmd.Root().Version)
	}

	cmd := &cli.Command{
		Usage:    "qBittorrent quick install tool",
		Version:  "v2.10",
		Commands: cmds,
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
