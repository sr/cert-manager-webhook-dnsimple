module github.com/sr/cert-manager-webhook-dnsimple

go 1.13

require (
	github.com/dnsimple/dnsimple-go v0.31.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jetstack/cert-manager v0.13.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	k8s.io/client-go v0.17.0
)

replace github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190203023257-5858425f7550
