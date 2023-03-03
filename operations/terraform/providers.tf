terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }

  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "kot-labs"

    workspaces {
      name = "plate-stack"
    }
  }

}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = var.do_token
}