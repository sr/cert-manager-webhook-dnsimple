package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jetstack/cert-manager/test/acme/dns"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

	fixture := dns.NewFixture(&dnsimpleSolver{},
		dns.SetResolvedZone(zone),
		dns.SetAllowAmbientCredentials(false),
		dns.SetManifestPath("testdata/dnsimple-solver"),
		dns.SetBinariesPath(filepath.Join(os.TempDir(), "kubebuilder/bin")),
	)

	fixture.RunConformance(t)
}
