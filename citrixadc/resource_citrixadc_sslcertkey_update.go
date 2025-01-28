package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSslcertkeyUpdate() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcertkeyUpdateFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"certkey": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"fipskey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"inform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodomaincheck": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passplain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			// New "linkcertkeyname" schema added
			"linkcertkeyname": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    false, // This is not computed as it must be set by the user if needed.
				ForceNew:    true,  // Forces re-creation of the resource if changed
				Description: "The name of the certificate key to link with this certificate key.",
			},
		},
	}
}

func createSslcertkeyUpdateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createSslcertkeyUpdateFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

	// Creating sslcertkey structure
	sslcertkey := ssl.Sslcertkey{
		Cert:          d.Get("cert").(string),
		Certkey:       d.Get("certkey").(string),
		Fipskey:       d.Get("fipskey").(string),
		Inform:        d.Get("inform").(string),
		Key:           d.Get("key").(string),
		Nodomaincheck: true,
		Passplain:     d.Get("passplain").(string),
		Password:      d.Get("password").(bool),
	}

	// Check for linkcertkeyname and add if specified
	if v, ok := d.GetOk("linkcertkeyname"); ok {
		sslcertkey.Linkcertkeyname = v.(string)
	}

	// Performing the update via Nitro API
	err := client.ActOnResource(service.Sslcertkey.Type(), &sslcertkey, "update")
	if err != nil {
		return err
	}

	// Set the ID for the resource (same as certkey name)
	d.SetId(sslcertkeyName)

	return nil
}
