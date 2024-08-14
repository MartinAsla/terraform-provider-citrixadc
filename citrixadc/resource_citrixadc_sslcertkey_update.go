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
			"linkcertkeyname": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    false,
				Description: "The name of the certificate key linked to this SSL cert key.",
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
		},
	}
}

func createSslcertkeyUpdateFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In createSslcertkeyUpdateFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcertkeyName := d.Get("certkey").(string)

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

	// Check for linkcertkeyname
	if v, ok := d.GetOk("linkcertkeyname"); ok {
		sslcertkey.Linkcertkeyname = v.(string)
	}

	// Perform the update action
	err := client.ActOnResource(service.Sslcertkey.Type(), &sslcertkey, "update")
	if err != nil {
		return err
	}

	d.SetId(sslcertkeyName)

	return nil
}
