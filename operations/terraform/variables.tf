# Set the variable value in *.tfvars file
# or using -var="do_token=..." CLI option
variable "do_token" {}

variable "do_domain" {
  type = string
  default = "kot-labs.com"
}

variable "do_region" {
  type = string
  default = "nyc3"
}