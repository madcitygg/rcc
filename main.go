package main

import (
	"errors"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/codegangsta/cli"
	"github.com/madcitygg/rcon"
	"github.com/mitchellh/colorstring"
	"io"
	"os"
	"strconv"
	"strings"
)

// Server addresses

const (
	defaultPort = 27015
)

type Address struct {
	Host string
	Port int
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func ParseAddress(address string) (a *Address, err error) {
	var host string
	var port int

	parts := strings.Split(address, ":")
	if len(parts) > 2 {
		err = errors.New("address contains multiple colons")
	}

	if len(parts) == 1 {
		host = parts[0]
		port = defaultPort
	} else if len(parts) == 2 {
		host = parts[0]
		port, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}
	}

	a = &Address{host, port}
	return
}

// Main loop

func main() {
	// global vars
	password := ""
	authenticated := false

	// set up cli app
	app := cli.NewApp()
	app.Name = "rcc"
	app.Usage = "control a Source RCON server"
	app.UsageText = "rcc [global options] server"
	app.Version = "0.0.1"

	// add global command line flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "rcon password",
			Destination: &password,
		},
	}

	// default action
	app.Action = func(ctx *cli.Context) {
		if len(ctx.Args()) == 0 {
			// the user did not provide a server to connect to
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}

		// parse server address
		address, err := ParseAddress(ctx.Args()[0])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		// dial the server
		colorstring.Printf("Connecting to %s\n", address)
		rc, err := rcon.Dial(address.String())
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		defer rc.Close()
		colorstring.Printf("[green]Connection successful![reset]\n")

		// create readline instance
		rl, err := readline.NewEx(&readline.Config{
			Prompt:          "> ",
			InterruptPrompt: "\nPress Ctrl+D to exit",
			EOFPrompt:       "exit",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		passwordConfig := rl.GenPasswordConfig()
		passwordConfig.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
			rl.SetPrompt("Enter password: ")
			rl.Refresh()
			return nil, 0, false
		})

		// start inner console
		console := cli.NewApp()
		console.Commands = []cli.Command{
			{
				Name:  "login",
				Usage: "Log in to rcon server",
				Action: func(c *cli.Context) {
					askForPassword := func(destination *string) {
						for len(*destination) == 0 {
							// read a password from command line
							pw, err := rl.ReadPasswordWithConfig(passwordConfig)
							if err == io.EOF {
								os.Exit(1)
							}
							if err != nil {
								colorstring.Printf("Press Ctrl+D to exit\n")
							}
							// clean it up and check if it makes sense
							*destination = strings.TrimSpace(string(pw))
							if len(*destination) == 0 {
								colorstring.Printf("[red]Please enter a valid password[reset]")
							}
						}
					}

					// ask for a password until we authenticate successfully
					for !authenticated {
						askForPassword(&password)
						err := rc.Authenticate(password)
						if err != nil {
							colorstring.Printf("[red]Authentication failed. Incorrect password?[reset]")
							password = ""
							continue
						}

						authenticated = true
					}
				},
			},
			{
				Name: "clear",
				Action: func(c *cli.Context) {
					colorstring.Printf("[red]`clear` has not been implemented yet.[reset]\n")
				},
			},
			{
				Name:    "exit",
				Aliases: []string{"quit", "q"},
				Action: func(c *cli.Context) {
					os.Exit(1)
				},
			},
		}
		console.Action = func(c *cli.Context) {
			command := strings.Join(c.Args(), " ")
			response, err := rc.Execute(command)
			if err != nil {
				colorstring.Printf("[red]Error: %s[reset]\n", err)
			}

			fmt.Println(response.Body)
		}

		// perform login to authenticate with rcon server
		console.Run(strings.Fields("cmd login"))

		// set a nice prompt with server name and port
		prompt := "[white]%s[reset]:[blue]%d[reset]> "
		rl.SetPrompt(colorstring.Color(fmt.Sprintf(prompt, address.Host, address.Port)))

		// start readline loop
		for {
			line, err := rl.Readline()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}

			console.Run([]string{"cmd", line})
		}
	}

	app.Run(os.Args)
}
