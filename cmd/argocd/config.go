package argocd

import "io"

type Config struct {
	Cluster    string
	Namespace  string
	Path       string
	Project    string
	Prune      bool
	Repository string
	Server     string

	Stdout io.Writer
	Stderr io.Writer
}
