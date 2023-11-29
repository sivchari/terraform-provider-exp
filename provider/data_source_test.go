package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSource(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: protoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSource(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.exp_datasource.test", "id", "test"),
					resource.TestCheckResourceAttr("data.exp_datasource.test", "cnt", "1"),
				),
			},
		},
	})
}

func testAccDataSource() string {
	return `
data "exp_datasource" "test" {
    id = "test"
}`
}
