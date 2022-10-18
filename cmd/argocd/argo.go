package argocd

import (
	"context"
	"log"
	"os/exec"
)

func Run(ctx context.Context, cfg Config) error {

	log.Printf("source")
	log.Printf("reposiroty: %s", cfg.Repository)

	argocdAppCreateCmd := argocdAppCreate(ctx, cfg)

	if err := argocdAppCreateCmd.Run(); err != nil {
		return err
	}
	return nil
}

func argocdAppCreate(ctx context.Context, cfg Config) *exec.Cmd {
	args := []string{
		"argocd",
		"help",
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	return cmd

}
