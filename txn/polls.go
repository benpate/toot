package txn

/******************************************
* Polls API Methods
* View and vode on polls attached to statuses. To discover pollID,
* you will need to GET a Status first and then check for a `poll` property.
* https://docs.joinmastodon.org/methods/polls/
******************************************/

// https://docs.joinmastodon.org/methods/polls/#get
// GET /api/v1/polls/:id
// Returns: Poll
// View a Poll
type GetPoll struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
}

// https://docs.joinmastodon.org/methods/polls/#vote
// POST /api/v1/polls/:id/votes
// Returns: Poll
// Vote on a Poll
type PostPoll_Votes struct {
	Host          string `header:"Host"`
	Authorization string `header:"Authorization"`
	ID            string `param:"id"`
	Choices       []int  `json:"choices"`
}
