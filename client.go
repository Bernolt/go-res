package res

import "github.com/jirenius/resgate/server/reserr"

// Response sends a response to the messaging system
type Response func(subj string, payload []byte, err error)

// Unsubscriber is the interface that wraps the basic Unsubscribe method
type Unsubscriber interface {
	// Unsubscribe cancels the subscription
	Unsubscribe() error
}

// Client is an interface that represents a client to a messaging system.
type Client interface {
	// SendRequest sends an asynchronous request on a subject, expecting the Response
	// callback to be called once.
	SendRequest(subject string, payload []byte, cb Response)

	// Subscribe to all events on a resource namespace.
	// The namespace has the format "event."+resource
	Subscribe(namespace string, cb Response) (Unsubscriber, error)

	// Close closes the connection.
	Close()

	// IsClosed tests if the connection has been closed.
	IsClosed() bool

	// Sets the closed handler
	SetClosedHandler(cb func(error))
}

// ErrRequestTimeout is the error the client should pass to the Response
// when a call to SendRequest times out
var ErrRequestTimeout = reserr.ErrTimeout
