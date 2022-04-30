/*
 * InterlockLedger Node REST API
 *
 * A node instance inside the peer-to-peer network of the InterlockLedger
 *
 * API version: v7.2.0
 * Contact: core@interlockledger.network
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

// Node details
type NodeDetailsModel struct {
	// List of owned chains, only the ids
	Chains []string `json:"chains,omitempty"`
	// Mapping color
	Color string `json:"color,omitempty"`
	// Dictionary with values for extensions on node configuration
	Extensions map[string]string `json:"extensions,omitempty"`
	// Unique node id
	Id string `json:"id"`
	// Node name
	Name string `json:"name,omitempty"`
	// Network this node participates on
	Network string `json:"network,omitempty"`
	// Node owner id [Optional]
	OwnerId string `json:"ownerId,omitempty"`
	// Node owner name [Optional]
	OwnerName string `json:"ownerName,omitempty"`
	// Node url for Peer2Peer protocol
	PeerAddress string `json:"peerAddress,omitempty"`
	// List of active roles running in the node
	Roles            []string          `json:"roles,omitempty"`
	SoftwareVersions *SoftwareVersions `json:"softwareVersions,omitempty"`
}
