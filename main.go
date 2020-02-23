package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	"golang.org/x/oauth2"
	"k8s.io/client-go/rest"

	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
	"github.com/jetstack/cert-manager/pkg/issuer/acme/dns/util"
)

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		panic("GROUP_NAME must be specified")
	}
	cmd.RunWebhookServer(GroupName, &dnsimpleSolver{})
}

// dnsimpleSolver implements the webhook.Solver interface.
// See: github.com/jetstack/cert-manager/pkg/acme/webhook
type dnsimpleSolver struct {
	client    *dnsimple.Client
	accountID string
}

func (s *dnsimpleSolver) Name() string {
	return "dnsimple"
}

func (s *dnsimpleSolver) Present(ch *v1alpha1.ChallengeRequest) error {
	_, err := s.client.Zones.CreateRecord(s.accountID, strings.TrimRight(ch.ResolvedZone, "."), dnsimple.ZoneRecord{
		Name:    extractRecordName(ch.ResolvedFQDN, ch.ResolvedZone),
		Type:    "TXT",
		TTL:     300,
		Content: ch.Key,
	})
	if err != nil && strings.Contains(err.Error(), "record already exists") {
		return nil
	}
	return err
}

func (s *dnsimpleSolver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	resp, err := s.client.Zones.ListRecords(s.accountID, util.UnFqdn(ch.ResolvedZone), &dnsimple.ZoneRecordListOptions{
		Name: extractRecordName(ch.ResolvedFQDN, ch.ResolvedZone),
	})
	if err != nil {
		return fmt.Errorf("listing zone records: %v", err)
	}
	for _, r := range resp.Data {
		_, err := s.client.Zones.DeleteRecord(s.accountID, util.UnFqdn(ch.ResolvedZone), r.ID)
		if err != nil {
			return fmt.Errorf("deleting record: %v", err)
		}
	}
	return nil
}

func (s *dnsimpleSolver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {
	if s.client == nil {
		s.client = dnsimple.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("DNSIMPLE_ACCESS_TOKEN")})))
	}

	if s.accountID == "" {
		resp, err := dnsimple.Whoami(s.client)
		if err != nil {
			return fmt.Errorf("whoami request: %v", err)
		}
		s.accountID = strconv.FormatInt(resp.Account.ID, 10)
	}

	return nil
}

func extractRecordName(fqdn, domain string) string {
	name := util.UnFqdn(fqdn)
	if idx := strings.Index(name, "."+util.UnFqdn(domain)); idx != -1 {
		return name[:idx]
	}
	return name
}
