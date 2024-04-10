//go:build !github

package cmd

import "os"

func handleError(err error) {
	os.Exit(1)
}
