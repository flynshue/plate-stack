# plate-stack

## Nginx configuration
Here are the current sites in nginx

```
ll /etc/nginx/sites-available/
total 24
drwxr-xr-x 2 root root 4096 Nov 16 06:21 ./
drwxr-xr-x 8 root root 4096 Feb 16 06:20 ../
-rw-r--r-- 1 root root  694 Nov 15  2020 codMaps
-rw-r--r-- 1 root root 2416 Mar 26  2020 default
-rw-r--r-- 1 root root  754 Nov 15  2020 lenslocked.com
-rw-r--r-- 1 root root  715 Aug 18  2022 snappass
```


* Create an ansible playbook to manage nginx
* Create an ansible playbook to manage letsencrypt

## Terraform stuff
Creating webserver for to host docker container, please see [operations/terraform](operations/terraform/README.md#setup)

Once the server is up, you can run the ansible playbook to bootstrap it with docker, nginx, certbot


## Ansible playbook stuff
Setup python environment
```bash
cd operations

python3 -m venv venv

venv/bin/pip install -r requirement.txt

venv/bin/ansible-galaxy -fr ansible-galaxy.yml
```

To generate cert against letsencrypt staging
```
venv/bin/ansible-playbook -i hosts.yml setup.yml --private-key ~/.ssh/id_rsa -e email_address=flynshue@gmail.com -e cert_type=staging
```

To generate cert against letsencrypt prod
```
venv/bin/ansible-playbook -i hosts.yml setup.yml --private-key ~/.ssh/id_rsa -e email_address=flynshue@gmail.com -e cert_type=prod
```

**Note: If you ran playbook to generate a staging cert, you'll have to delete the stage cert before generating the prod cert**
```
certbot certificates
Saving debug log to /var/log/letsencrypt/letsencrypt.log

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Found the following certs:
  Certificate Name: test-plate-stack.kot-labs.com
    Serial Number: fab59fdd9ddd19499ccc21b6dffabca3b5d1
    Key Type: ECDSA
    Domains: test-plate-stack.kot-labs.com
    Expiry Date: 2023-06-04 19:51:05+00:00 (INVALID: TEST_CERT)
    Certificate Path: /etc/letsencrypt/live/test-plate-stack.kot-labs.com/fullchain.pem
    Private Key Path: /etc/letsencrypt/live/test-plate-stack.kot-labs.com/privkey.pem
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

```

```
openssl x509 -text -noout -in /etc/letsencrypt/live/test-plate-stack.kot-labs.com/cert.pem 
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            fa:b5:9f:dd:9d:dd:19:49:9c:cc:21:b6:df:fa:bc:a3:b5:d1
        Signature Algorithm: ecdsa-with-SHA384
        Issuer: C = US, O = (STAGING) Let's Encrypt, CN = (STAGING) Ersatz Edamame E1
        Validity
            Not Before: Mar  6 19:51:06 2023 GMT
            Not After : Jun  4 19:51:05 2023 GMT
        Subject: CN = test-plate-stack.kot-labs.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:c1:ba:59:9c:04:f4:9f:51:0d:e1:00:a9:45:e4:
                    db:59:d7:6d:90:bf:35:41:70:61:f9:05:72:8b:e9:
                    23:d8:55:14:37:9d:ec:cb:42:34:f7:62:41:fe:8e:
                    2e:50:fd:64:cb:f7:12:77:d6:bc:01:04:38:a1:10:
                    b0:6f:4e:de:3c
                ASN1 OID: prime256v1
                NIST CURVE: P-256
```

Deleting staging cert
```
certbot delete --cert-name test-plate-stack.kot-labs.com
```

Not sure how ansible is going to work with the automating nginx and certbot configs though.