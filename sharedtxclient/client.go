package sharedtxclient

import (
	"sync"
//	"time"

	"google.golang.org/grpc"

	"github.com/shweini/dcrdtx/utils"
	"github.com/decred/dcrwallet/sharedtxclient/service"
)

type Config struct {
	Host          string
   	Enable        bool
	RetryInterval uint32
}

type Client struct {
	*Config
	connection *grpc.ClientConn
	lock       sync.Mutex
	wg         sync.WaitGroup

	*service.TransactionService
	*service.VersionService
	*service.MessageVerificationService
}

var (
	logger     = utils.MustGetLogger("dcrdtxclient")
	logFormat  = "[dcrdtx.%{module}:%{level}] %{message}"
	logModules = []string{
		"dcrdtxclient",
	}
)

func NewClient(config *Config) (*Client, error) {
	client := &Client{
		Config: config,
	}

	// connect to grpc server
	conn, err := client.Connect()
	if err != nil {
		return nil, err
	}

	// somehow, connection is still null
	// return error
	if conn == nil {
		return nil, ErrCannotConnect
	}

	client.connection = conn

	// register our client services
    err = client.registerServiceClients()
    if err != nil {
        return nil, err
    }

	return client, nil
}

// Connect connects to sharedtx server based on config passed upon
// initilization.

// TODO @jokyjor..can we trust the default grpc auto reconnect?
// TODO @jokyjor..autoconnect backoff implementation?
func (c *Client) Connect() (*grpc.ClientConn, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.IsConnected() {
		return nil, ErrClientAlreadyConnected
	}

	logger.Infof("Attempting to connect to SharedTx Server")

	conn, err := grpc.Dial(c.Host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	logger.Infof("Successfully connected to SharedTx Server on: %s", c.Host)
	return conn, nil
}

// registerServiceClients registers the services that should be made available
// on dcrwallet
func (c *Client) registerServiceClients() error {
	if !c.IsConnected() {
		return ErrClientNotConnected
	}

	c.TransactionService = service.NewTransactionService(c.connection)
	c.MessageVerificationService = service.NewMessageVerificationService(c.connection)
	c.VersionService = service.NewVersionService(c.connection)
	return nil
}


// Disconnect disconnects client from server
// Returns an error if client is not connected
func (c *Client) Disconnect() error {
	if c.IsConnected() {
		c.connection.Close()
		return nil
	}
	return ErrClientNotConnected
}

// IsConnected returns a boolean which trutthiness depends
// On whether the client is connected or not
func (c *Client) IsConnected() bool {
	if c.connection != nil {
		return true
	}
	return false
}
