package object

// DomainBlock represents a domain that is blocked by the instance.
// https://docs.joinmastodon.org/entities/DomainBlock/
type DomainBlock struct {
	Domain   string `json:"domain"`            // The domain which is blocked. This may be obfuscated or partially censored.
	Digest   string `json:"digest"`            // The SHA256 hash digest of the domain string.
	Severity string `json:"severity"`          // The level to which the domain is blocked. [silence | suspend]
	Comment  string `json:"comment,omitempty"` // An optional reason for the domain block.
}

// Users from this domain will be hidden from timelines, threads, and notifications (unless you follow the user).
const DomainBlockSeveritySilence = "silence"

// Incoming messages from this domain will be rejected and dropped entirely.
const DomainBlockSeveritySuspend = "suspend"
