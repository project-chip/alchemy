package action

import (
	"fmt"
	"os"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/cmd/action/github/templates"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/sethvargo/go-githubactions"
)

type Comment struct {
}

func (c *Comment) Run(cc *cli.Context) (err error) {

	action := githubactions.New()

	action.Infof("Alchemy %s", config.Version())

	templateName := githubactions.GetInput("template_name")
	if templateName == "" {
		githubactions.Fatalf("missing 'template_name'")
		return
	}

	var t *raymond.Template
	t, err = templates.LoadTemplate(templateName)

	if err != nil {
		err = fmt.Errorf("error loading template %s: %w", templateName, err)
		return
	}

	templateContext := make(map[string]any)
	templateContext["patch_url"] = githubactions.GetInput("patch_url")

	patchPath := githubactions.GetInput("patch_path")

	if patchPath != "" {
		var patch []byte
		patch, err = os.ReadFile(patchPath)
		if err != nil {
			err = fmt.Errorf("error loading patch file: %w", err)
			return
		}
		if len(patch) < 60000 {
			templateContext["diff"] = string(patch)
		}
	}

	var comment string
	comment, err = t.Exec(templateContext)
	if err != nil {
		err = fmt.Errorf("error rendering disco comment template: %w", err)
		return
	}

	action.SetOutput("comment", comment)

	return
}
