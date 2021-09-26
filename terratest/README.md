    $ go test

```
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 retry.go:91: terraform [init -upgrade=false]
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: Running command terraform with args [init -upgrade=false]
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: Initializing the backend...
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: Initializing provider plugins...
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: - Reusing previous version of yandex-cloud/yandex from the dependency lock file
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: - Using previously-installed yandex-cloud/yandex v0.64.0
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: Terraform has been successfully initialized!
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: You may now begin working with Terraform. Try running "terraform plan" to see
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: any changes that are required for your infrastructure. All Terraform commands
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: should now work.
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: If you ever set or change modules or backend configuration for Terraform,
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: rerun this command to reinitialize your working directory. If you forget, other
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: commands will detect it and remind you to do so if necessary.
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 retry.go:91: terraform [apply -input=false -auto-approve -lock=false]
TestTerraformHelloWorldExample 2021-09-26T18:44:45+03:00 logger.go:66: Running command terraform with args [apply -input=false -auto-approve -lock=false]
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: Terraform used the selected providers to generate the following execution
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: plan. Resource actions are indicated with the following symbols:
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:   + create
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: Terraform will perform the following actions:
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:   # yandex_compute_instance.wp-app-1 will be created
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:   + resource "yandex_compute_instance" "wp-app-1" {
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:       + created_at                = (known after apply)
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:       + folder_id                 = (known after apply)
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:       + fqdn                      = (known after apply)
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:       + hostname                  = (known after apply)
TestTerraformHelloWorldExample 2021-09-26T18:44:46+03:00 logger.go:66:       + id                        = (known after apply)

    ...

TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: Apply complete! Resources: 9 added, 0 changed, 0 destroyed.
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: Outputs:
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: 
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: database_host_fqdn = tolist([
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66:   "rc1b-w2d0ib9j3yk3d3yt.mdb.yandexcloud.net",
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66:   "rc1c-5y6uqgn13d099j6h.mdb.yandexcloud.net",
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: ])
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: load_balancer_public_ip = tolist([
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66:   "62.84.121.138",
TestTerraformHelloWorldExample 2021-09-26T18:52:37+03:00 logger.go:66: ])

    ...

TestTerraformHelloWorldExample 2021-09-26T18:54:17+03:00 http_helper.go:32: Making an HTTP GET call to URL http://[62.84.121.138]

    ...

TestTerraformHelloWorldExample 2021-09-26T18:42:05+03:00 logger.go:66: Destroy complete! Resources: 9 destroyed.
TestTerraformHelloWorldExample 2021-09-26T18:42:05+03:00 logger.go:66: 
--- FAIL: TestTerraformHelloWorldExample (177.20s)
    http_helper.go:24: Get "http://[62.84.118.33]": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
FAIL
exit status 1
FAIL	github.com/test/of_test	177.202s
```

### Тест провален потому что на ноды wp-app-1 и wp-app-2 не установлены web сервера.
