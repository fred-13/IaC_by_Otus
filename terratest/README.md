    $ go test

```
    ...

TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66: {
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:   "load_balancer_public_ip": {
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:     "sensitive": false,
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:     "type": [
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:       "list",
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:       "string"
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:     ],
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:     "value": [
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:       "62.84.113.22"
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:     ]
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66:   }
TestTerraformHelloWorldExample 2021-09-26T23:03:45+03:00 logger.go:66: }
TestTerraformHelloWorldExample 2021-09-26T23:05:25+03:00 http_helper.go:32: Making an HTTP GET call to URL http://[62.84.113.22]

    ...

TestTerraformHelloWorldExample 2021-09-26T23:05:25+03:00 http_helper.go:32: Making an HTTP GET call to URL http://[62.84.113.22]

    ...

TestTerraformHelloWorldExample 2021-09-26T23:05:52+03:00 logger.go:66: Destroy complete! Resources: 9 destroyed.
TestTerraformHelloWorldExample 2021-09-26T23:05:52+03:00 logger.go:66: 
PASS
ok  	github.com/test/of_test	181.174s
```
