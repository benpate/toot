package route

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// https://docs.joinmastodon.org/methods/accounts/#create
const PostAccount = "/api/v1/accounts"

// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
const GetAccount_VerifyCredentials = "/api/v1/accounts/verify_credentials"

// https://docs.joinmastodon.org/methods/accounts/#update_credentials
const PatchAccount_UpdateCredentials = "/api/v1/accounts/update_credentials"

// https://docs.joinmastodon.org/methods/accounts/#get
const GetAccount = "/api/v1/accounts/:id"

// https://docs.joinmastodon.org/methods/accounts/#statuses
const GetAccount_Statuses = "/api/v1/accounts/:id/statuses"

// https://docs.joinmastodon.org/methods/accounts/#followers
const GetAccount_Followers = "/api/v1/accounts/:id/followers"

// https://docs.joinmastodon.org/methods/accounts/#following
const GetAccount_Following = "/api/v1/accounts/:id/following"

// https://docs.joinmastodon.org/methods/accounts/#featured_tags
const GetAccount_FeaturedTags = "/api/v1/accounts/:id/featured_tags"

// https://docs.joinmastodon.org/methods/accounts/#lists
const GetAccount_Lists = "/api/v1/accounts/:id/lists"

// https://docs.joinmastodon.org/methods/accounts/#follow
const PostAccont_Follow = "/api/v1/accounts/:id/follow"

// https://docs.joinmastodon.org/methods/accounts/#unfollow
const PostAccount_Unfollow = "/api/v1/accounts/:id/unfollow"

// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
const PostAccount_RemoveFromFollowers = "/api/v1/accounts/:id/remove_from_followers"

// https://docs.joinmastodon.org/methods/accounts/#block
const PostAccount_Block = "/api/v1/accounts/:id/block"

// https://docs.joinmastodon.org/methods/accounts/#unblock
const PostAccount_Unblock = "/api/v1/accounts/:id/unblock"

// https://docs.joinmastodon.org/methods/accounts/#mute
const PostAccount_Mute = "/api/v1/accounts/:id/mute"

// https://docs.joinmastodon.org/methods/accounts/#unmute
const PostAccount_Unmute = "/api/v1/accounts/:id/unmute"

// https://docs.joinmastodon.org/methods/accounts/#pin
const PostAccount_Pin = "/api/v1/accounts/:id/pin"

// https://docs.joinmastodon.org/methods/accounts/#unpin
const PostAccount_Unpin = "/api/v1/accounts/:id/unpin"

// https://docs.joinmastodon.org/methods/accounts/#note
const PostAccount_Note = "/api/v1/accounts/:id/note"

// https://docs.joinmastodon.org/methods/accounts/#relationships
const PostAccount_Relationships = "/api/v1/accounts/relationships"

// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
const GetAccount_FamiliarFollowers = "/api/v1/accounts/:id/familiar_followers"

// https://docs.joinmastodon.org/methods/accounts/#search
const GetAccount_Search = "/api/v1/accounts/search"

// https://docs.joinmastodon.org/methods/accounts/#lookup
const GetAccount_Lookup = "/api/v1/accounts/lookup"
