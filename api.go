// Package toot implements a Mastodon API server, which opens up other applications to Mastodon clients.
package toot

import (
	"github.com/benpate/toot/object"
	"github.com/benpate/toot/txn"
)

// API represents the collection of routes and handlers that make up a Mastodon API.
// The API structure is dependent on an app-defined `AuthToken` that should be something like
// a JWT key that was obtained from an OAuth2 handshake.
type API[AuthToken ScopesGetter] struct {

	// Authorize method uses an http.Request to generate a app-defined AuthToken.
	// This property MUST be present for the API server to work correctly.
	Authorize Authorizer[AuthToken]

	// https://docs.joinmastodon.org/methods/accounts/
	PostAccount                    func(AuthToken, txn.PostAccount) (object.Token, error)
	GetAccount_VerifyCredentials   func(AuthToken, txn.GetAccount_VerifyCredentials) (object.Account, error)
	PatchAccount_UpdateCredentials func(AuthToken, txn.PatchAccount_UpdateCredentials) (object.Account, error)
	GetAccount                     func(AuthToken, txn.GetAccount) (object.Account, error)
	GetAccount_Statuses            func(AuthToken, txn.GetAccount_Statuses) ([]object.Status, error)
	GetAccount_Followers           func(AuthToken, txn.GetAccount_Followers) ([]object.Account, error)
	GetAccount_Following           func(AuthToken, txn.GetAccount_Following) ([]object.Account, error)
	GetAccount_FeaturedTags        func(AuthToken, txn.GetAccount_FeaturedTags) ([]object.Tag, error)
	PostAccount_Follow             func(AuthToken, txn.PostAccount_Follow) (object.Relationship, error)
	PostAccount_Unfollow           func(AuthToken, txn.PostAccount_Unfollow) (object.Relationship, error)
	PostAccount_Block              func(AuthToken, txn.PostAccount_Block) (object.Relationship, error)
	PostAccount_Unblock            func(AuthToken, txn.PostAccount_Unblock) (object.Relationship, error)
	PostAccount_Mute               func(AuthToken, txn.PostAccount_Mute) (object.Relationship, error)
	PostAccount_Unmute             func(AuthToken, txn.PostAccount_Unmute) (object.Relationship, error)
	PostAccount_Pin                func(AuthToken, txn.PostAccount_Pin) (object.Relationship, error)
	PostAccount_Unpin              func(AuthToken, txn.PostAccount_Unpin) (object.Relationship, error)
	PostAccount_Note               func(AuthToken, txn.PostAccount_Note) (object.Relationship, error)
	GetAccount_Relationships       func(AuthToken, txn.GetAccount_Relationships) ([]object.Relationship, error)
	GetAccount_FamiliarFollowers   func(AuthToken, txn.GetAccount_FamiliarFollowers) (object.FamiliarFollowers, error)
	GetAccount_Search              func(AuthToken, txn.GetAccount_Search) ([]object.Account, error)
	GetAccount_Lookup              func(AuthToken, txn.GetAccount_Lookup) (object.Account, error)

	// https://docs.joinmastodon.org/methods/announcements/
	GetAnnouncements            func(AuthToken, txn.GetAnnouncements) ([]object.Announcement, error)
	PostAnnouncement_Dismiss    func(AuthToken, txn.PostAnnouncement_Dismiss) (struct{}, error)
	PutAnnouncement_Reaction    func(AuthToken, txn.PutAnnouncement_Reaction) (struct{}, error)
	DeleteAnnouncement_Reaction func(AuthToken, txn.DeleteAnnouncement_Reaction) (struct{}, error)

	// https://docs.joinmastodon.org/methods/apps/
	PostApplication                  func(AuthToken, txn.PostApplication) (object.Application, error)
	GetApplication_VerifyCredentials func(AuthToken, txn.GetApplication_VerifyCredentials) (object.Application, error)

	// https://docs.joinmastodon.org/methods/blocks/
	GetBlocks func(AuthToken, txn.GetBlocks) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/bookmarks/
	GetBookmarks func(AuthToken, txn.GetBookmarks) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/conversations/
	GetConversations     func(AuthToken, txn.GetConversations) ([]object.Conversation, error)
	DeleteConversation   func(AuthToken, txn.DeleteConversation) (struct{}, error)
	PostConversationRead func(AuthToken, txn.PostConversationRead) (struct{}, error)

	// https://docs.joinmastodon.org/methods/custom_emojis/
	GetCustomEmojis func(AuthToken, txn.GetCustomEmojis) ([]object.CustomEmoji, error)

	// https://docs.joinmastodon.org/methods/directory/
	GetDirectory func(AuthToken, txn.GetDirectory) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/domain_blocks/
	GetDomainBlocks   func(AuthToken, txn.GetDomainBlocks) ([]string, error)
	PostDomainBlock   func(AuthToken, txn.PostDomainBlock) (struct{}, error)
	DeleteDomainBlock func(AuthToken, txn.DeleteDomainBlock) (struct{}, error)

	// https://docs.joinmastodon.org/methods/emails/
	PostEmailConfirmation func(AuthToken, txn.PostEmailConfirmation) (struct{}, error)

	// https://docs.joinmastodon.org/methods/endorsements/
	GetEndorsements func(AuthToken, txn.GetEndorsements) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/favourites/
	GetFavourites func(AuthToken, txn.GetFavourites) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/featured_tags/
	GetFeaturedTags             func(AuthToken, txn.GetFeaturedTags) ([]object.FeaturedTag, error)
	PostFeaturedTag             func(AuthToken, txn.PostFeaturedTag) (object.FeaturedTag, error)
	DeleteFeaturedTag           func(AuthToken, txn.DeleteFeaturedTag) (struct{}, error)
	GetFeaturedTags_Suggestions func(AuthToken, txn.GetFeaturedTags_Suggestions) ([]object.FeaturedTag, error)

	// https://docs.joinmastodon.org/methods/filters/
	GetFilters           func(AuthToken, txn.GetFilters) ([]object.Filter, error)
	GetFilter            func(AuthToken, txn.GetFilter) (object.Filter, error)
	PostFilter           func(AuthToken, txn.PostFilter) (object.Filter, error)
	PutFilter            func(AuthToken, txn.PutFilter) (object.Filter, error)
	DeleteFilter         func(AuthToken, txn.DeleteFilter) (struct{}, error)
	GetFilter_Keywords   func(AuthToken, txn.GetFilter_Keywords) ([]string, error)
	PostFilter_Keyword   func(AuthToken, txn.PostFilter_Keyword) (struct{}, error)
	GetFilter_Keyword    func(AuthToken, txn.GetFilter_Keyword) (struct{}, error)
	PutFilter_Keyword    func(AuthToken, txn.PutFilter_Keyword) (struct{}, error)
	DeleteFilter_Keyword func(AuthToken, txn.DeleteFilter_Keyword) (struct{}, error)
	GetFilter_Statuses   func(AuthToken, txn.GetFilter_Statuses) ([]object.FilterStatus, error)
	PostFilter_Status    func(AuthToken, txn.PostFilter_Status) (object.FilterStatus, error)
	GetFilter_Status     func(AuthToken, txn.GetFilter_Status) (object.FilterStatus, error)
	DeleteFilter_Status  func(AuthToken, txn.DeleteFilter_Status) (struct{}, error)
	GetFilter_V1         func(AuthToken, txn.GetFilter_V1) (object.Filter, error)
	PostFilter_V1        func(AuthToken, txn.PostFilter_V1) (object.Filter, error)
	PutFilter_V1         func(AuthToken, txn.PutFilter_V1) (object.Filter, error)
	DeleteFilter_V1      func(AuthToken, txn.DeleteFilter_V1) (struct{}, error)

	// https://docs.joinmastodon.org/methods/follow_requests/
	GetFollowRequests           func(AuthToken, txn.GetFollowRequests) ([]object.Account, error)
	PostFollowRequest_Authorize func(AuthToken, txn.PostFollowRequest_Authorize) (object.Relationship, error)
	PostFollowRequest_Reject    func(AuthToken, txn.PostFollowRequest_Reject) (object.Relationship, error)

	// https://docs.joinmastodon.org/methods/followed_tags/
	GetFollowedTags func(AuthToken, txn.GetFollowedTags) ([]object.Tag, error)

	// https://docs.joinmastodon.org/methods/instance/
	GetInstance                     func(AuthToken, txn.GetInstance) (object.Instance, error)
	GetInstance_Peers               func(AuthToken, txn.GetInstance_Peers) ([]string, error)
	GetInstance_Activity            func(AuthToken, txn.GetInstance_Activity) (map[string]any, error)
	GetInstance_Rules               func(AuthToken, txn.GetInstance_Rules) ([]object.Rule, error)
	GetInstance_DomainBlocks        func(AuthToken, txn.GetInstance_DomainBlocks) ([]object.DomainBlock, error)
	GetInstance_ExtendedDescription func(AuthToken, txn.GetInstance_ExtendedDescription) (object.ExtendedDescription, error)

	// https://docs.joinmastodon.org/methods/lists/
	GetLists            func(AuthToken, txn.GetLists) ([]object.List, error)
	GetList             func(AuthToken, txn.GetList) (object.List, error)
	PostList            func(AuthToken, txn.PostList) (object.List, error)
	PutList             func(AuthToken, txn.PutList) (object.List, error)
	DeleteList          func(AuthToken, txn.DeleteList) (struct{}, error)
	GetList_Accounts    func(AuthToken, txn.GetList_Accounts) ([]object.Account, error)
	PostList_Accounts   func(AuthToken, txn.PostList_Accounts) (struct{}, error)
	DeleteList_Accounts func(AuthToken, txn.DeleteList_Accounts) (struct{}, error)

	// https://docs.joinmastodon.org/methods/markers/
	GetMarkers func(AuthToken, txn.GetMarkers) (object.Marker, error)
	PostMarker func(AuthToken, txn.PostMarker) (object.Marker, error)

	// https://docs.joinmastodon.org/methods/media/
	PostMedia func(AuthToken, txn.PostMedia) (object.MediaAttachment, error)

	// https://docs.joinmastodon.org/methods/mutes/
	GetMutes func(AuthToken, txn.GetMutes) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/notifications/
	GetNotifications         func(AuthToken, txn.GetNotifications) ([]object.Notification, error)
	GetNotification          func(AuthToken, txn.GetNotification) (object.Notification, error)
	PostNotifications_Clear  func(AuthToken, txn.PostNotifications_Clear) (object.Notification, error)
	PostNotification_Dismiss func(AuthToken, txn.PostNotification_Dismiss) (object.Notification, error)

	// https://docs.joinmastodon.org/methods/oauth/
	GetOAuth_Authorize func(AuthToken, txn.GetOAuth_Authorize) (struct{}, error)
	PostOAuth_Token    func(AuthToken, txn.PostOAuth_Token) (object.Token, error)
	PostOAuth_Revoke   func(AuthToken, txn.PostOAuth_Revoke) (struct{}, error)

	// https://docs.joinmastodon.org/methods/oembed/
	GetOEmbed func(AuthToken, txn.GetOEmbed) (map[string]any, error)

	// https://docs.joinmastodon.org/methods/polls/
	GetPoll        func(AuthToken, txn.GetPoll) ([]object.Poll, error)
	PostPoll_Votes func(AuthToken, txn.PostPoll_Votes) ([]object.Poll, error)

	// https://docs.joinmastodon.org/methods/preferences/
	GetPreferences func(AuthToken, txn.GetPreferences) (map[string]any, error)

	// https://docs.joinmastodon.org/methods/profile/
	DeleteProfile_Avatar func(AuthToken, txn.DeleteProfile_Avatar) (object.Account, error)
	DeleteProfile_Header func(AuthToken, txn.DeleteProfile_Header) (object.Account, error)

	// https://docs.joinmastodon.org/methods/reports/
	PostReport func(AuthToken, txn.PostReport) (object.Report, error)

	// https://docs.joinmastodon.org/methods/scheduled_statuses/
	GetScheduledStatuses  func(AuthToken, txn.GetScheduledStatuses) ([]object.ScheduledStatus, error)
	GetScheduledStatus    func(AuthToken, txn.GetScheduledStatus) (object.ScheduledStatus, error)
	PutScheduledStatus    func(AuthToken, txn.PutScheduledStatus) (object.ScheduledStatus, error)
	DeleteScheduledStatus func(AuthToken, txn.DeleteScheduledStatus) (struct{}, error)

	// https://docs.joinmastodon.org/methods/search/
	GetSearch func(AuthToken, txn.GetSearch) (object.Search, error)

	// https://docs.joinmastodon.org/methods/statuses/#create
	PostStatus             func(AuthToken, txn.PostStatus) (object.Status, error)
	GetStatus              func(AuthToken, txn.GetStatus) (object.Status, error)
	DeleteStatus           func(AuthToken, txn.DeleteStatus) (struct{}, error)
	GetStatus_Context      func(AuthToken, txn.GetStatus_Context) (object.Context, error)
	PostStatus_Translate   func(AuthToken, txn.PostStatus_Translate) (object.Status, error)
	GetStatus_RebloggedBy  func(AuthToken, txn.GetStatus_RebloggedBy) ([]object.Account, error)
	GetStatus_FavouritedBy func(AuthToken, txn.GetStatus_FavouritedBy) ([]object.Account, error)
	PostStatus_Favourite   func(AuthToken, txn.PostStatus_Favourite) (object.Status, error)
	PostStatus_Unfavourite func(AuthToken, txn.PostStatus_Unfavourite) (object.Status, error)
	PostStatus_Reblog      func(AuthToken, txn.PostStatus_Reblog) (object.Status, error)
	PostStatus_Unreblog    func(AuthToken, txn.PostStatus_Unreblog) (object.Status, error)
	PostStatus_Bookmark    func(AuthToken, txn.PostStatus_Bookmark) (object.Status, error)
	PostStatus_Unbookmark  func(AuthToken, txn.PostStatus_Unbookmark) (object.Status, error)
	PostStatus_Mute        func(AuthToken, txn.PostStatus_Mute) (object.Status, error)
	PostStatus_Unmute      func(AuthToken, txn.PostStatus_Unmute) (object.Status, error)
	PostStatus_Pin         func(AuthToken, txn.PostStatus_Pin) (object.Status, error)
	PostStatus_Unpin       func(AuthToken, txn.PostStatus_Unpin) (object.Status, error)
	PutStatus              func(AuthToken, txn.PutStatus) (object.Status, error)
	GetStatus_History      func(AuthToken, txn.GetStatus_History) ([]object.StatusEdit, error)
	GetStatus_Source       func(AuthToken, txn.GetStatus_Source) (object.StatusSource, error)

	// https://docs.joinmastodon.org/methods/suggestions/
	GetSuggestions   func(AuthToken, txn.GetSuggestions) ([]object.Suggestion, error)
	DeleteSuggestion func(AuthToken, txn.DeleteSuggestion) (struct{}, error)

	// https://docs.joinmastodon.org/methods/tags/
	GetTag           func(AuthToken, txn.GetTag) (object.Tag, error)
	PostTag_Follow   func(AuthToken, txn.PostTag_Follow) (object.Tag, error)
	PostTag_Unfollow func(AuthToken, txn.PostTag_Unfollow) (object.Tag, error)

	// https://docs.joinmastodon.org/methods/timelines/
	GetTimeline_Public  func(AuthToken, txn.GetTimeline_Public) ([]object.Status, error)
	GetTimeline_Hashtag func(AuthToken, txn.GetTimeline_Hashtag) ([]object.Status, error)
	GetTimeline_Home    func(AuthToken, txn.GetTimeline_Home) ([]object.Status, error)
	GetTimeline_List    func(AuthToken, txn.GetTimeline_List) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/trends/
	GetTrends          func(AuthToken, txn.GetTrends) ([]object.Tag, error)
	GetTrends_Statuses func(AuthToken, txn.GetTrends_Statuses) ([]object.Status, error)
	GetTrends_Links    func(AuthToken, txn.GetTrends_Links) ([]object.PreviewCard, error)
}

// New returns a bare bones API structure with no methods implemented.
func New[AuthToken ScopesGetter](authorizer Authorizer[AuthToken]) API[AuthToken] {
	return API[AuthToken]{
		Authorize: authorizer,
	}
}
