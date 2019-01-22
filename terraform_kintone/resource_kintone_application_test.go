package terraform_kintone

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/client"
	"os"
	"strings"
	"testing"
)

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("KINTONE_HOST"); v == "" {
		t.Fatal("KINTONE_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("KINTONE_USER"); v == "" {
		t.Fatal("KINTONE_USER must be set for acceptance tests")
	}
	if v := os.Getenv("KINTONE_PASSWORD"); v == "" {
		t.Fatal("KINTONE_PASSWORD must be set for acceptance tests")
	}
}

// example.Widget represents a concrete Go type that represents an API resource
func TestAccKintoneApplication_basic(t *testing.T) {
	resourceName := "app_" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	kintoneHost := os.Getenv("KINTONE_HOST")
	kintoneUser := os.Getenv("KINTONE_USER")
	kintonePassword := os.Getenv("KINTONE_PASSWORD")

	apiClient := client.New(kintone.ApiClientConfig{
		Host:     kintoneHost,
		User:     kintoneUser,
		Password: kintonePassword,
	})

	tmp := testAccKintoneApplicationConfig(resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleResourceDestroy(apiClient),
		Steps: []resource.TestStep{
			{
				Config: tmp,
				Check: resource.ComposeTestCheckFunc(
					testAccKintoneApplicationExists(resourceName, apiClient),
				),
			},
		},
	})
}

func testAccKintoneApplicationConfig(name string) string {
	return fmt.Sprintf(`
resource "kintone_application" "%s" {
  name = "%s"
  description = "This is my application!!"
  theme = "BLUE"

  field = { code = "code_app", label = "AppId", type = "SINGLE_LINE_TEXT" }
}`, name, name)
}

func testAccKintoneApplicationExists(resourceName string, apiClient kintone.ApiClient) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources["kintone_application."+resourceName]
		if !ok {
			return fmt.Errorf("kintone application not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no kintone application id is set")
		}

		ctx := context.Background()
		_, err := apiClient.FetchApplication(ctx, kintone.AppId(rs.Primary.ID))
		if err != nil {
			return fmt.Errorf("fetch application error: %s", err)
		}

		return nil
	}
}

func testAccCheckExampleResourceDestroy(apiClient kintone.ApiClient) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "kintone_application" {
				continue
			}

			_, err := apiClient.FetchApplication(ctx, kintone.AppId(rs.Primary.ID))
			if err == nil {
				return fmt.Errorf("kintone application (id: %s) still exists", rs.Primary.ID)
			}

			if !strings.Contains(err.Error(), "404") {
				return err
			}
		}

		return nil
	}
}
