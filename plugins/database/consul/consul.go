package consul

import (
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/vault/builtin/logical/database/dbplugin"
	"github.com/hashicorp/vault/plugins"
	"github.com/hashicorp/vault/plugins/helper/database/connutil"
	"github.com/hashicorp/vault/plugins/helper/database/credsutil"
)

const consulTypeName = "consul"

// Consul type is an implementation of the Database interface
type Consul struct {
	connutil.ConnectionProducer
	credsutil.CredentialsProducer
}

// New create a new Consul Connection instance
func New() (interface{}, error) {
	connProducer := &consulConnectionProducer{}
	connProducer.Type = consulTypeName

	dbType := &Consul{
		ConnectionProducer:  connProducer,
		CredentialsProducer: nil,
	}

	return dbType, nil
}

// Run instantiates a Consul object, and runs the RPC server for the plugin
func Run(apiTLSConfig *api.TLSConfig) error {
	dbType, err := New()
	if err != nil {
		return err
	}

	plugins.Serve(dbType.(*Consul), apiTLSConfig)

	return nil
}

// Type returns the TypeName for this backend
func (c *Consul) Type() (string, error) {
	return consulTypeName, nil
}

func (c *Consul) getConnection() (*api.Client, error) {
	conn, err := c.Connection()
	if err != nil {
		return nil, err
	}

	return conn.(*api.Client), nil
}

// CreateUser generates a Consul token
func (c *Consul) CreateUser(statements dbplugin.Statements, usernameConfig dbplugin.UsernameConfig, expiration time.Time) (username string, password string, err error) {
	// Lock Consul
	c.Lock()
	defer c.Unlock()

	client := c.getConnection()
	client.ACL().
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
