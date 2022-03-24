// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package billingconductor_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSBillingConductorPricingPlanDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::BillingConductor::PricingPlan", "awscc_billingconductor_pricing_plan", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSBillingConductorPricingPlanDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::BillingConductor::PricingPlan", "awscc_billingconductor_pricing_plan", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}