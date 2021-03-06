package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOktaDataSourceAuthServer_read(t *testing.T) {
	ri := acctest.RandInt()
	mgr := newFixtureManager(authServer)
	config := mgr.GetFixtures("datasource.tf", ri, t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("okta_auth_server.test", "id"),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "name", fmt.Sprintf("testAcc_%d", ri)),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "status", statusActive),
				),
			},
		},
	})
}
