package route

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// https://docs.joinmastodon.org/methods/lists/#get
const GetLists = "/api/v1/lists"

// https://docs.joinmastodon.org/methods/lists/#get-one
const GetList = "/api/v1/lists/:id"

// https://docs.joinmastodon.org/methods/lists/#create
const PostList = "/api/v1/lists"

// https://docs.joinmastodon.org/methods/lists/#update
const PutList = "/api/v1/lists/:id"

// https://docs.joinmastodon.org/methods/lists/#delete
const DeleteList = "/api/v1/lists/:id"

// https://docs.joinmastodon.org/methods/lists/#accounts
const GetList_Accounts = "/api/v1/lists/:id/accounts"

// https://docs.joinmastodon.org/methods/lists/#accounts-add
const PostList_Accounts = "/api/v1/lists/:id/accounts"

// https://docs.joinmastodon.org/methods/lists/#accounts-remove
const DeleteList_Accounts = "/api/v1/lists/:id/accounts"
