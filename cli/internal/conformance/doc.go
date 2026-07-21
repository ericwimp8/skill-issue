// Package conformance drives complete skill-issue evaluate run invocations
// against scripted fake harness binaries that speak each supported harness's
// native protocol. The suite proves the full adapter, instrumentation,
// attribution, artifact, and cleanup path on any machine with no vendor
// credentials: happy paths execute the real injected signal commands, and
// failure cases cover protocol violations, mid-turn death, session-ID
// changes, configuration rejection, and version-pin enforcement.
package conformance
