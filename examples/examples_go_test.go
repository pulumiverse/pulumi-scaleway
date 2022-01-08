// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccWebserverGo(t *testing.T) {
	project_id := getProjectId(t)
	test := integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "go/server"),
		Config: map[string]string{
			"project_id": project_id,
		},
		Dependencies: []string{
			"github.com/jaxxstorm/pulumi-scaleway/sdk",
		},
	}

	integration.ProgramTest(t, &test)
}
