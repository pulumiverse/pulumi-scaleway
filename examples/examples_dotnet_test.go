//go:build dotnet || all
// +build dotnet all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccWebserver(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dotnet/server"),
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	project_id := getProjectId(t)
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"project_id": project_id,
		},
		Dependencies: []string{
			"Pulumi.Scaleway",
		},
	})

	return baseCsharp
}
