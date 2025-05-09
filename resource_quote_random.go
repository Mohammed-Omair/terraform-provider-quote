package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceQuoteRandom() *schema.Resource {
	return &schema.Resource{
		Create: resourceQuoteRandomCreate,
		Read:   resourceQuoteRandomRead,
		Delete: resourceQuoteRandomDelete,
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"author": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceQuoteRandomCreate(d *schema.ResourceData, meta interface{}) error {
	// HTTP client with disabled TLS verification
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://api.quotable.io/random")
	if err != nil {
		return fmt.Errorf("error fetching quote: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().UnixNano(), 10))
	d.Set("content", result.Content)
	d.Set("author", result.Author)

	return nil
}

func resourceQuoteRandomRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceQuoteRandomDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
