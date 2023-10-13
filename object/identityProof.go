package object

// IdentityProof represents a proof from an external identity provider.
// https://docs.joinmastodon.org/entities/IdentityProof/
type IdentityProof struct {
	Provider         string `json:"provider"`          // The name of the identity provider.
	ProviderUsername string `json:"provider_username"` // The account owner’s username on the identity provider’s service.
	UpdatedAt        string `json:"updated_at"`        // When the identity proof was last updated. (ISO 8601 Datetime)
	ProofURL         string `json:"proof_url"`         // A link to a statement of identity proof, hosted by the identity provider.
	ProfileURL       string `json:"profile_url"`       // The account owner’s profile URL on the identity provider.
}
