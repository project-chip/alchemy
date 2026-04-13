package action

import (
	"encoding/json"
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
	switch templateName {
	case "disco/combined":
		t, err = templates.LoadDiscoCombinedTemplate()
	case "disco/unpatched":
		t, err = templates.LoadDiscoUnpatchedTemplate()
	default:
		t, err = templates.LoadTemplate(templateName)
	}

	if err != nil {
		err = fmt.Errorf("error loading template %s: %w", templateName, err)
		return
	}

	templateContext := make(map[string]any)

	templateDataStr := githubactions.GetInput("template_data")
	if templateDataStr != "" {
		var templateData map[string]any
		err = json.Unmarshal([]byte(templateDataStr), &templateData)
		if err != nil {
			err = fmt.Errorf("error parsing template_data JSON: %w", err)
			return
		}
		for k, v := range templateData {
			templateContext[k] = v
		}
	}

	fileDataStr := githubactions.GetInput("file_data")
	if fileDataStr != "" {
		var fileData map[string]string
		err = json.Unmarshal([]byte(fileDataStr), &fileData)
		if err != nil {
			err = fmt.Errorf("error parsing file_data JSON: %w", err)
			return
		}
		for k, path := range fileData {
			if path == "" {
				continue
			}
			var content []byte
			content, err = os.ReadFile(path)
			if err != nil {
				err = fmt.Errorf("error reading file %s for key %s: %w", path, k, err)
				return
			}
			if len(content) < 60000 {
				templateContext[k] = string(content)
			}
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
