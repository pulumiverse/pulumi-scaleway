package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	scaleway "github.com/scaleway/terraform-provider-scaleway/v2/internal/provider"
)

func NewProvider() plugin.ProviderFunc {
	return scaleway.Provider(scaleway.DefaultConfig())
}
