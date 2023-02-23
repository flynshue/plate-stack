<!-- BEGIN_TF_DOCS -->
## Setup
Configure TF Cloud
```
export TF_CLI_CONFIG_FILE="$HOME/.terraformrc.home"

terraform init
```

## Running terraform
```
terraform plan -var do_token=$(lpass show --password blah)
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_digitalocean"></a> [digitalocean](#requirement\_digitalocean) | ~> 2.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_digitalocean"></a> [digitalocean](#provider\_digitalocean) | 2.26.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [digitalocean_droplet.plate-stack](https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/droplet) | resource |
| [digitalocean_record.test-plate-stack](https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/resources/record) | resource |
| [digitalocean_ssh_keys.keys](https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs/data-sources/ssh_keys) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_do_domain"></a> [do\_domain](#input\_do\_domain) | n/a | `string` | `"kot-labs.com"` | no |
| <a name="input_do_region"></a> [do\_region](#input\_do\_region) | n/a | `string` | `"nyc2"` | no |
| <a name="input_do_token"></a> [do\_token](#input\_do\_token) | Set the variable value in *.tfvars file or using -var="do\_token=..." CLI option | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_droplet"></a> [droplet](#output\_droplet) | n/a |
| <a name="output_droplet_record"></a> [droplet\_record](#output\_droplet\_record) | n/a |