# plate-stack

## Operations
All of the operations has now moved to [kot-labs-ops](https://github.com/flynshue/kot-labs-ops#kot-labs-ops) repo

The kot-labs-ops repo contains terraform to create the droplet in DigitalOcean which will be used as a Docker engine

The ansible playbooks will handle installing and configuring docker engine, nginx, and certbot

## Creating New Release
* Bump VERSION file.  Use semantic versioning. See [semver cheatsheet](https://github.com/rstacruz/cheatsheets/blob/master/semver.md)
* Create PR and verify tests pass
* Merge PR, that will kick off GHA workflow which will do a new build and push new image version to AWS ECR
* Go to [kot-labs-ops](https://github.com/flynshue/kot-labs-ops/blob/main/ansible/hosts.yml#L17) and update image tag
* Use instructions [kot-labs-ops](https://github.com/flynshue/kot-labs-ops#running-ansible-playbook-to-update-cod-maps-or-plate-stack) to run ansible-playbook to update plate-stack