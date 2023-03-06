output "droplet" {
  value = digitalocean_droplet.plate-stack
}

output "droplet_record" {
    value = digitalocean_record.test-plate-stack
}