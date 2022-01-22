package mgmarkdownfmt

import (
	"context"
	"os/exec"

	"github.com/magefile/mage/mg"
	"go.einride.tech/mage-tools/mgtool"
)

const version = "75134924a9fd3335f76a9709314c5f5cef4d6ddc"

// nolint: gochecknoglobals
var commandPath string

func Command(ctx context.Context, args ...string) *exec.Cmd {
	mg.CtxDeps(ctx, Prepare.MarkdownFmt)
	return mgtool.Command(ctx, commandPath, args...)
}

type Prepare mgtool.Prepare

func (Prepare) MarkdownFmt(ctx context.Context) error {
	binary, err := mgtool.GoInstall(ctx, "github.com/shurcooL/markdownfmt", version)
	if err != nil {
		return err
	}
	commandPath = binary
	return nil
}