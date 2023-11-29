package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExpResource(t *testing.T) {
	t.Parallel()
	const version = "v1.0.0"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccExpResource(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("exp_resource.test", "cnt", "1"),
				),
			},
		},
	})
}

func testAccExpResource() string {
	return `
provider "exp" {}

resource "exp_resource" "test" {}
`
}
