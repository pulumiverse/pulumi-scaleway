// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticmetal

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-scaleway/sdk/go/scaleway/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "scaleway:elasticmetal/ip:Ip":
		r = &Ip{}
	case "scaleway:elasticmetal/ipMacAddress:IpMacAddress":
		r = &IpMacAddress{}
	case "scaleway:elasticmetal/server:Server":
		r = &Server{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"scaleway",
		"elasticmetal/ip",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"elasticmetal/ipMacAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"elasticmetal/server",
		&module{version},
	)
}
