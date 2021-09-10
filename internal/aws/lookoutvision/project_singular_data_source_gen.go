// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lookoutvision

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lookoutvision_project", projectDataSourceType)
}

// projectDataSourceType returns the Terraform awscc_lookoutvision_project data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LookoutVision::Project resource type.
func projectDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon Lookout for Vision project",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the Amazon Lookout for Vision project",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::LookoutVision::Project",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutVision::Project").WithTerraformTypeName("awscc_lookoutvision_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":          "Arn",
		"project_name": "ProjectName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_lookoutvision_project", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
