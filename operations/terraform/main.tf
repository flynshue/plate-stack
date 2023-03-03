# locals {
#    image_sizes = { for image in data.digitalocean_images.ubuntu_images.images : image.slug => "${image.size_gigabytes}gb" }
# }

# data "digitalocean_images" "ubuntu_images" {
#     filter {
#       key = "distribution"
#       values = ["Ubuntu"]
#     }
#     filter {
#         key = "slug"
#         values = ["ubuntu"]
#         match_by = "substring"
#     }
# }

# output "images" {
#   value = local.image_sizes
# }

data "digitalocean_ssh_keys" "keys" {
  sort {
    key       = "name"
    direction = "asc"
  }
}

resource "digitalocean_droplet" "plate-stack" {
  image = "ubuntu-22-10-x64"
  name = "test-plate-stack-${var.do_region}-01"
  region = var.do_region
  size = "s-1vcpu-1gb"
  tags = ["plate-stack"]
  ssh_keys = data.digitalocean_ssh_keys.keys.ssh_keys[*].id
}

resource "digitalocean_record" "test-plate-stack" {
    domain = var.do_domain
    name = "test-plate-stack"
    type = "A"
    value = digitalocean_droplet.plate-stack.ipv4_address
    ttl = 60
}