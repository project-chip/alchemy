package testplan

import (
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func renderClusterTestPlan(doc *ascii.Doc, cluster *matter.Cluster) (output string, err error) {
	var out strings.Builder
	err = renderHeader(cluster, &out)
	if err != nil {
		return
	}
	err = renderServer(doc, cluster, &out)
	if err != nil {
		return
	}
	out.WriteString(testCases)
	err = renderGlobalAttributesTestCase(doc, cluster, &out)
	if err != nil {
		return
	}
	renderAttributesTest(cluster, &out)
	output = out.String()
	return
}

var testCases = `== Test Case List

|===
| *TC UUID*         | *Test Case Name*
| TC-{picsCode}-1.1 | Global Attributes with {DUT_Server}
| TC-{picsCode}-2.1 | Attributes with Server as DUT
| TC-{picsCode}-2.2 | Primary Functionality with Server as DUT
|===


`
