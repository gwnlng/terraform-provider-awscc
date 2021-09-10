---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iot_security_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A security profile defines a set of expected behaviors for devices in your account.
---

# awscc_iot_security_profile (Resource)

A security profile defines a set of expected behaviors for devices in your account.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **additional_metrics_to_retain_v2** (Attributes Set) A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here. (see [below for nested schema](#nestedatt--additional_metrics_to_retain_v2))
- **alert_targets** (Attributes Map) Specifies the destinations to which alerts are sent. (see [below for nested schema](#nestedatt--alert_targets))
- **behaviors** (Attributes Set) Specifies the behaviors that, when violated by a device (thing), cause an alert. (see [below for nested schema](#nestedatt--behaviors))
- **security_profile_description** (String) A description of the security profile.
- **security_profile_name** (String) A unique identifier for the security profile.
- **tags** (Attributes Set, Max: 50) Metadata that can be used to manage the security profile. (see [below for nested schema](#nestedatt--tags))
- **target_arns** (Set of String) A set of target ARNs that the security profile is attached to.

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **security_profile_arn** (String) The ARN (Amazon resource name) of the created security profile.

<a id="nestedatt--additional_metrics_to_retain_v2"></a>
### Nested Schema for `additional_metrics_to_retain_v2`

Optional:

- **metric** (String) What is measured by the behavior.
- **metric_dimension** (Attributes) The dimension of a metric. (see [below for nested schema](#nestedatt--additional_metrics_to_retain_v2--metric_dimension))

<a id="nestedatt--additional_metrics_to_retain_v2--metric_dimension"></a>
### Nested Schema for `additional_metrics_to_retain_v2.metric_dimension`

Optional:

- **dimension_name** (String) A unique identifier for the dimension.
- **operator** (String) Defines how the dimensionValues of a dimension are interpreted.



<a id="nestedatt--alert_targets"></a>
### Nested Schema for `alert_targets`

Optional:

- **alert_target_arn** (String) The ARN of the notification target to which alerts are sent.
- **role_arn** (String) The ARN of the role that grants permission to send alerts to the notification target.


<a id="nestedatt--behaviors"></a>
### Nested Schema for `behaviors`

Optional:

- **criteria** (Attributes) The criteria by which the behavior is determined to be normal. (see [below for nested schema](#nestedatt--behaviors--criteria))
- **metric** (String) What is measured by the behavior.
- **metric_dimension** (Attributes) The dimension of a metric. (see [below for nested schema](#nestedatt--behaviors--metric_dimension))
- **name** (String) The name for the behavior.
- **suppress_alerts** (Boolean) Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.

<a id="nestedatt--behaviors--criteria"></a>
### Nested Schema for `behaviors.criteria`

Optional:

- **comparison_operator** (String) The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).
- **consecutive_datapoints_to_alarm** (Number) If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.
- **consecutive_datapoints_to_clear** (Number) If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.
- **duration_seconds** (Number) Use this to specify the time duration over which the behavior is evaluated.
- **ml_detection_config** (Attributes) The configuration of an ML Detect Security Profile. (see [below for nested schema](#nestedatt--behaviors--criteria--ml_detection_config))
- **statistical_threshold** (Attributes) A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior. (see [below for nested schema](#nestedatt--behaviors--criteria--statistical_threshold))
- **value** (Attributes) The value to be compared with the metric. (see [below for nested schema](#nestedatt--behaviors--criteria--value))

<a id="nestedatt--behaviors--criteria--ml_detection_config"></a>
### Nested Schema for `behaviors.criteria.ml_detection_config`

Optional:

- **confidence_level** (String) The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.


<a id="nestedatt--behaviors--criteria--statistical_threshold"></a>
### Nested Schema for `behaviors.criteria.statistical_threshold`

Optional:

- **statistic** (String) The percentile which resolves to a threshold value by which compliance with a behavior is determined


<a id="nestedatt--behaviors--criteria--value"></a>
### Nested Schema for `behaviors.criteria.value`

Optional:

- **cidrs** (Set of String) If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.
- **count** (String) If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.
- **number** (Number) The numeral value of a metric.
- **numbers** (Set of Number) The numeral values of a metric.
- **ports** (Set of Number) If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.
- **strings** (Set of String) The string values of a metric.



<a id="nestedatt--behaviors--metric_dimension"></a>
### Nested Schema for `behaviors.metric_dimension`

Optional:

- **dimension_name** (String) A unique identifier for the dimension.
- **operator** (String) Defines how the dimensionValues of a dimension are interpreted.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The tag's key.
- **value** (String) The tag's value.

