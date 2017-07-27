package consul

import (
	"time"

	"github.com/hashicorp/vault/builtin/logical/database/dbplugin"
	"github.com/hashicorp/vault/plugins/helper/database/connutil"
	"github.com/hashicorp/vault/plugins/helper/database/credsutil"
)

const consulTypeName = "consul"

// Consul type is an implementation of the Database interface
type Consul struct {
	connutil.ConnectionProducer
	credsutil.CredentialsProducer
}

// New create a new Consul instance
func New() (interface{}, error) {
	// connProducer := &consulConnectionProducer{}
	return nil, nil
}

// Type returns the TypeName for this backend
func (c *Consul) Type() (string, error) {
	return consulTypeName, nil
}

// CreateUser generates a Consul token
func (c *Consul) CreateUser(statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error) {
	// Lock Consul
	c.Lock()
	defer c.Unlock()
	return "", "", nil
}

// RenewUser renews Consul token
func (c *Consul) RenewUser(statements dbplugin.Statements, username string, expiration time.Time) error {
	return nil
}

// RevokeUser destroys the Consul token
func (c *Consul) RevokeUser(statements dbplugin.Statements, username string) error {
	return nil
}

// Close is not supported on Consul, so this will be a no-op
func (c *Consul) Close() error {
	return nil
}
