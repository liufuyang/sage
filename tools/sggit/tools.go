package sggit

import (
	"bytes"
	"context"
	"fmt"

	"go.einride.tech/sage/sg"
)

func VerifyNoDiff(ctx context.Context) error {
	cmd := sg.Command(ctx, "git", "status", "--porcelain")
	var status bytes.Buffer
	cmd.Stdout = &status
	if err := cmd.Run(); err != nil {
		return err
	}
	if status.String() != "" {
		output := sg.Output(sg.Command(ctx, "git", "diff", "--patch"))
		if output != "" {
			return fmt.Errorf("staging area is dirty, please add all files created by the build to .gitignore: %s", output)
		}
	}
	return nil
}

func Version(ctx context.Context) string {
	revision := sg.Output(
		sg.Command(ctx, "git", "rev-parse", "--verify", "--short", "HEAD"),
	)
	diff := sg.Output(
		sg.Command(ctx, "git", "status", "--porcelain"),
	)
	if diff != "" {
		revision += "-dirty"
	}
	return revision
}
