# kube-demo-site

Kubernetes Demo Site

## Pulumi

You will need several env vars from the Prometheus Team secrets storage in order to apply changes.

* `export AWS_ACCESS_KEY_ID=XXX`
* `export AWS_SECRET_ACCESS_KEY=XXX`
* `export AWS_REGION=ams3`
* `export PULUMI_CONFIG_PASSPHRASE=XXX`
* `PULUMI_BACKEND_URL='s3://prometheus-ci/kube-demo-site/do-k8s?endpoint=ams3.digitaloceanspaces.com&s3ForcePathStyle=true'`

After this you can `cd do-k8s` and run Pulumi commands.

```console
pulumi login
pulumi up
```
