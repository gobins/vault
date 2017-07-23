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

// New function for creating new connection
func New() (interface{}, error) {
	// connProducer := &consulConnectionProducer{}
	return nil, nil
}

// Type returns the TypeName for this backend
func (c *Consul) Type() (string, error) {
	return consulTypeName, nil
}

// CreateUser generates a token
func (c *Consul) CreateUser(statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error) {
	// Lock Consul
	c.Lock()
	defer c.Unlock()

}
