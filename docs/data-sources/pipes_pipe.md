---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pipes_pipe Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Pipes::Pipe
---

# awscc_pipes_pipe (Data Source)

Data Source schema for AWS::Pipes::Pipe



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `creation_time` (String)
- `current_state` (String)
- `description` (String)
- `desired_state` (String)
- `enrichment` (String)
- `enrichment_parameters` (Attributes) (see [below for nested schema](#nestedatt--enrichment_parameters))
- `last_modified_time` (String)
- `name` (String)
- `role_arn` (String)
- `source` (String)
- `source_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters))
- `state_reason` (String)
- `tags` (Map of String)
- `target` (String)
- `target_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters))

<a id="nestedatt--enrichment_parameters"></a>
### Nested Schema for `enrichment_parameters`

Read-Only:

- `http_parameters` (Attributes) (see [below for nested schema](#nestedatt--enrichment_parameters--http_parameters))
- `input_template` (String)

<a id="nestedatt--enrichment_parameters--http_parameters"></a>
### Nested Schema for `enrichment_parameters.http_parameters`

Read-Only:

- `header_parameters` (Map of String)
- `path_parameter_values` (List of String)
- `query_string_parameters` (Map of String)



<a id="nestedatt--source_parameters"></a>
### Nested Schema for `source_parameters`

Read-Only:

- `active_mq_broker_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--active_mq_broker_parameters))
- `dynamo_db_stream_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--dynamo_db_stream_parameters))
- `filter_criteria` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--filter_criteria))
- `kinesis_stream_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--kinesis_stream_parameters))
- `managed_streaming_kafka_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--managed_streaming_kafka_parameters))
- `rabbit_mq_broker_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--rabbit_mq_broker_parameters))
- `self_managed_kafka_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--self_managed_kafka_parameters))
- `sqs_queue_parameters` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--sqs_queue_parameters))

<a id="nestedatt--source_parameters--active_mq_broker_parameters"></a>
### Nested Schema for `source_parameters.active_mq_broker_parameters`

Read-Only:

- `batch_size` (Number)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--active_mq_broker_parameters--credentials))
- `maximum_batching_window_in_seconds` (Number)
- `queue_name` (String)

<a id="nestedatt--source_parameters--active_mq_broker_parameters--credentials"></a>
### Nested Schema for `source_parameters.active_mq_broker_parameters.credentials`

Read-Only:

- `basic_auth` (String) Optional SecretManager ARN which stores the database credentials



<a id="nestedatt--source_parameters--dynamo_db_stream_parameters"></a>
### Nested Schema for `source_parameters.dynamo_db_stream_parameters`

Read-Only:

- `batch_size` (Number)
- `dead_letter_config` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--dynamo_db_stream_parameters--dead_letter_config))
- `maximum_batching_window_in_seconds` (Number)
- `maximum_record_age_in_seconds` (Number)
- `maximum_retry_attempts` (Number)
- `on_partial_batch_item_failure` (String)
- `parallelization_factor` (Number)
- `starting_position` (String)

<a id="nestedatt--source_parameters--dynamo_db_stream_parameters--dead_letter_config"></a>
### Nested Schema for `source_parameters.dynamo_db_stream_parameters.dead_letter_config`

Read-Only:

- `arn` (String)



<a id="nestedatt--source_parameters--filter_criteria"></a>
### Nested Schema for `source_parameters.filter_criteria`

Read-Only:

- `filters` (Attributes List) (see [below for nested schema](#nestedatt--source_parameters--filter_criteria--filters))

<a id="nestedatt--source_parameters--filter_criteria--filters"></a>
### Nested Schema for `source_parameters.filter_criteria.filters`

Read-Only:

- `pattern` (String)



<a id="nestedatt--source_parameters--kinesis_stream_parameters"></a>
### Nested Schema for `source_parameters.kinesis_stream_parameters`

Read-Only:

- `batch_size` (Number)
- `dead_letter_config` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--kinesis_stream_parameters--dead_letter_config))
- `maximum_batching_window_in_seconds` (Number)
- `maximum_record_age_in_seconds` (Number)
- `maximum_retry_attempts` (Number)
- `on_partial_batch_item_failure` (String)
- `parallelization_factor` (Number)
- `starting_position` (String)
- `starting_position_timestamp` (String)

<a id="nestedatt--source_parameters--kinesis_stream_parameters--dead_letter_config"></a>
### Nested Schema for `source_parameters.kinesis_stream_parameters.dead_letter_config`

Read-Only:

- `arn` (String)



<a id="nestedatt--source_parameters--managed_streaming_kafka_parameters"></a>
### Nested Schema for `source_parameters.managed_streaming_kafka_parameters`

Read-Only:

- `batch_size` (Number)
- `consumer_group_id` (String)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--managed_streaming_kafka_parameters--credentials))
- `maximum_batching_window_in_seconds` (Number)
- `starting_position` (String)
- `topic_name` (String)

<a id="nestedatt--source_parameters--managed_streaming_kafka_parameters--credentials"></a>
### Nested Schema for `source_parameters.managed_streaming_kafka_parameters.credentials`

Read-Only:

- `client_certificate_tls_auth` (String) Optional SecretManager ARN which stores the database credentials
- `sasl_scram_512_auth` (String) Optional SecretManager ARN which stores the database credentials



<a id="nestedatt--source_parameters--rabbit_mq_broker_parameters"></a>
### Nested Schema for `source_parameters.rabbit_mq_broker_parameters`

Read-Only:

- `batch_size` (Number)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--rabbit_mq_broker_parameters--credentials))
- `maximum_batching_window_in_seconds` (Number)
- `queue_name` (String)
- `virtual_host` (String)

<a id="nestedatt--source_parameters--rabbit_mq_broker_parameters--credentials"></a>
### Nested Schema for `source_parameters.rabbit_mq_broker_parameters.credentials`

Read-Only:

- `basic_auth` (String) Optional SecretManager ARN which stores the database credentials



<a id="nestedatt--source_parameters--self_managed_kafka_parameters"></a>
### Nested Schema for `source_parameters.self_managed_kafka_parameters`

Read-Only:

- `additional_bootstrap_servers` (List of String)
- `batch_size` (Number)
- `consumer_group_id` (String)
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--self_managed_kafka_parameters--credentials))
- `maximum_batching_window_in_seconds` (Number)
- `server_root_ca_certificate` (String) Optional SecretManager ARN which stores the database credentials
- `starting_position` (String)
- `topic_name` (String)
- `vpc` (Attributes) (see [below for nested schema](#nestedatt--source_parameters--self_managed_kafka_parameters--vpc))

<a id="nestedatt--source_parameters--self_managed_kafka_parameters--credentials"></a>
### Nested Schema for `source_parameters.self_managed_kafka_parameters.credentials`

Read-Only:

- `basic_auth` (String) Optional SecretManager ARN which stores the database credentials
- `client_certificate_tls_auth` (String) Optional SecretManager ARN which stores the database credentials
- `sasl_scram_256_auth` (String) Optional SecretManager ARN which stores the database credentials
- `sasl_scram_512_auth` (String) Optional SecretManager ARN which stores the database credentials


<a id="nestedatt--source_parameters--self_managed_kafka_parameters--vpc"></a>
### Nested Schema for `source_parameters.self_managed_kafka_parameters.vpc`

Read-Only:

- `security_group` (List of String) List of SecurityGroupId.
- `subnets` (List of String) List of SubnetId.



<a id="nestedatt--source_parameters--sqs_queue_parameters"></a>
### Nested Schema for `source_parameters.sqs_queue_parameters`

Read-Only:

- `batch_size` (Number)
- `maximum_batching_window_in_seconds` (Number)



<a id="nestedatt--target_parameters"></a>
### Nested Schema for `target_parameters`

Read-Only:

- `batch_job_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters))
- `cloudwatch_logs_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--cloudwatch_logs_parameters))
- `ecs_task_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters))
- `event_bridge_event_bus_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--event_bridge_event_bus_parameters))
- `http_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--http_parameters))
- `input_template` (String)
- `kinesis_stream_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--kinesis_stream_parameters))
- `lambda_function_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--lambda_function_parameters))
- `redshift_data_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--redshift_data_parameters))
- `sage_maker_pipeline_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--sage_maker_pipeline_parameters))
- `sqs_queue_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--sqs_queue_parameters))
- `step_function_state_machine_parameters` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--step_function_state_machine_parameters))

<a id="nestedatt--target_parameters--batch_job_parameters"></a>
### Nested Schema for `target_parameters.batch_job_parameters`

Read-Only:

- `array_properties` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--array_properties))
- `container_overrides` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--container_overrides))
- `depends_on` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--depends_on))
- `job_definition` (String)
- `job_name` (String)
- `parameters` (Map of String)
- `retry_strategy` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--retry_strategy))

<a id="nestedatt--target_parameters--batch_job_parameters--array_properties"></a>
### Nested Schema for `target_parameters.batch_job_parameters.array_properties`

Read-Only:

- `size` (Number)


<a id="nestedatt--target_parameters--batch_job_parameters--container_overrides"></a>
### Nested Schema for `target_parameters.batch_job_parameters.container_overrides`

Read-Only:

- `command` (List of String)
- `environment` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--container_overrides--environment))
- `instance_type` (String)
- `resource_requirements` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--batch_job_parameters--container_overrides--resource_requirements))

<a id="nestedatt--target_parameters--batch_job_parameters--container_overrides--environment"></a>
### Nested Schema for `target_parameters.batch_job_parameters.container_overrides.resource_requirements`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedatt--target_parameters--batch_job_parameters--container_overrides--resource_requirements"></a>
### Nested Schema for `target_parameters.batch_job_parameters.container_overrides.resource_requirements`

Read-Only:

- `type` (String)
- `value` (String)



<a id="nestedatt--target_parameters--batch_job_parameters--depends_on"></a>
### Nested Schema for `target_parameters.batch_job_parameters.depends_on`

Read-Only:

- `job_id` (String)
- `type` (String)


<a id="nestedatt--target_parameters--batch_job_parameters--retry_strategy"></a>
### Nested Schema for `target_parameters.batch_job_parameters.retry_strategy`

Read-Only:

- `attempts` (Number)



<a id="nestedatt--target_parameters--cloudwatch_logs_parameters"></a>
### Nested Schema for `target_parameters.cloudwatch_logs_parameters`

Read-Only:

- `log_stream_name` (String)
- `timestamp` (String)


<a id="nestedatt--target_parameters--ecs_task_parameters"></a>
### Nested Schema for `target_parameters.ecs_task_parameters`

Read-Only:

- `capacity_provider_strategy` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--capacity_provider_strategy))
- `enable_ecs_managed_tags` (Boolean)
- `enable_execute_command` (Boolean)
- `group` (String)
- `launch_type` (String)
- `network_configuration` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--network_configuration))
- `overrides` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides))
- `placement_constraints` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--placement_constraints))
- `placement_strategy` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--placement_strategy))
- `platform_version` (String)
- `propagate_tags` (String)
- `reference_id` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--tags))
- `task_count` (Number)
- `task_definition_arn` (String)

<a id="nestedatt--target_parameters--ecs_task_parameters--capacity_provider_strategy"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.capacity_provider_strategy`

Read-Only:

- `base` (Number)
- `capacity_provider` (String)
- `weight` (Number)


<a id="nestedatt--target_parameters--ecs_task_parameters--network_configuration"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.network_configuration`

Read-Only:

- `awsvpc_configuration` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--network_configuration--awsvpc_configuration))

<a id="nestedatt--target_parameters--ecs_task_parameters--network_configuration--awsvpc_configuration"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.network_configuration.awsvpc_configuration`

Read-Only:

- `assign_public_ip` (String)
- `security_groups` (List of String)
- `subnets` (List of String)



<a id="nestedatt--target_parameters--ecs_task_parameters--overrides"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides`

Read-Only:

- `container_overrides` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--container_overrides))
- `cpu` (String)
- `ephemeral_storage` (Attributes) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--ephemeral_storage))
- `execution_role_arn` (String)
- `inference_accelerator_overrides` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--inference_accelerator_overrides))
- `memory` (String)
- `task_role_arn` (String)

<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--container_overrides"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn`

Read-Only:

- `command` (List of String)
- `cpu` (Number)
- `environment` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--environment))
- `environment_files` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--environment_files))
- `memory` (Number)
- `memory_reservation` (Number)
- `name` (String)
- `resource_requirements` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--resource_requirements))

<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--environment"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn.environment`

Read-Only:

- `name` (String)
- `value` (String)


<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--environment_files"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn.environment_files`

Read-Only:

- `type` (String)
- `value` (String)


<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--task_role_arn--resource_requirements"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn.resource_requirements`

Read-Only:

- `type` (String)
- `value` (String)



<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--ephemeral_storage"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn`

Read-Only:

- `size_in_gi_b` (Number)


<a id="nestedatt--target_parameters--ecs_task_parameters--overrides--inference_accelerator_overrides"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.overrides.task_role_arn`

Read-Only:

- `device_name` (String)
- `device_type` (String)



<a id="nestedatt--target_parameters--ecs_task_parameters--placement_constraints"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.placement_constraints`

Read-Only:

- `expression` (String)
- `type` (String)


<a id="nestedatt--target_parameters--ecs_task_parameters--placement_strategy"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.placement_strategy`

Read-Only:

- `field` (String)
- `type` (String)


<a id="nestedatt--target_parameters--ecs_task_parameters--tags"></a>
### Nested Schema for `target_parameters.ecs_task_parameters.tags`

Read-Only:

- `key` (String)
- `value` (String)



<a id="nestedatt--target_parameters--event_bridge_event_bus_parameters"></a>
### Nested Schema for `target_parameters.event_bridge_event_bus_parameters`

Read-Only:

- `detail_type` (String)
- `endpoint_id` (String)
- `resources` (List of String)
- `source` (String)
- `time` (String)


<a id="nestedatt--target_parameters--http_parameters"></a>
### Nested Schema for `target_parameters.http_parameters`

Read-Only:

- `header_parameters` (Map of String)
- `path_parameter_values` (List of String)
- `query_string_parameters` (Map of String)


<a id="nestedatt--target_parameters--kinesis_stream_parameters"></a>
### Nested Schema for `target_parameters.kinesis_stream_parameters`

Read-Only:

- `partition_key` (String)


<a id="nestedatt--target_parameters--lambda_function_parameters"></a>
### Nested Schema for `target_parameters.lambda_function_parameters`

Read-Only:

- `invocation_type` (String)


<a id="nestedatt--target_parameters--redshift_data_parameters"></a>
### Nested Schema for `target_parameters.redshift_data_parameters`

Read-Only:

- `database` (String) Redshift Database
- `db_user` (String) Database user name
- `secret_manager_arn` (String) Optional SecretManager ARN which stores the database credentials
- `sqls` (List of String) A list of SQLs.
- `statement_name` (String) A name for Redshift DataAPI statement which can be used as filter of ListStatement.
- `with_event` (Boolean)


<a id="nestedatt--target_parameters--sage_maker_pipeline_parameters"></a>
### Nested Schema for `target_parameters.sage_maker_pipeline_parameters`

Read-Only:

- `pipeline_parameter_list` (Attributes List) (see [below for nested schema](#nestedatt--target_parameters--sage_maker_pipeline_parameters--pipeline_parameter_list))

<a id="nestedatt--target_parameters--sage_maker_pipeline_parameters--pipeline_parameter_list"></a>
### Nested Schema for `target_parameters.sage_maker_pipeline_parameters.pipeline_parameter_list`

Read-Only:

- `name` (String)
- `value` (String)



<a id="nestedatt--target_parameters--sqs_queue_parameters"></a>
### Nested Schema for `target_parameters.sqs_queue_parameters`

Read-Only:

- `message_deduplication_id` (String)
- `message_group_id` (String)


<a id="nestedatt--target_parameters--step_function_state_machine_parameters"></a>
### Nested Schema for `target_parameters.step_function_state_machine_parameters`

Read-Only:

- `invocation_type` (String)

