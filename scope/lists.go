package scope

/******************************************
 * Lists API Methods
 * View and manage lists. See also /api/v1/timelines/list/:id for loading a list timeline
 * https://docs.joinmastodon.org/methods/lists/
 ******************************************/

// https://docs.joinmastodon.org/methods/lists/#get
const GetLists = ReadLists

// https://docs.joinmastodon.org/methods/lists/#get-one
const GetList = WriteLists

// https://docs.joinmastodon.org/methods/lists/#create
const PostList = WriteLists

// https://docs.joinmastodon.org/methods/lists/#update
const PutList = WriteLists

// https://docs.joinmastodon.org/methods/lists/#delete
const DeleteList = WriteLists

// https://docs.joinmastodon.org/methods/lists/#accounts
const GetList_Accounts = ReadLists

// https://docs.joinmastodon.org/methods/lists/#accounts-add
const PostList_Accounts = WriteLists

// https://docs.joinmastodon.org/methods/lists/#accounts-remove
const DeleteList_Accounts = WriteLists
