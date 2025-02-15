---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_logs_destination Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Logs::Destination
---

# awscc_logs_destination (Data Source)

Data Source schema for AWS::Logs::Destination



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `destination_name` (String) The name of the destination resource
- `destination_policy` (String) An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
- `role_arn` (String) The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
- `target_arn` (String) The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)


