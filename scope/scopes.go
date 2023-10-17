package scopes

/******************************************
 * OAuth Scopes
 * Defining what you have permission to do with the API
 * https://docs.joinmastodon.org/api/oauth-scopes/
 ******************************************/

// https://docs.joinmastodon.org/api/oauth-scopes/#read
const Read = "read"

const ReadAccounts = "read:accounts"

const ReadBlocks = "read:blocks"

const ReadBookmarks = "read:bookmarks"

const ReadFavourites = "read:favourites"

const ReadFilters = "read:filters"

const ReadFollows = "read:follows"

const ReadLists = "read:lists"

const ReadMutes = "read:mutes"

const ReadNotifications = "read:notifications"

const ReadSearch = "read:search"

const ReadStatuses = "read:statuses"

// https://docs.joinmastodon.org/api/oauth-scopes/#write
const Write = "write"

const WriteAccounts = "write:accounts"

const WriteBlocks = "write:blocks"

const WriteBookmarks = "write:bookmarks"

const WriteConversations = "write:conversations"

const WriteFavourites = "write:favourites"

const WriteFilters = "write:filters"

const WriteFollows = "write:follows"

const WriteLists = "write:lists"

const WriteMedia = "write:media"

const WriteMutes = "write:mutes"

const WriteNotifications = "write:notifications"

const WriteReports = "write:reports"

const WriteStatuses = "write:statuses"

// https://docs.joinmastodon.org/api/oauth-scopes/#push
const Push = "push"

// https://docs.joinmastodon.org/api/oauth-scopes/#admin
const AdminRead = "admin:read"

const AdminReadAccounts = "admin:read:accounts"

const AdminReadReports = "admin:read:reports"

const AdminReadDomainAllows = "admin:read:domain_allows"

const AdminReadDomainBlocks = "admin:read:domain_blocks"

const AdminReadIPBlocks = "admin:read:ip_blocks"

const AdminReadEmailDomainBlocks = "admin:read:email_domain_blocks"

const AdminReadCanonicalEmailBlocks = "admin:read:canonical_email_blocks"

const AdminWrite = "admin:write"

const AdminWriteAccounts = "admin:write:accounts"

const AdminWriteReports = "admin:write:reports"

const AdminWriteDomainAllows = "admin:write:domain_allows"

const AdminWriteDomainBlocks = "admin:write:domain_blocks"

const AdminWriteIPBlocks = "admin:write:ip_blocks"

const AdminWriteEmailDomainBlocks = "admin:write:email_domain_blocks"

const AdminWriteCanonicalEmailBlocks = "admin:write:canonical_email_blocks"
