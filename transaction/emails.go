package transaction

/******************************************
 * Emails API Methods
 * Request a new confirmation email, potentially to a new email address
 * https://docs.joinmastodon.org/methods/emails/
 ******************************************/

// https://docs.joinmastodon.org/methods/emails/#confirmation
// POST /api/v1/email/confirmation
// Returns: Empty object
type PostEmailConfirmation struct{}
