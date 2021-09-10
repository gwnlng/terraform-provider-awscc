// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package devopsguru

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
	registry.AddDataSourceTypeFactory("awscc_devopsguru_notification_channel", notificationChannelDataSourceType)
}

// notificationChannelDataSourceType returns the Terraform awscc_devopsguru_notification_channel data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DevOpsGuru::NotificationChannel resource type.
func notificationChannelDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"config": {
			// Property: Config
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about notification channels you have configured with DevOps Guru.",
			//   "properties": {
			//     "Sns": {
			//       "additionalProperties": false,
			//       "description": "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
			//       "properties": {
			//         "TopicArn": {
			//           "maxLength": 1024,
			//           "minLength": 36,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about notification channels you have configured with DevOps Guru.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sns": {
						// Property: Sns
						Description: "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"topic_arn": {
									// Property: TopicArn
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of a notification channel.",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of a notification channel.",
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
		Description: "Data Source schema for AWS::DevOpsGuru::NotificationChannel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::NotificationChannel").WithTerraformTypeName("awscc_devopsguru_notification_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":    "Config",
		"id":        "Id",
		"sns":       "Sns",
		"topic_arn": "TopicArn",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_devopsguru_notification_channel", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
