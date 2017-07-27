package consul

import (
	"fmt"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/vault/plugins/helper/database/connutil"
	"github.com/mitchellh/mapstructure"
)

type consulConnectionProducer struct {
	Address     string `json:"address" structs:"address" mapstructure:"address"`
	Scheme      string `json:"scheme" structs:"scheme" mapstructure:"scheme"`
	Token       string `json:"token" structs:"token" mapstructure:"token"`
	Initialized bool
	Type        string
	sync.Mutex
	session *api.Client
}

// Initialize sets the consul client configuration
func (c *consulConnectionProducer) Initialize(conf map[string]interface{}, verifyConnection bool) error {
	c.Lock()
	defer c.Unlock()

	err := mapstructure.WeakDecode(conf, c)
	if err != nil {
		return err
	}
	switch {
	case len(c.Address) == 0:
		return fmt.Errorf("address cannot be empty")
	case len(c.Scheme) == 0:
		return fmt.Errorf("scheme cannot be empty")
	case len(c.Token) == 0:
		return fmt.Errorf("token cannot be empty")
	}

	c.Initialized = true

	if verifyConnection {
		if _, err := c.Connection(); err != nil {
			return fmt.Errorf("error verifying connection: %s", err)
		}

		if _, err := c.session.Status().Leader(); err != nil {
			return fmt.Errorf("error verifying connection: %s", err)
		}
	}

	return nil
}

//Connection creates a Consul client
func (c *consulConnectionProducer) Connection() (interface{}, error) {
	var err error
	if !c.Initialized {
		return nil, connutil.ErrNotInitialized
	}
	if c.session != nil {
		return c.session, nil
	}

	consulConf := api.DefaultNonPooledConfig()
	consulConf.Address = c.Address
	consulConf.Scheme = c.Scheme
	consulConf.Token = c.Token

	c.session, err = api.NewClient(consulConf)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *consulConnectionProducer) Close() error {
	c.Lock()
	defer c.Unlock()
	if c.session != nil {
	}
	return nil
}
