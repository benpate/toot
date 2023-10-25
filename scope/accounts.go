package scope

/******************************************
 * Accounts API Methods
 * Methods concerning accounts and profiles
 * https://docs.joinmastodon.org/methods/accounts/
 ******************************************/

// https://docs.joinmastodon.org/methods/accounts/#create
const PostAccount = WriteAccounts

// https://docs.joinmastodon.org/methods/accounts/#verify_credentials
const GetAccount_VerifyCredentials = ReadAccounts

// https://docs.joinmastodon.org/methods/accounts/#update_credentials
const PatchAccount_UpdateCredentials = WriteAccounts

// https://docs.joinmastodon.org/methods/accounts/#get
const GetAccount = Public

// https://docs.joinmastodon.org/methods/accounts/#statuses
const GetAccount_Statuses = ReadStatuses

// https://docs.joinmastodon.org/methods/accounts/#followers
const GetAccount_Followers = Private

// https://docs.joinmastodon.org/methods/accounts/#following
const GetAccount_Following = Private

// https://docs.joinmastodon.org/methods/accounts/#featured_tags
const GetAccount_FeaturedTags = Private

// https://docs.joinmastodon.org/methods/accounts/#lists
const GetAccount_Lists = ReadLists

// https://docs.joinmastodon.org/methods/accounts/#follow
const PostAccont_Follow = WriteFollows

// https://docs.joinmastodon.org/methods/accounts/#unfollow
const PostAccount_Unfollow = WriteFollows

// https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
const PostAccount_RemoveFromFollowers = WriteFollows

// https://docs.joinmastodon.org/methods/accounts/#block
const PostAccount_Block = WriteBlocks

// https://docs.joinmastodon.org/methods/accounts/#unblock
const PostAccount_Unblock = WriteBlocks

// https://docs.joinmastodon.org/methods/accounts/#mute
const PostAccount_Mute = WriteMutes

// https://docs.joinmastodon.org/methods/accounts/#unmute
const PostAccount_Unmute = WriteMutes

// https://docs.joinmastodon.org/methods/accounts/#pin
const PostAccount_Pin = WriteAccounts

// https://docs.joinmastodon.org/methods/accounts/#unpin
const PostAccount_Unpin = WriteAccounts

// https://docs.joinmastodon.org/methods/accounts/#note
const PostAccount_Note = WriteAccounts

// https://docs.joinmastodon.org/methods/accounts/#relationships
const PostAccount_Relationships = ReadFollows

// https://docs.joinmastodon.org/methods/accounts/#familiar_followers
const GetAccount_FamiliarFollowers = ReadFollows

// https://docs.joinmastodon.org/methods/accounts/#search
const GetAccount_Search = ReadAccounts

// https://docs.joinmastodon.org/methods/accounts/#lookup
const GetAccount_Lookup = Private
