// Package version carries the build-time version string stamped into
// davinci-dkg binaries via `-ldflags "-X .../version.Version=..."`.
//
// The default value is "dev"; release builds substitute a semantic version
// or git describe output. See the Dockerfile for the link flags.
package version
