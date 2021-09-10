// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_local_gateway_route_table_vpc_associations", localGatewayRouteTableVPCAssociationsDataSourceType)
}

// localGatewayRouteTableVPCAssociationsDataSourceType returns the Terraform awscc_ec2_local_gateway_route_table_vpc_associations data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVPCAssociation resource type.
func localGatewayRouteTableVPCAssociationsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        providertypes.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::EC2::LocalGatewayRouteTableVPCAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVPCAssociation").WithTerraformTypeName("awscc_ec2_local_gateway_route_table_vpc_associations")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ec2_local_gateway_route_table_vpc_associations", "schema", hclog.Fmt("%v", schema))

	return pluralDataSourceType, nil
}
