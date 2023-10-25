package scope

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// https://docs.joinmastodon.org/methods/filters/#get
const GetFilters = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#get-one
const GetFilter = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#create
const PostFilter = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#update
const PutFilter = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#delete
const DeleteFilter = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#keywords-get
const GetFilter_Keywords = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#keywords-create
const PostFilter_Keyword = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
const GetFilter_Keyword = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#keywords-update
const PutFilter_Keyword = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#keywords-delete
const DeleteFilter_Keyword = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#statuses-get
const GetFilter_Statuses = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#statuses-add
const PostFilter_Status = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
const GetFilter_Status = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#statuses-remove
const DeleteFilter_Status = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#get-v1
const GetFilters_V1 = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#get-one-v1
const GetFilter_V1 = ReadFilters

// https://docs.joinmastodon.org/methods/filters/#create-v1
const PostFilter_V1 = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#update-v1
const PutFilter_V1 = WriteFilters

// https://docs.joinmastodon.org/methods/filters/#delete-v1
const DeleteFilter_V1 = WriteFilters
