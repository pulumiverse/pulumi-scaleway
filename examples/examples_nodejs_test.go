// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccWebserver(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "ts/webserver"),
		})

	integration.ProgramTest(t, &test)
}


func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	project_id := getProjectId(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		Config: map[string]string{
			"default_organization_id": org_id,
		},
		Dependencies: []string{
			"@jaxxstorm/pulumi-scaleway",
		},
	})

	return baseJS
}