package client

import (
	"crypto/x509"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	c := NewConfiguration()
	assert.Equal(t, "/", c.BasePath)
	assert.NotNil(t, c.DefaultHeader)
	assert.Equal(t, "Swagger-Codegen/1.0.0/go", c.UserAgent)
}

func TestConfiguration_AddDefaultHeader(t *testing.T) {
	c := NewConfiguration()

	c.AddDefaultHeader("X-Header", "X-Header-Value")
	v := c.DefaultHeader["X-Header"]
	assert.Equal(t, "X-Header-Value", v)
}

func TestConfiguration_SetClientCertificate(t *testing.T) {
	c := NewConfiguration()

	assert.Nil(t, c.SetClientCertificate(
		path.Join("..", "samples", "cert.pem"),
		path.Join("..", "samples", "key.pem")))

	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)
}

func TestConfiguration_SetClientCertificateEx(t *testing.T) {
	c := NewConfiguration()

	assert.Nil(t, c.SetClientCertificateEx(
		path.Join("..", "samples", "cert.pem"),
		path.Join("..", "samples", "key.pem"), true))

	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)

	c = NewConfiguration()
	assert.Error(t, c.SetClientCertificate(
		path.Join("..", "samples", "cert.pem-none"),
		path.Join("..", "samples", "key.pem")))

	c = NewConfiguration()
	assert.Error(t, c.SetClientCertificate(
		path.Join("..", "samples", "cert.pem"),
		path.Join("..", "samples", "key.pem-none")))
}

func TestConfiguration_SetClientCertificatePKCS12(t *testing.T) {
	c := NewConfiguration()

	assert.Nil(t, c.SetClientCertificatePKCS12(
		path.Join("..", "samples", "sample.pfx"),
		"password"))
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)
}

func TestConfiguration_SetClientCertificatePKCS12Ex(t *testing.T) {
	c := NewConfiguration()

	assert.Nil(t, c.SetClientCertificatePKCS12Ex(
		path.Join("..", "samples", "sample.pfx"),
		"password", true))
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)

	c = NewConfiguration()
	assert.Error(t, c.SetClientCertificatePKCS12Ex(
		path.Join("..", "samples", "sample.pfx"),
		"bad-password", true))
}

func TestConfiguration_createHTTPClient(t *testing.T) {
	ref := NewConfiguration()

	require.Nil(t, ref.SetClientCertificatePKCS12Ex(
		path.Join("..", "samples", "sample.pfx"),
		"password", true))

	c := NewConfiguration()
	c.ClientCertificates = ref.ClientCertificates
	c.CertPool = &x509.CertPool{}
	assert.Nil(t, c.createHTTPClient(true))
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)
}

func TestConfiguration_Init(t *testing.T) {

	// PEM
	c := NewConfiguration()
	c.CertFile = path.Join("..", "samples", "cert.pem")
	c.KeyFile = path.Join("..", "samples", "key.pem")
	assert.Nil(t, c.Init())
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)

	// Ensure pfx over PEM
	c = NewConfiguration()
	c.CertFile = path.Join("..", "samples", "cert.pem2")
	c.KeyFile = path.Join("..", "samples", "key.pem2")
	c.PFXFile = path.Join("..", "samples", "sample.pfx")
	c.PFXPassword = "password"
	assert.Nil(t, c.Init())
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)

	// Ensure pfx over PEM
	clientCerts := c.ClientCertificates
	certPool := c.CertPool
	c = NewConfiguration()
	c.ClientCertificates = clientCerts
	c.CertPool = certPool
	c.CertFile = path.Join("..", "samples", "cert.pem2")
	c.KeyFile = path.Join("..", "samples", "key.pem2")
	c.PFXFile = path.Join("..", "samples", "sample.pfx2")
	c.PFXPassword = "password2"
	assert.Nil(t, c.Init())
	assert.NotNil(t, c.HTTPClient)
	assert.NotNil(t, c.ClientCertificates)
	assert.NotNil(t, c.CertPool)

	// None
	c = NewConfiguration()
	assert.ErrorContains(t, c.Init(), "no client certificate set")
}
