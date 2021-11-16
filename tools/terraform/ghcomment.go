package terraform

import (
	"fmt"

	"github.com/einride/mage-tools/file"
	"github.com/einride/mage-tools/tools"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var ghCommentVersion string

func SetGhCommentVersion(v string) (string, error) {
	ghCommentVersion = v
	return ghCommentVersion, nil
}

func GhReviewTerraformPlan(prNumber string, gcpProject string) {
	terraformPlanFile := "terraform.plan"
	mg.Deps(
		mg.F(tools.Terraform, tfVersion),
		mg.F(tools.GHComment, ghCommentVersion),
		mg.F(file.Exists, terraformPlanFile),
	)

	comment, _ := sh.Output(
		"terraform",
		"show",
		"-no-color",
		terraformPlanFile,
	)
	comment = fmt.Sprintf("```"+"hcl\n%s\n"+"```", comment)
	ghComment := fmt.Sprintf(`
<div>
<img align="right" width="120" src="https://www.terraform.io/assets/images/logo-text-8c3ba8a6.svg" />
<h2>Terraform Plan (%s)</h2>
</div>

%s
`, gcpProject, comment)

	fmt.Println("[ghcomment] commenting terraform plan on pr...")
	err := sh.RunV(
		tools.GHCommentPath,
		"--pr",
		prNumber,
		"--signkey",
		gcpProject,
		"--comment",
		ghComment,
	)
	if err != nil {
		panic(err)
	}
}
