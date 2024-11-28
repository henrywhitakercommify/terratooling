package terratooling

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// Type alias for terraform options/variables to make them a bit easier to work with
//
// e.g. Options{}.With(Options{"resource_group_name": "uksouth-123")
type Options map[string]any

// Merges the provided `with` options with the set the existing options
func (o Options) With(with Options) Options {
	for k, v := range with {
		o[k] = v
	}
	return o
}

// Removes the given keys provided from the options
func (o Options) Without(keys ...string) Options {
	for _, key := range keys {
		delete(o, key)
	}
	return o
}

// Setup a new terraform env. This will copy the terraform module provided to it into
// a new temp directory so two tests running in parallel won't have the same local terraform state.
// root is relative to the test directory, so for tests in the `tests` directory, and the module
// setup in `example`:
//
// - example/
// - tests/
// - main.tf
// - providers.tf
// - variables.tf
//
// You would call Setup(t, "..", "example", opts)
func Setup(t *testing.T, root, module string, opts Options) *terraform.Options {
	tempFolder := test_structure.CopyTerraformFolderToTemp(t, "..", "examples")
	return &terraform.Options{
		TerraformDir: tempFolder,
		Vars:         opts,
	}
}

// Returns a fake azure subnet id
func FakeSubnetId() string {
	return "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/virtualNetworks/virtualNetworksValue/subnets/subnetValue"
}

// Returns a fake azure private dns zone id
func FakePrivateDnsZoneId() string {
	return "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.Network/privateDnsZones/privateDnsZoneValue"
}

// Returns a fake azure eventhub id
func FakeEventhubId() string {
	return "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.EventHub/namespaces/namespaceValue/authorizationRules/authorizationRuleValue"
}
