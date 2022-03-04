package synapse

import (
	"testing"

	"github.com/aquasecurity/defsec/adapters/terraform/testutil"

	"github.com/aquasecurity/defsec/providers/azure/synapse"
)

func Test_Adapt(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  synapse.Synapse
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: synapse.Synapse{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutil.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := Adapt(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

func Test_adaptWorkspaces(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  []synapse.Workspace
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: []synapse.Workspace{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutil.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := adaptWorkspaces(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}

func Test_adaptWorkspace(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		terraform string
		expected  synapse.Workspace
	}{
		{
			name: "basic",
			terraform: `
resource "" "example" {
    
}
`,
			expected: synapse.Workspace{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := testutil.CreateModulesFromSource(test.terraform, ".tf", t)
			adapted := adaptWorkspace(modules.GetBlocks()[0])
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}