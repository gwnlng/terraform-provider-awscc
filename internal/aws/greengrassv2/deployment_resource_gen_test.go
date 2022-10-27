// Code generated by generators/resource/main.go; DO NOT EDIT.

package greengrassv2_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGreengrassV2Deployment_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GreengrassV2::Deployment", "awscc_greengrassv2_deployment", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}