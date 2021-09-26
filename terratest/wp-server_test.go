package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
	})

	defer terraform.Destroy(t, terraformOptions)
		  terraform.InitAndApply(t, terraformOptions)

	output := terraform.OutputForKeys(t, terraformOptions, []string{"load_balancer_public_ip"})
	ip := output["load_balancer_public_ip"]
	time.Sleep(100 * time.Second)
	status_code, _ := http_helper.HttpGet(t, fmt.Sprintf("http://%s", ip), &tls.Config{})
	assert.Equal(t, 200, status_code)
}
