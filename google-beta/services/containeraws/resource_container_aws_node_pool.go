// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package containeraws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	containeraws "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/beta"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgdclresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceContainerAwsNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerAwsNodePoolCreate,
		Read:   resourceContainerAwsNodePoolRead,
		Update: resourceContainerAwsNodePoolUpdate,
		Delete: resourceContainerAwsNodePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceContainerAwsNodePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
			tpgdclresource.ResourceContainerAwsNodePoolCustomizeDiffFunc,
			tpgresource.SetAnnotationsDiff,
		),

		Schema: map[string]*schema.Schema{
			"autoscaling": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Autoscaler configuration for this node pool.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolAutoscalingSchema(),
			},

			"cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "The awsCluster for the resource",
			},

			"config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The configuration of the node pool.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigSchema(),
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The location for the resource",
			},

			"max_pods_constraint": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolMaxPodsConstraintSchema(),
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of this resource.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet where the node pool node run.",
			},

			"version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.",
			},

			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
			},

			"kubelet_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "The kubelet configuration for the node pool.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolKubeletConfigSchema(),
			},

			"management": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "The Management configuration for this node pool.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolManagementSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"update_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Optional. Update settings control the speed and disruption of the node pool update.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolUpdateSettingsSchema(),
			},

			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field `effective_annotations` for all of the annotations present on the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The time at which this node pool was created.",
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.",
			},

			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Output only. If set, there are currently changes in flight to the node pool.",
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The lifecycle state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED",
			},

			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. A globally unique identifier for the node pool.",
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. The time at which this node pool was last updated.",
			},
		},
	}
}

func ContainerAwsNodePoolAutoscalingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Maximum number of nodes in the NodePool. Must be >= min_node_count.",
			},

			"min_node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Minimum number of nodes in the NodePool. Must be >= 1 and <= max_node_count.",
			},
		},
	}
}

func ContainerAwsNodePoolConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_encryption": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The ARN of the AWS KMS key used to encrypt node pool configuration.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigConfigEncryptionSchema(),
			},

			"iam_instance_profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the AWS IAM role assigned to nodes in the pool.",
			},

			"autoscaling_metrics_collection": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optional. Configuration related to CloudWatch metrics collection on the Auto Scaling group of the node pool. When unspecified, metrics collection is disabled.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigAutoscalingMetricsCollectionSchema(),
			},

			"image_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "The OS image type to use on node pool instances.",
			},

			"instance_placement": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Details of placement information for an instance.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigInstancePlacementSchema(),
			},

			"instance_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.",
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Optional. The initial labels assigned to nodes of this node pool. An object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"proxy_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Proxy configuration for outbound HTTP(S) traffic.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigProxyConfigSchema(),
			},

			"root_volume": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Optional. Template for the root volume provisioned for node pool nodes. Volumes will be provisioned in the availability zone assigned to the node pool subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigRootVolumeSchema(),
			},

			"security_group_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optional. The IDs of additional security groups to add to nodes in this pool. The manager will automatically create security groups with minimum rules needed for a functioning cluster.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"spot_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. When specified, the node pool will provision Spot instances from the set of spot_config.instance_types. This field is mutually exclusive with `instance_type`",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigSpotConfigSchema(),
			},

			"ssh_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Optional. The SSH configuration.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolConfigSshConfigSchema(),
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Optional. Key/value metadata to assign to each underlying AWS resource. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"taints": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The initial taints assigned to nodes of this node pool.",
				Elem:        ContainerAwsNodePoolConfigTaintsSchema(),
			},
		},
	}
}

func ContainerAwsNodePoolConfigConfigEncryptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kms_key_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ARN of the AWS KMS key used to encrypt node pool configuration.",
			},
		},
	}
}

func ContainerAwsNodePoolConfigAutoscalingMetricsCollectionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"granularity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The frequency at which EC2 Auto Scaling sends aggregated data to AWS CloudWatch. The only valid value is \"1Minute\".",
			},

			"metrics": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The metrics to enable. For a list of valid metrics, see https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html. If you specify granularity and don't specify any metrics, all metrics are enabled.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ContainerAwsNodePoolConfigInstancePlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tenancy": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CaseDiffSuppress,
				Description:      "The tenancy for the instance. Possible values: TENANCY_UNSPECIFIED, DEFAULT, DEDICATED, HOST",
			},
		},
	}
}

func ContainerAwsNodePoolConfigProxyConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_arn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.",
			},

			"secret_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.",
			},
		},
	}
}

func ContainerAwsNodePoolConfigRootVolumeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iops": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.",
			},

			"kms_key_arn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.",
			},

			"size_gib": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.",
			},

			"throughput": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The throughput to provision for the volume, in MiB/s. Only valid if the volume type is GP3. If volume type is gp3 and throughput is not specified, the throughput will defaults to 125.",
			},

			"volume_type": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CaseDiffSuppress,
				Description:      "Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3",
			},
		},
	}
}

func ContainerAwsNodePoolConfigSpotConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_types": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "List of AWS EC2 instance types for creating a spot node pool's nodes. The specified instance types must have the same number of CPUs and memory. You can use the Amazon EC2 Instance Selector tool (https://github.com/aws/amazon-ec2-instance-selector) to choose instance types with matching CPU and memory",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ContainerAwsNodePoolConfigSshConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ec2_key_pair": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the EC2 key pair used to login into cluster machines.",
			},
		},
	}
}

func ContainerAwsNodePoolConfigTaintsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CaseDiffSuppress,
				Description:      "The taint effect. Possible values: EFFECT_UNSPECIFIED, NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE",
			},

			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Key for the taint.",
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Value for the taint.",
			},
		},
	}
}

func ContainerAwsNodePoolMaxPodsConstraintSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_pods_per_node": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The maximum number of pods to schedule on a single node.",
			},
		},
	}
}

func ContainerAwsNodePoolKubeletConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_cfs_quota": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether or not to enable CPU CFS quota. Defaults to true.",
			},

			"cpu_cfs_quota_period": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The CPU CFS quota period to use for the node. Defaults to \"100ms\".",
			},

			"cpu_manager_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: "The CpuManagerPolicy to use for the node. Defaults to \"none\".",
			},

			"pod_pids_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional. The maximum number of PIDs in each pod running on the node. The limit scales automatically based on underlying machine size if left unset.",
			},
		},
	}
}

func ContainerAwsNodePoolManagementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_repair": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				Description: "Optional. Whether or not the nodes will be automatically repaired.",
			},
		},
	}
}

func ContainerAwsNodePoolUpdateSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"surge_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Optional. Settings for surge update.",
				MaxItems:    1,
				Elem:        ContainerAwsNodePoolUpdateSettingsSurgeSettingsSchema(),
			},
		},
	}
}

func ContainerAwsNodePoolUpdateSettingsSurgeSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_surge": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The maximum number of nodes that can be created beyond the current size of the node pool during the update process.",
			},

			"max_unavailable": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				Description: "Optional. The maximum number of nodes that can be simultaneously unavailable during the update process. A node is considered unavailable if its status is not Ready.",
			},
		},
	}
}

func resourceContainerAwsNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &containeraws.NodePool{
		Autoscaling:       expandContainerAwsNodePoolAutoscaling(d.Get("autoscaling")),
		Cluster:           dcl.String(d.Get("cluster").(string)),
		Config:            expandContainerAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandContainerAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       tpgresource.CheckStringMap(d.Get("effective_annotations")),
		KubeletConfig:     expandContainerAwsNodePoolKubeletConfig(d.Get("kubelet_config")),
		Management:        expandContainerAwsNodePoolManagement(d.Get("management")),
		Project:           dcl.String(project),
		UpdateSettings:    expandContainerAwsNodePoolUpdateSettings(d.Get("update_settings")),
	}

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	directive := tpgdclresource.CreateDirective
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLContainerAwsClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyNodePool(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating NodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NodePool %q: %#v", d.Id(), res)

	return resourceContainerAwsNodePoolRead(d, meta)
}

func resourceContainerAwsNodePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &containeraws.NodePool{
		Autoscaling:       expandContainerAwsNodePoolAutoscaling(d.Get("autoscaling")),
		Cluster:           dcl.String(d.Get("cluster").(string)),
		Config:            expandContainerAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandContainerAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       tpgresource.CheckStringMap(d.Get("effective_annotations")),
		KubeletConfig:     expandContainerAwsNodePoolKubeletConfig(d.Get("kubelet_config")),
		Management:        expandContainerAwsNodePoolManagement(d.Get("management")),
		Project:           dcl.String(project),
		UpdateSettings:    expandContainerAwsNodePoolUpdateSettings(d.Get("update_settings")),
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLContainerAwsClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetNodePool(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("ContainerAwsNodePool %q", d.Id())
		return tpgdclresource.HandleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("autoscaling", flattenContainerAwsNodePoolAutoscaling(res.Autoscaling)); err != nil {
		return fmt.Errorf("error setting autoscaling in state: %s", err)
	}
	if err = d.Set("cluster", res.Cluster); err != nil {
		return fmt.Errorf("error setting cluster in state: %s", err)
	}
	if err = d.Set("config", flattenContainerAwsNodePoolConfig(res.Config)); err != nil {
		return fmt.Errorf("error setting config in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("max_pods_constraint", flattenContainerAwsNodePoolMaxPodsConstraint(res.MaxPodsConstraint)); err != nil {
		return fmt.Errorf("error setting max_pods_constraint in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("subnet_id", res.SubnetId); err != nil {
		return fmt.Errorf("error setting subnet_id in state: %s", err)
	}
	if err = d.Set("version", res.Version); err != nil {
		return fmt.Errorf("error setting version in state: %s", err)
	}
	if err = d.Set("effective_annotations", res.Annotations); err != nil {
		return fmt.Errorf("error setting effective_annotations in state: %s", err)
	}
	if err = d.Set("kubelet_config", flattenContainerAwsNodePoolKubeletConfig(res.KubeletConfig)); err != nil {
		return fmt.Errorf("error setting kubelet_config in state: %s", err)
	}
	if err = d.Set("management", tpgresource.FlattenContainerAwsNodePoolManagement(res.Management, d, config)); err != nil {
		return fmt.Errorf("error setting management in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("update_settings", flattenContainerAwsNodePoolUpdateSettings(res.UpdateSettings)); err != nil {
		return fmt.Errorf("error setting update_settings in state: %s", err)
	}
	if err = d.Set("annotations", flattenContainerAwsNodePoolAnnotations(res.Annotations, d)); err != nil {
		return fmt.Errorf("error setting annotations in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("reconciling", res.Reconciling); err != nil {
		return fmt.Errorf("error setting reconciling in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("uid", res.Uid); err != nil {
		return fmt.Errorf("error setting uid in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceContainerAwsNodePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &containeraws.NodePool{
		Autoscaling:       expandContainerAwsNodePoolAutoscaling(d.Get("autoscaling")),
		Cluster:           dcl.String(d.Get("cluster").(string)),
		Config:            expandContainerAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandContainerAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       tpgresource.CheckStringMap(d.Get("effective_annotations")),
		KubeletConfig:     expandContainerAwsNodePoolKubeletConfig(d.Get("kubelet_config")),
		Management:        expandContainerAwsNodePoolManagement(d.Get("management")),
		Project:           dcl.String(project),
		UpdateSettings:    expandContainerAwsNodePoolUpdateSettings(d.Get("update_settings")),
	}
	directive := tpgdclresource.UpdateDirective
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLContainerAwsClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutUpdate))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyNodePool(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating NodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NodePool %q: %#v", d.Id(), res)

	return resourceContainerAwsNodePoolRead(d, meta)
}

func resourceContainerAwsNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	obj := &containeraws.NodePool{
		Autoscaling:       expandContainerAwsNodePoolAutoscaling(d.Get("autoscaling")),
		Cluster:           dcl.String(d.Get("cluster").(string)),
		Config:            expandContainerAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandContainerAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       tpgresource.CheckStringMap(d.Get("effective_annotations")),
		KubeletConfig:     expandContainerAwsNodePoolKubeletConfig(d.Get("kubelet_config")),
		Management:        expandContainerAwsNodePoolManagement(d.Get("management")),
		Project:           dcl.String(project),
		UpdateSettings:    expandContainerAwsNodePoolUpdateSettings(d.Get("update_settings")),
	}

	log.Printf("[DEBUG] Deleting NodePool %q", d.Id())
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := transport_tpg.NewDCLContainerAwsClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := tpgresource.ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	if err := client.DeleteNodePool(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting NodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting NodePool %q", d.Id())
	return nil
}

func resourceContainerAwsNodePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/awsClusters/(?P<cluster>[^/]+)/awsNodePools/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<cluster>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<cluster>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/awsClusters/{{cluster}}/awsNodePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandContainerAwsNodePoolAutoscaling(o interface{}) *containeraws.NodePoolAutoscaling {
	if o == nil {
		return containeraws.EmptyNodePoolAutoscaling
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolAutoscaling
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolAutoscaling{
		MaxNodeCount: dcl.Int64(int64(obj["max_node_count"].(int))),
		MinNodeCount: dcl.Int64(int64(obj["min_node_count"].(int))),
	}
}

func flattenContainerAwsNodePoolAutoscaling(obj *containeraws.NodePoolAutoscaling) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_node_count": obj.MaxNodeCount,
		"min_node_count": obj.MinNodeCount,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfig(o interface{}) *containeraws.NodePoolConfig {
	if o == nil {
		return containeraws.EmptyNodePoolConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfig{
		ConfigEncryption:             expandContainerAwsNodePoolConfigConfigEncryption(obj["config_encryption"]),
		IamInstanceProfile:           dcl.String(obj["iam_instance_profile"].(string)),
		AutoscalingMetricsCollection: expandContainerAwsNodePoolConfigAutoscalingMetricsCollection(obj["autoscaling_metrics_collection"]),
		ImageType:                    dcl.StringOrNil(obj["image_type"].(string)),
		InstancePlacement:            expandContainerAwsNodePoolConfigInstancePlacement(obj["instance_placement"]),
		InstanceType:                 dcl.StringOrNil(obj["instance_type"].(string)),
		Labels:                       tpgresource.CheckStringMap(obj["labels"]),
		ProxyConfig:                  expandContainerAwsNodePoolConfigProxyConfig(obj["proxy_config"]),
		RootVolume:                   expandContainerAwsNodePoolConfigRootVolume(obj["root_volume"]),
		SecurityGroupIds:             tpgdclresource.ExpandStringArray(obj["security_group_ids"]),
		SpotConfig:                   expandContainerAwsNodePoolConfigSpotConfig(obj["spot_config"]),
		SshConfig:                    expandContainerAwsNodePoolConfigSshConfig(obj["ssh_config"]),
		Tags:                         tpgresource.CheckStringMap(obj["tags"]),
		Taints:                       expandContainerAwsNodePoolConfigTaintsArray(obj["taints"]),
	}
}

func flattenContainerAwsNodePoolConfig(obj *containeraws.NodePoolConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"config_encryption":              flattenContainerAwsNodePoolConfigConfigEncryption(obj.ConfigEncryption),
		"iam_instance_profile":           obj.IamInstanceProfile,
		"autoscaling_metrics_collection": flattenContainerAwsNodePoolConfigAutoscalingMetricsCollection(obj.AutoscalingMetricsCollection),
		"image_type":                     obj.ImageType,
		"instance_placement":             flattenContainerAwsNodePoolConfigInstancePlacement(obj.InstancePlacement),
		"instance_type":                  obj.InstanceType,
		"labels":                         obj.Labels,
		"proxy_config":                   flattenContainerAwsNodePoolConfigProxyConfig(obj.ProxyConfig),
		"root_volume":                    flattenContainerAwsNodePoolConfigRootVolume(obj.RootVolume),
		"security_group_ids":             obj.SecurityGroupIds,
		"spot_config":                    flattenContainerAwsNodePoolConfigSpotConfig(obj.SpotConfig),
		"ssh_config":                     flattenContainerAwsNodePoolConfigSshConfig(obj.SshConfig),
		"tags":                           obj.Tags,
		"taints":                         flattenContainerAwsNodePoolConfigTaintsArray(obj.Taints),
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigConfigEncryption(o interface{}) *containeraws.NodePoolConfigConfigEncryption {
	if o == nil {
		return containeraws.EmptyNodePoolConfigConfigEncryption
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfigConfigEncryption
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigConfigEncryption{
		KmsKeyArn: dcl.String(obj["kms_key_arn"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigConfigEncryption(obj *containeraws.NodePoolConfigConfigEncryption) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kms_key_arn": obj.KmsKeyArn,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigAutoscalingMetricsCollection(o interface{}) *containeraws.NodePoolConfigAutoscalingMetricsCollection {
	if o == nil {
		return containeraws.EmptyNodePoolConfigAutoscalingMetricsCollection
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfigAutoscalingMetricsCollection
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigAutoscalingMetricsCollection{
		Granularity: dcl.String(obj["granularity"].(string)),
		Metrics:     tpgdclresource.ExpandStringArray(obj["metrics"]),
	}
}

func flattenContainerAwsNodePoolConfigAutoscalingMetricsCollection(obj *containeraws.NodePoolConfigAutoscalingMetricsCollection) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"granularity": obj.Granularity,
		"metrics":     obj.Metrics,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigInstancePlacement(o interface{}) *containeraws.NodePoolConfigInstancePlacement {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigInstancePlacement{
		Tenancy: containeraws.NodePoolConfigInstancePlacementTenancyEnumRef(obj["tenancy"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigInstancePlacement(obj *containeraws.NodePoolConfigInstancePlacement) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"tenancy": obj.Tenancy,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigProxyConfig(o interface{}) *containeraws.NodePoolConfigProxyConfig {
	if o == nil {
		return containeraws.EmptyNodePoolConfigProxyConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfigProxyConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigProxyConfig{
		SecretArn:     dcl.String(obj["secret_arn"].(string)),
		SecretVersion: dcl.String(obj["secret_version"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigProxyConfig(obj *containeraws.NodePoolConfigProxyConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"secret_arn":     obj.SecretArn,
		"secret_version": obj.SecretVersion,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigRootVolume(o interface{}) *containeraws.NodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigRootVolume{
		Iops:       dcl.Int64OrNil(int64(obj["iops"].(int))),
		KmsKeyArn:  dcl.String(obj["kms_key_arn"].(string)),
		SizeGib:    dcl.Int64OrNil(int64(obj["size_gib"].(int))),
		Throughput: dcl.Int64OrNil(int64(obj["throughput"].(int))),
		VolumeType: containeraws.NodePoolConfigRootVolumeVolumeTypeEnumRef(obj["volume_type"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigRootVolume(obj *containeraws.NodePoolConfigRootVolume) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"iops":        obj.Iops,
		"kms_key_arn": obj.KmsKeyArn,
		"size_gib":    obj.SizeGib,
		"throughput":  obj.Throughput,
		"volume_type": obj.VolumeType,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigSpotConfig(o interface{}) *containeraws.NodePoolConfigSpotConfig {
	if o == nil {
		return containeraws.EmptyNodePoolConfigSpotConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfigSpotConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigSpotConfig{
		InstanceTypes: tpgdclresource.ExpandStringArray(obj["instance_types"]),
	}
}

func flattenContainerAwsNodePoolConfigSpotConfig(obj *containeraws.NodePoolConfigSpotConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"instance_types": obj.InstanceTypes,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolConfigSshConfig(o interface{}) *containeraws.NodePoolConfigSshConfig {
	if o == nil {
		return containeraws.EmptyNodePoolConfigSshConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolConfigSshConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolConfigSshConfig{
		Ec2KeyPair: dcl.String(obj["ec2_key_pair"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigSshConfig(obj *containeraws.NodePoolConfigSshConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ec2_key_pair": obj.Ec2KeyPair,
	}

	return []interface{}{transformed}

}
func expandContainerAwsNodePoolConfigTaintsArray(o interface{}) []containeraws.NodePoolConfigTaints {
	if o == nil {
		return make([]containeraws.NodePoolConfigTaints, 0)
	}

	objs := o.([]interface{})
	if len(objs) == 0 || objs[0] == nil {
		return make([]containeraws.NodePoolConfigTaints, 0)
	}

	items := make([]containeraws.NodePoolConfigTaints, 0, len(objs))
	for _, item := range objs {
		i := expandContainerAwsNodePoolConfigTaints(item)
		items = append(items, *i)
	}

	return items
}

func expandContainerAwsNodePoolConfigTaints(o interface{}) *containeraws.NodePoolConfigTaints {
	if o == nil {
		return containeraws.EmptyNodePoolConfigTaints
	}

	obj := o.(map[string]interface{})
	return &containeraws.NodePoolConfigTaints{
		Effect: containeraws.NodePoolConfigTaintsEffectEnumRef(obj["effect"].(string)),
		Key:    dcl.String(obj["key"].(string)),
		Value:  dcl.String(obj["value"].(string)),
	}
}

func flattenContainerAwsNodePoolConfigTaintsArray(objs []containeraws.NodePoolConfigTaints) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenContainerAwsNodePoolConfigTaints(&item)
		items = append(items, i)
	}

	return items
}

func flattenContainerAwsNodePoolConfigTaints(obj *containeraws.NodePoolConfigTaints) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"effect": obj.Effect,
		"key":    obj.Key,
		"value":  obj.Value,
	}

	return transformed

}

func expandContainerAwsNodePoolMaxPodsConstraint(o interface{}) *containeraws.NodePoolMaxPodsConstraint {
	if o == nil {
		return containeraws.EmptyNodePoolMaxPodsConstraint
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return containeraws.EmptyNodePoolMaxPodsConstraint
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64(int64(obj["max_pods_per_node"].(int))),
	}
}

func flattenContainerAwsNodePoolMaxPodsConstraint(obj *containeraws.NodePoolMaxPodsConstraint) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_pods_per_node": obj.MaxPodsPerNode,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolKubeletConfig(o interface{}) *containeraws.NodePoolKubeletConfig {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolKubeletConfig{
		CpuCfsQuota:       dcl.Bool(obj["cpu_cfs_quota"].(bool)),
		CpuCfsQuotaPeriod: dcl.String(obj["cpu_cfs_quota_period"].(string)),
		CpuManagerPolicy:  containeraws.NodePoolKubeletConfigCpuManagerPolicyEnumRef(obj["cpu_manager_policy"].(string)),
		PodPidsLimit:      dcl.Int64(int64(obj["pod_pids_limit"].(int))),
	}
}

func flattenContainerAwsNodePoolKubeletConfig(obj *containeraws.NodePoolKubeletConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cpu_cfs_quota":        obj.CpuCfsQuota,
		"cpu_cfs_quota_period": obj.CpuCfsQuotaPeriod,
		"cpu_manager_policy":   obj.CpuManagerPolicy,
		"pod_pids_limit":       obj.PodPidsLimit,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolManagement(o interface{}) *containeraws.NodePoolManagement {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolManagement{
		AutoRepair: dcl.Bool(obj["auto_repair"].(bool)),
	}
}

func flattenContainerAwsNodePoolManagement(obj *containeraws.NodePoolManagement) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"auto_repair": obj.AutoRepair,
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolUpdateSettings(o interface{}) *containeraws.NodePoolUpdateSettings {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolUpdateSettings{
		SurgeSettings: expandContainerAwsNodePoolUpdateSettingsSurgeSettings(obj["surge_settings"]),
	}
}

func flattenContainerAwsNodePoolUpdateSettings(obj *containeraws.NodePoolUpdateSettings) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"surge_settings": flattenContainerAwsNodePoolUpdateSettingsSurgeSettings(obj.SurgeSettings),
	}

	return []interface{}{transformed}

}

func expandContainerAwsNodePoolUpdateSettingsSurgeSettings(o interface{}) *containeraws.NodePoolUpdateSettingsSurgeSettings {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &containeraws.NodePoolUpdateSettingsSurgeSettings{
		MaxSurge:       dcl.Int64OrNil(int64(obj["max_surge"].(int))),
		MaxUnavailable: dcl.Int64OrNil(int64(obj["max_unavailable"].(int))),
	}
}

func flattenContainerAwsNodePoolUpdateSettingsSurgeSettings(obj *containeraws.NodePoolUpdateSettingsSurgeSettings) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_surge":       obj.MaxSurge,
		"max_unavailable": obj.MaxUnavailable,
	}

	return []interface{}{transformed}

}

func flattenContainerAwsNodePoolAnnotations(v map[string]string, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}

	transformed := make(map[string]interface{})
	if l, ok := d.Get("annotations").(map[string]interface{}); ok {
		for k := range l {
			transformed[k] = v[k]
		}
	}

	return transformed
}
