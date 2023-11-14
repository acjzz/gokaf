// Gokaf is a simple In-memory PubSub Engine
package gokaf

// Logging Interface
type Logger interface {
	Printf(format string, v ...interface{})
}
