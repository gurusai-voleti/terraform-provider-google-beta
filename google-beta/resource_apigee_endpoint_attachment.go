// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApigeeEndpointAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEndpointAttachmentCreate,
		Read:   resourceApigeeEndpointAttachmentRead,
		Delete: resourceApigeeEndpointAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEndpointAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"endpoint_attachment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the endpoint attachment.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location of the endpoint attachment.`,
			},
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee Organization associated with the Apigee instance,
in the format 'organizations/{{org_name}}'.`,
			},
			"service_attachment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Format: projects/*/regions/*/serviceAttachments/*`,
			},
			"connection_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the endpoint attachment connection to the service attachment.`,
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Name of the Endpoint Attachment in the following format:
organizations/{organization}/endpointAttachments/{endpointAttachment}.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEndpointAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	locationProp, err := expandApigeeEndpointAttachmentLocation(d.Get("location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location"); !isEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	serviceAttachmentProp, err := expandApigeeEndpointAttachmentServiceAttachment(d.Get("service_attachment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_attachment"); !isEmptyValue(reflect.ValueOf(serviceAttachmentProp)) && (ok || !reflect.DeepEqual(v, serviceAttachmentProp)) {
		obj["serviceAttachment"] = serviceAttachmentProp
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/endpointAttachments?endpointAttachmentId={{endpoint_attachment_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EndpointAttachment: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EndpointAttachment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = apigeeOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating EndpointAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create EndpointAttachment: %s", err)
	}

	if err := d.Set("name", flattenApigeeEndpointAttachmentName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EndpointAttachment %q: %#v", d.Id(), res)

	return resourceApigeeEndpointAttachmentRead(d, meta)
}

func resourceApigeeEndpointAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeEndpointAttachment %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeEndpointAttachmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("location", flattenApigeeEndpointAttachmentLocation(res["location"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("host", flattenApigeeEndpointAttachmentHost(res["host"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("service_attachment", flattenApigeeEndpointAttachmentServiceAttachment(res["serviceAttachment"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("connection_state", flattenApigeeEndpointAttachmentConnectionState(res["connectionState"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}

	return nil
}

func resourceApigeeEndpointAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{ApigeeBasePath}}{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EndpointAttachment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EndpointAttachment")
	}

	err = apigeeOperationWaitTime(
		config, res, "Deleting EndpointAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EndpointAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEndpointAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	nameParts := strings.Split(d.Get("name").(string), "/")
	if len(nameParts) == 4 {
		// `organizations/{{org_name}}/endpointAttachment/{{endpoint_attachment_id}}`
		orgId := fmt.Sprintf("organizations/%s", nameParts[1])
		if err := d.Set("org_id", orgId); err != nil {
			return nil, fmt.Errorf("Error setting org_id: %s", err)
		}
		if err := d.Set("endpoint_attachment_id", nameParts[3]); err != nil {
			return nil, fmt.Errorf("Error setting endpoint_attachment_id: %s", err)
		}
	} else {
		return nil, fmt.Errorf(
			"Saw %s when the name is expected to have shape %s",
			d.Get("name"),
			"organizations/{{org_name}}/environments/{{name}}")
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEndpointAttachmentName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEndpointAttachmentLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEndpointAttachmentHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEndpointAttachmentServiceAttachment(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEndpointAttachmentConnectionState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApigeeEndpointAttachmentLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEndpointAttachmentServiceAttachment(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
