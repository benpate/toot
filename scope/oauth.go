package scope

/******************************************
 * OAuth API Methods
 * Generate and manage OAuth tokens
 * https://docs.joinmastodon.org/methods/oauth/
******************************************/

// https://docs.joinmastodon.org/methods/oauth/#authorize
const GetOAuth_Authorize = Public

// https://docs.joinmastodon.org/methods/oauth/#token
const PostOAuth_Token = Public

// https://docs.joinmastodon.org/methods/oauth/#revoke
const PostOAuth_Revoke = Public
