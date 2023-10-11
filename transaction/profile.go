package transaction

/******************************************
 * Profile API Methods
 * Methods concerning profiles
 * https://docs.joinmastodon.org/methods/profile/
 ******************************************/

// https://docs.joinmastodon.org/methods/profile/#delete-profile-avatar
// DELETE /api/v1/profile/avatar
// Returns: CredentialAccount
type DeleteProfile_Avatar struct{}

// https://docs.joinmastodon.org/methods/profile/#delete-profile-header
// DELETE /api/v1/profile/header
// Returns: CredentialAccount
type DeleteProfile_Header struct{}
