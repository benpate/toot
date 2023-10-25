package scope

/******************************************
* Polls API Methods
* View and vode on polls attached to statuses. To discover pollID,
* you will need to GET a Status first and then check for a `poll` property.
* https://docs.joinmastodon.org/methods/polls/
******************************************/

// https://docs.joinmastodon.org/methods/polls/#get
const GetPoll = ReadStatuses

// https://docs.joinmastodon.org/methods/polls/#vote
const PostPoll_Votes = WriteStatuses
