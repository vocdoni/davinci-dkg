// Package log is a thin wrapper around zerolog that standardizes the log
// format used by every davinci-dkg binary.
//
// It exposes structured helpers (Infow, Warnw, Errorw, Debugw) and a single
// Init function to configure level + output. All services in this repo
// must use this package — do not create fresh zerolog / log.Logger
// instances elsewhere, so the output remains consistent.
package log
