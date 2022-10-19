package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = argoCmd
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang",
			EnvVar: "GITHUB_REPOSITORY",
			Usage:  "language for the greeting",
		},
		cli.StringFlag{
			Name:   "fruit",
			EnvVar: "INPUT_FRUIT",
			Usage:  "Name of fruit",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GetInput(i string) string {
	e := strings.ReplaceAll(i, " ", "_")
	e = strings.ToUpper(e)
	e = "INPUT_" + e
	return strings.TrimSpace(os.Getenv(e))
}
func argoCmd(c *cli.Context) error {
	fmt.Println("fruit repo is: ", c.String("fruit"))
	cmd := exec.Command("argocd", "help")

	// cmd.Stdin = strings.NewReader("and old falcon")

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}

	return nil

}

// func main() {
// 	a := githubactions.New()
// 	fruit := a.GetInput("fruit")
// 	if fruit == "" {
// 		githubactions.Fatalf("missing input 'fruit'")
// 	}
// 	// githubactions.AddMask(fruit)
// 	app := cli.NewApp()
// 	app.Name = "greet"
// 	app.Usage = "fight the loneliness!"
// 	app.Action = runact
// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func runact(c *cli.Context) error {

// 	cfg := Config{
// 		Cluster:    "argocd",
// 		Namespace:  "dev",
// 		Path:       "kustomize/dev",
// 		Project:    "argocd",
// 		Prune:      true,
// 		Repository: "spunkandgritt/kustomize-argocd",
// 		Server:     "localhost",
// 	}

// 	return Runer(context.Background(), cfg)

// }

// type Config struct {
// 	Cluster    string
// 	Namespace  string
// 	Path       string
// 	Project    string
// 	Prune      bool
// 	Repository string
// 	Server     string

// 	Stdout io.Writer
// 	Stderr io.Writer
// }

// func Runer(ctx context.Context, cfg Config) error {

// 	log.Printf("source")
// 	log.Printf("repository: %s", cfg.Repository)

// 	argocdAppCreateCmd := argocdAppCreate(ctx, cfg)
// 	log.Printf("returned the cmd")
// 	stdout, error := argocdAppCreateCmd.StdoutPipe()
// 	log.Printf("got handle on stdout")
// 	if error != nil {
// 		log.Fatal(error)
// 	}
// 	log.Printf("now going to run the command")
// 	if err := argocdAppCreateCmd.Run(); err != nil {
// 		log.Printf("got error")
// 		return err
// 	}
// 	log.Printf("no errors now run funct")
// 	go func() {
// 		log.Printf("inside func")
// 		defer argocdAppCreateCmd.Wait()
// 		scanner := bufio.NewScanner(stdout)
// 		for scanner.Scan() {
// 			s := scanner.Text()
// 			log.Printf("trying to print")
// 			log.Printf(s) // echo to stdout
// 			// Do something with s
// 		}
// 	}()
// 	time.Sleep(15 * time.Second)
// 	log.Printf("all done returning nil")
// 	return nil
// }

// func argocdAppCreate(ctx context.Context, cfg Config) *exec.Cmd {
// 	log.Printf("now in argocd function")
// 	args := []string{
// 		"argocd",
// 		"help",
// 	}
// 	log.Printf("initialized args")
// 	cmd := exec.CommandContext(ctx, args[0], args[1])
// 	cmd.Stdout = cfg.Stdout
// 	cmd.Stderr = cfg.Stderr
// 	log.Printf("returning cmd")
// 	return cmd

// }
