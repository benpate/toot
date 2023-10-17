package route

/******************************************
 * Filters API Methods
 * Create and manage filters
 * https://docs.joinmastodon.org/methods/filters/
 ******************************************/

// https://docs.joinmastodon.org/methods/filters/#get
const GetFilters = "/api/v2/filters"

// https://docs.joinmastodon.org/methods/filters/#get-one
const GetFilter = "/api/v2/filters/:id"

// https://docs.joinmastodon.org/methods/filters/#create
const PostFilter = "/api/v2/filters"

// https://docs.joinmastodon.org/methods/filters/#update
const PutFilter = "/api/v2/filters/:id"

// https://docs.joinmastodon.org/methods/filters/#delete
const DeleteFilter = "/api/v2/filters/:id"

// https://docs.joinmastodon.org/methods/filters/#keywords-get
const GetFilter_Keywords = "/api/v2/filters/:filter_id/keywords"

// https://docs.joinmastodon.org/methods/filters/#keywords-create
const PostFilter_Keyword = "/api/v2/filters/:filter_id/keywords"

// https://docs.joinmastodon.org/methods/filters/#keywords-get-one
const GetFilter_Keyword = "/api/v2/filters/keywords/:id"

// https://docs.joinmastodon.org/methods/filters/#keywords-update
const PutFilter_Keyword = "/api/v2/filters/keywords/:id"

// https://docs.joinmastodon.org/methods/filters/#keywords-delete
const DeleteFilter_Keyword = "/api/v2/filters/keywords/:id"

// https://docs.joinmastodon.org/methods/filters/#statuses-get
const GetFilter_Statuses = "/api/v2/filters/:filter_id/statuses"

// https://docs.joinmastodon.org/methods/filters/#statuses-add
const PostFilter_Status = "/api/v2/filters/:filter_id/statuses"

// https://docs.joinmastodon.org/methods/filters/#statuses-get-one
const GetFilter_Status = "/api/v2/filters/statuses/:id"

// https://docs.joinmastodon.org/methods/filters/#statuses-remove
const DeleteFilter_Status = "/api/v2/filters/statuses/:id"

// https://docs.joinmastodon.org/methods/filters/#get-v1
const GetFilters_V1 = "/api/v1/filters"

// https://docs.joinmastodon.org/methods/filters/#get-one-v1
const GetFilter_V1 = "/api/v1/filters/:id"

// https://docs.joinmastodon.org/methods/filters/#create-v1
const PostFilter_V1 = "/api/v1/filters"

// https://docs.joinmastodon.org/methods/filters/#update-v1
const PutFilter_V1 = "/api/v1/filters/:id"

// https://docs.joinmastodon.org/methods/filters/#delete-v1
const DeleteFilter_V1 = "/api/v1/filters/:id"
