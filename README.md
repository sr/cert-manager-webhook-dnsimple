# cert-manager-webhook-dnsimple

[DNSimple](https://dnsimple.com) provider for [cert-manager](https://cert-manager.io/).

## Running the test suite


```bash
$ hack/hack/fetch-test-binaries.sh
$ TEST_ZONE_NAME="example.com." DNSIMPLE_ACCESS_TOKEN="<TOKEN>" go test .
```
