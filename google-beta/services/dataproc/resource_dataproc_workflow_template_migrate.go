// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dataproc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
)

func resourceDataprocWorkflowTemplateResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"jobs": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The Directed Acyclic Graph of Jobs to submit.",
				Elem:        DataprocWorkflowTemplateJobsSchemaV0(),
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For `projects.locations.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`",
			},

			"placement": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. WorkflowTemplate scheduling information.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementSchemaV0(),
			},

			"dag_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes (\"600s\") to 24 hours (\"86400s\"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a [managed cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.",
			},

			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
			},

			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.",
				Elem:        DataprocWorkflowTemplateParametersSchemaV0(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Output only. The current version of this workflow template.",
				Deprecated:  "version is not useful as a configurable field, and will be removed in the future.",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The time template was created.",
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field `effective_labels` for all of the labels present on the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"terraform_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: "The combination of labels configured directly on the resource and default labels configured on the provider.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The time template was last updated.",
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"step_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job `goog-dataproc-workflow-step-id` label, and in prerequisiteStepIds field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.",
			},

			"hadoop_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a Hadoop job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHadoopJobSchemaV0(),
			},

			"hive_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a Hive job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHiveJobSchemaV0(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given job.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"pig_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a Pig job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobSchemaV0(),
			},

			"prerequisite_step_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"presto_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a Presto job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobSchemaV0(),
			},

			"pyspark_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a PySpark job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPysparkJobSchemaV0(),
			},

			"scheduling": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job scheduling configuration.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSchedulingSchemaV0(),
			},

			"spark_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a Spark job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkJobSchemaV0(),
			},

			"spark_r_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a SparkR job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkRJobSchemaV0(),
			},

			"spark_sql_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Job is a SparkSql job.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHadoopJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of archives to be extracted in the working directory of Hadoop drivers and tasks. Supported file types: .jar, .tar, .tar.gz, .tgz, or .zip.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The arguments to pass to the driver. Do not include arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS (Hadoop Compatible Filesystem) URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Jar file URIs to add to the CLASSPATHs of the Hadoop driver and tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHadoopJobLoggingConfigSchemaV0(),
			},

			"main_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the driver's main class. The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`.",
			},

			"main_jar_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'",
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHiveJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.",
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the script that contains Hive queries.",
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A list of queries.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHiveJobQueryListSchemaV0(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Mapping of query variable names to values (equivalent to the Hive command: `SET name=\"value\";`).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHiveJobQueryListSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.",
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobLoggingConfigSchemaV0(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the script that contains the Pig queries.",
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A list of queries.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobQueryListSchemaV0(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Mapping of query variable names to values (equivalent to the Pig command: `name=[value]`).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobQueryListSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_tags": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Presto client tags to attach to this query",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.",
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobLoggingConfigSchemaV0(),
			},

			"output_format": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The format in which query output will be displayed. See the Presto documentation for supported output formats",
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values. Used to set Presto [session properties](https://prestodb.io/docs/current/sql/set-session.html) Equivalent to using the --session flag in the Presto CLI",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the script that contains SQL queries.",
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A list of queries.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobQueryListSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobQueryListSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPysparkJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"main_python_file_uri": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file.",
			},

			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPysparkJobLoggingConfigSchemaV0(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"python_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSchedulingSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_failures_per_hour": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. A job may be reported as thrashing if driver exits with non-zero code 4 times within 10 minute window. Maximum value is 10.",
			},

			"max_failures_total": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. Maximum value is 240.",
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkJobLoggingConfigSchemaV0(),
			},

			"main_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the driver's main class. The jar file that contains the class must be in the default CLASSPATH or specified in `jar_file_uris`.",
			},

			"main_jar_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the jar file that contains the main class.",
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkRJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"main_r_file_uri": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The HCFS URI of the main R file to use as the driver. Must be a .R file.",
			},

			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkRJobLoggingConfigSchemaV0(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure SparkR. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The runtime log config for job execution.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigSchemaV0(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Dataproc API may be overwritten.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The HCFS URI of the script that contains SQL queries.",
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A list of queries.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobQueryListSchemaV0(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET `name=\"value\";`).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The per-package log levels for the driver. This may include \"root\" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobQueryListSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The queries to execute. You do not need to end a query expression with a semicolon. Multiple queries can be specified in one string by separating each with a semicolon. Here is an example of a Dataproc API snippet that uses a QueryList to specify a HiveJob: \"hiveJob\": { \"queryList\": { \"queries\": [ \"query1\", \"query2\", \"query3;query4\", ] } }",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_selector": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A selector that chooses target cluster for jobs based on metadata. The selector is evaluated at the time each job is submitted.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementClusterSelectorSchemaV0(),
			},

			"managed_cluster": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A cluster that is managed by the workflow.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementClusterSelectorSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_labels": {
				Type:        schema.TypeMap,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The cluster labels. Cluster must have all labels to match.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The zone where workflow process executes. This parameter does not affect the selection of the cluster. If unspecified, the zone of the first cluster matching the selector is used.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The cluster name prefix. A unique cluster name will be formed by appending a random suffix. The name must contain only lower-case letters (a-z), numbers (0-9), and hyphens (-). Must begin with a letter. Cannot begin or end with hyphen. Must consist of between 2 and 35 characters.",
			},

			"config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. The cluster configuration.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSchemaV0(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given cluster.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigSchemaV0(),
			},

			"encryption_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Encryption settings for the cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigSchemaV0(),
			},

			"endpoint_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Port/endpoint configuration for this cluster",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigSchemaV0(),
			},

			"gce_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The shared Compute Engine config settings for all instances in a cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigSchemaV0(),
			},
			"gke_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. BETA. The Kubernetes Engine config for Dataproc clusters deployed to Kubernetes. Setting this is considered mutually exclusive with Compute Engine-based options such as `gce_cluster_config`, `master_config`, `worker_config`, `secondary_worker_config`, and `autoscaling_config`.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigSchemaV0(),
			},
			"initialization_actions": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's `role` metadata to run an executable on a master or worker node, as shown below using `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if [[ \"${ROLE}\" == 'Master' ]]; then ... master specific actions ... else ... worker specific actions ... fi",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsSchemaV0(),
			},

			"lifecycle_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Lifecycle setting for the cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigSchemaV0(),
			},

			"master_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine config settings for the master instance in a cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigSchemaV0(),
			},
			"metastore_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Metastore configuration.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigSchemaV0(),
			},
			"secondary_worker_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine config settings for additional worker instances in a cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigSchemaV0(),
			},

			"security_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Security settings for the cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigSchemaV0(),
			},

			"software_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The config settings for software inside the cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigSchemaV0(),
			},

			"staging_bucket": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
			},

			"temp_bucket": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**",
			},

			"worker_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine config settings for worker instances in a cluster.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policy": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` Note that the policy must be in the same project and Dataproc region.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gce_pd_kms_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_http_port_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.",
			},

			"http_ports": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"internal_ip_only": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses.",
			},

			"metadata": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the \"default\" network of the project is used, if it exists. Cannot be a \"Custom Subnet Network\" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` * `default`",
			},

			"node_group_affinity": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Node Group Affinity for sole-tenant clusters.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinitySchemaV0(),
			},

			"private_ipv6_google_access": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL",
			},

			"reservation_affinity": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Reservation Affinity for consuming Zonal reservation.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinitySchemaV0(),
			},

			"service_account": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.",
			},

			"service_account_scopes": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"shielded_instance_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Shielded Instance Config for clusters using Compute Engine Shielded VMs.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigSchemaV0(),
			},

			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0` * `sub0`",
			},

			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "The Compute Engine tags to add to all instances (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The zone where the Compute Engine cluster will be located. On a create request, it is required in the \"global\" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` * `projects/[project_id]/zones/[zone]` * `us-central1-f`",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinitySchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Required. The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `node-group-1`",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinitySchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"consume_reservation_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION",
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Corresponds to the label key of reservation resource.",
			},

			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Corresponds to the label values of reservation resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_integrity_monitoring": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Defines whether instances have integrity monitoring enabled. Integrity monitoring compares the most recent boot measurements to the integrity policy baseline and returns a pair of pass/fail results depending on whether they match or not.",
			},

			"enable_secure_boot": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Defines whether the instances have Secure Boot enabled. Secure Boot helps ensure that the system only runs authentic software by verifying the digital signature of all boot components, and halting the boot process if signature verification fails.",
			},

			"enable_vtpm": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Defines whether the instance have the vTPM enabled. Virtual Trusted Platform Module protects objects like keys, certificates and enables Measured Boot by performing the measurements needed to create a known good boot baseline, called the integrity policy baseline.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"namespaced_gke_deployment_target": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A target for the deployment.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. A namespace within the GKE cluster to deploy into.",
			},

			"target_gke_cluster": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"executable_file": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Required. Cloud Storage URI of executable file.",
			},

			"execution_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_delete_time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
			},

			"auto_delete_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
			},

			"idle_delete_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
			},

			"idle_start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerators": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine accelerator configuration for these instances.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsSchemaV0(),
			},

			"disk_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Disk option config settings.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigSchemaV0(),
			},

			"image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
			},

			"min_cpu_platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
			},

			"num_instances": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
			},

			"preemptibility": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
			},

			"instance_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"is_preemptible": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Output only. Specifies that this instance group contains preemptible instances.",
			},

			"managed_group_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of the accelerator cards of this type exposed to this instance.",
			},

			"accelerator_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"boot_disk_size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Size in GB of the boot disk (default is 500GB).",
			},

			"boot_disk_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
			},

			"num_local_ssds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_group_manager_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Group Manager for this group.",
			},

			"instance_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dataproc_metastore_service": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Required. Resource name of an existing Dataproc Metastore service. Example: * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerators": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine accelerator configuration for these instances.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsSchemaV0(),
			},

			"disk_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Disk option config settings.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigSchemaV0(),
			},

			"image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
			},

			"min_cpu_platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
			},

			"num_instances": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
			},

			"preemptibility": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
			},

			"instance_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"is_preemptible": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Output only. Specifies that this instance group contains preemptible instances.",
			},

			"managed_group_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of the accelerator cards of this type exposed to this instance.",
			},

			"accelerator_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"boot_disk_size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Size in GB of the boot disk (default is 500GB).",
			},

			"boot_disk_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
			},

			"num_local_ssds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_group_manager_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Group Manager for this group.",
			},

			"instance_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kerberos_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Kerberos related configuration.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cross_realm_trust_admin_server": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
			},

			"cross_realm_trust_kdc": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship.",
			},

			"cross_realm_trust_realm": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust.",
			},

			"cross_realm_trust_shared_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship.",
			},

			"enable_kerberos": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster.",
			},

			"kdc_db_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database.",
			},

			"key_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc.",
			},

			"keystore": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
			},

			"keystore_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc.",
			},

			"kms_key": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The uri of the KMS key used to encrypt various sensitive files.",
			},

			"realm": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm.",
			},

			"root_principal_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password.",
			},

			"tgt_lifetime_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used.",
			},

			"truststore": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate.",
			},

			"truststore_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions), such as \"1.2\" (including a subminor version, such as \"1.2.29\"), or the [\"preview\" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version.",
			},

			"optional_components": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The set of components to activate on the cluster.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerators": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine accelerator configuration for these instances.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsSchemaV0(),
			},

			"disk_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Disk option config settings.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigSchemaV0(),
			},

			"image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.",
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`.",
			},

			"min_cpu_platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).",
			},

			"num_instances": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**.",
			},

			"preemptibility": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE",
			},

			"instance_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"is_preemptible": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Output only. Specifies that this instance group contains preemptible instances.",
			},

			"managed_group_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups.",
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of the accelerator cards of this type exposed to this instance.",
			},

			"accelerator_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"boot_disk_size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Size in GB of the boot disk (default is 500GB).",
			},

			"boot_disk_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Type of the boot disk (default is \"pd-standard\"). Valid values: \"pd-balanced\" (Persistent Disk Balanced Solid State Drive), \"pd-ssd\" (Persistent Disk Solid State Drive), or \"pd-standard\" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).",
			},

			"num_local_ssds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.",
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_group_manager_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Group Manager for this group.",
			},

			"instance_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The name of the Instance Template used for the Managed Instance Group.",
			},
		},
	}
}

func DataprocWorkflowTemplateParametersSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fields": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Paths to all fields that the parameter replaces. A field is allowed to appear in at most one parameter's list of field paths. A field path is similar in syntax to a google.protobuf.FieldMask. For example, a field path that references the zone field of a workflow template's cluster selector would be specified as `placement.clusterSelector.zone`. Also, field paths can reference fields using the following syntax: * Values in maps can be referenced by key: * labels['key'] * placement.clusterSelector.clusterLabels['key'] * placement.managedCluster.labels['key'] * placement.clusterSelector.clusterLabels['key'] * jobs['step-id'].labels['key'] * Jobs in the jobs list can be referenced by step-id: * jobs['step-id'].hadoopJob.mainJarFileUri * jobs['step-id'].hiveJob.queryFileUri * jobs['step-id'].pySparkJob.mainPythonFileUri * jobs['step-id'].hadoopJob.jarFileUris[0] * jobs['step-id'].hadoopJob.archiveUris[0] * jobs['step-id'].hadoopJob.fileUris[0] * jobs['step-id'].pySparkJob.pythonFileUris[0] * Items in repeated fields can be referenced by a zero-based index: * jobs['step-id'].sparkJob.args[0] * Other examples: * jobs['step-id'].hadoopJob.properties['key'] * jobs['step-id'].hadoopJob.args[0] * jobs['step-id'].hiveJob.scriptVariables['key'] * jobs['step-id'].hadoopJob.mainJarFileUri * placement.clusterSelector.zone It may not be possible to parameterize maps and repeated fields in their entirety since only individual map values and individual items in repeated fields can be referenced. For example, the following field paths are invalid: - placement.clusterSelector.clusterLabels - jobs['step-id'].sparkJob.args",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Brief description of the parameter. Must not exceed 1024 characters.",
			},

			"validation": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. Validation rules to be applied to this parameter's value.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regex": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Validation based on regular expressions.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationRegexSchemaV0(),
			},

			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Validation based on a list of allowed values.",
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationValuesSchemaV0(),
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationRegexSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regexes": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationValuesSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"values": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Required. List of allowed values for the parameter.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ResourceDataprocWorkflowTemplateUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	return tpgresource.TerraformLabelsStateUpgrade(rawState)
}