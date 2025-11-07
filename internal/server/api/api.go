package api

import "strings"

const (
	Wallet  = "/api/{version}/wallet"
	Wallets = "/api/{version}/wallets/"
)

func ConfigureEndpoint(endpoint string, version string) string {
	return strings.Replace(endpoint, "{version}", version, 1)
}
