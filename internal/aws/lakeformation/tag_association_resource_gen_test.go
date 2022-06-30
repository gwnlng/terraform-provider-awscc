// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLakeFormationTagAssociation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::LakeFormation::TagAssociation", "awscc_lakeformation_tag_association", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}