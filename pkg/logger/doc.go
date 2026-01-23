// Package logger defines interface for logging Debug, Info, Warning and Error messages.
// At the moment, the implementation relies on the std log package.
// The idea is, however, that it's possible to swap this for another
// implementation without disturbing client code.
package logger
