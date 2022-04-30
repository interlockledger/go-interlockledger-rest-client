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

type DocumentsMetadataModel struct {
	Comment              string           `json:"comment,omitempty"`
	Compression          string           `json:"compression,omitempty"`
	Encryption           string           `json:"encryption,omitempty"`
	EncryptionParameters *Parameters      `json:"encryptionParameters,omitempty"`
	PublicDirectory      []DirectoryEntry `json:"publicDirectory,omitempty"`
}
