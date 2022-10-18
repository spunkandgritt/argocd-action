package main

import (
	"argocd/argocd"
	"context"
	"log"
	"os"

	"github.com/sethvargo/go-githubactions"
	"github.com/urfave/cli"
)

func main() {
	a := githubactions.New()
	fruit := a.GetInput("fruit")
	if fruit == "" {
		githubactions.Fatalf("missing input 'fruit'")
	}
	// githubactions.AddMask(fruit)
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {

	cfg := argocd.Config{
		Cluster:    "argocd",
		Namespace:  "dev",
		Path:       "kustomize/dev",
		Project:    "argocd",
		Prune:      true,
		Repository: "spunkandgritt/kustomize-argocd",
		Server:     "localhost",
	}

	return argocd.Run(context.Background(), cfg)

}
