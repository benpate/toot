// Package toot implements a Mastodon API server, which opens up other applications to Mastodon clients.
package toot

import (
	"net/http"

	"github.com/benpate/toot/object"
	"github.com/benpate/toot/txn"
)

type API struct {

	// Authorize returns a middleware handler that verifies that the current user has the specified scopes.
	Authorize func(*http.Request, ...string) bool

	// https://docs.joinmastodon.org/methods/accounts/
	PostAccount                    func(*http.Request, txn.PostAccount) (object.Token, error)
	GetAccount_VerifyCredentials   func(*http.Request, txn.GetAccount_VerifyCredentials) (object.Account, error)
	PatchAccount_UpdateCredentials func(*http.Request, txn.PatchAccount_UpdateCredentials) (object.Account, error)
	GetAccount                     func(*http.Request, txn.GetAccount) (object.Account, error)
	GetAccount_Statuses            func(*http.Request, txn.GetAccount_Statuses) ([]object.Status, error)
	GetAccount_Followers           func(*http.Request, txn.GetAccount_Followers) ([]object.Account, error)
	GetAccount_Following           func(*http.Request, txn.GetAccount_Following) ([]object.Account, error)
	GetAccount_FeaturedTags        func(*http.Request, txn.GetAccount_FeaturedTags) ([]object.Tag, error)
	PostAccount_Follow             func(*http.Request, txn.PostAccount_Follow) (object.Relationship, error)
	PostAccount_Unfollow           func(*http.Request, txn.PostAccount_Unfollow) (object.Relationship, error)
	PostAccount_Block              func(*http.Request, txn.PostAccount_Block) (object.Relationship, error)
	PostAccount_Unblock            func(*http.Request, txn.PostAccount_Unblock) (object.Relationship, error)
	PostAccount_Mute               func(*http.Request, txn.PostAccount_Mute) (object.Relationship, error)
	PostAccount_Unmute             func(*http.Request, txn.PostAccount_Unmute) (object.Relationship, error)
	PostAccount_Pin                func(*http.Request, txn.PostAccount_Pin) (object.Relationship, error)
	PostAccount_Unpin              func(*http.Request, txn.PostAccount_Unpin) (object.Relationship, error)
	PostAccount_Note               func(*http.Request, txn.PostAccount_Note) (object.Status, error)
	GetAccount_Relationships       func(*http.Request, txn.GetAccount_Relationships) ([]object.Relationship, error)
	GetAccount_FamiliarFollowers   func(*http.Request, txn.GetAccount_FamiliarFollowers) (object.FamiliarFollowers, error)
	GetAccount_Search              func(*http.Request, txn.GetAccount_Search) ([]object.Account, error)
	GetAccount_Lookup              func(*http.Request, txn.GetAccount_Lookup) (object.Account, error)

	// https://docs.joinmastodon.org/methods/announcements/
	GetAnnouncements            func(*http.Request, txn.GetAnnouncements) ([]object.Announcement, error)
	PostAnnouncement_Dismiss    func(*http.Request, txn.PostAnnouncement_Dismiss) (struct{}, error)
	PutAnnouncement_Reaction    func(*http.Request, txn.PutAnnouncement_Reaction) (struct{}, error)
	DeleteAnnouncement_Reaction func(*http.Request, txn.DeleteAnnouncement_Reaction) (struct{}, error)

	// https://docs.joinmastodon.org/methods/apps/
	PostApplication                  func(*http.Request, txn.PostApplication) (object.Application, error)
	GetApplication_VerifyCredentials func(*http.Request, txn.GetApplication_VerifyCredentials) (object.Application, error)

	// https://docs.joinmastodon.org/methods/blocks/
	GetBlocks func(*http.Request, txn.GetBlocks) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/bookmarks/
	GetBookmarks func(*http.Request, txn.GetBookmarks) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/conversations/
	GetConversations     func(*http.Request, txn.GetConversations) ([]object.Conversation, error)
	DeleteConversation   func(*http.Request, txn.DeleteConversation) (struct{}, error)
	PostConversationRead func(*http.Request, txn.PostConversationRead) (struct{}, error)

	// https://docs.joinmastodon.org/methods/custom_emojis/
	GetCustomEmojis func(*http.Request, txn.GetCustomEmojis) ([]object.CustomEmoji, error)

	// https://docs.joinmastodon.org/methods/directory/
	GetDirectory func(*http.Request, txn.GetDirectory) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/domain_blocks/
	GetDomainBlocks   func(*http.Request, txn.GetDomainBlocks) ([]string, error)
	PostDomainBlock   func(*http.Request, txn.PostDomainBlock) (struct{}, error)
	DeleteDomainBlock func(*http.Request, txn.DeleteDomainBlock) (struct{}, error)

	// https://docs.joinmastodon.org/methods/emails/
	PostEmailConfirmation func(*http.Request, txn.PostEmailConfirmation) (struct{}, error)

	// https://docs.joinmastodon.org/methods/endorsements/
	GetEndorsements func(*http.Request, txn.GetEndorsements) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/favourites/
	GetFavourites func(*http.Request, txn.GetFavourites) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/featured_tags/
	GetFeaturedTags             func(*http.Request, txn.GetFeaturedTags) ([]object.FeaturedTag, error)
	PostFeaturedTag             func(*http.Request, txn.PostFeaturedTag) (object.FeaturedTag, error)
	DeleteFeaturedTag           func(*http.Request, txn.DeleteFeaturedTag) (struct{}, error)
	GetFeaturedTags_Suggestions func(*http.Request, txn.GetFeaturedTags_Suggestions) ([]object.FeaturedTag, error)

	// https://docs.joinmastodon.org/methods/filters/
	GetFilters           func(*http.Request, txn.GetFilters) ([]object.Filter, error)
	GetFilter            func(*http.Request, txn.GetFilter) (object.Filter, error)
	PostFilter           func(*http.Request, txn.PostFilter) (object.Filter, error)
	PutFilter            func(*http.Request, txn.PutFilter) (object.Filter, error)
	DeleteFilter         func(*http.Request, txn.DeleteFilter) (struct{}, error)
	GetFilter_Keywords   func(*http.Request, txn.GetFilter_Keywords) ([]string, error)
	PostFilter_Keyword   func(*http.Request, txn.PostFilter_Keyword) (struct{}, error)
	GetFilter_Keyword    func(*http.Request, txn.GetFilter_Keyword) (struct{}, error)
	PutFilter_Keyword    func(*http.Request, txn.PutFilter_Keyword) (struct{}, error)
	DeleteFilter_Keyword func(*http.Request, txn.DeleteFilter_Keyword) (struct{}, error)
	GetFilter_Statuses   func(*http.Request, txn.GetFilter_Statuses) ([]object.FilterStatus, error)
	PostFilter_Status    func(*http.Request, txn.PostFilter_Status) (object.FilterStatus, error)
	GetFilter_Status     func(*http.Request, txn.GetFilter_Status) (object.FilterStatus, error)
	DeleteFilter_Status  func(*http.Request, txn.DeleteFilter_Status) (struct{}, error)
	GetFilter_V1         func(*http.Request, txn.GetFilter_V1) (object.Filter, error)
	PostFilter_V1        func(*http.Request, txn.PostFilter_V1) (object.Filter, error)
	PutFilter_V1         func(*http.Request, txn.PutFilter_V1) (object.Filter, error)
	DeleteFilter_V1      func(*http.Request, txn.DeleteFilter_V1) (struct{}, error)

	// https://docs.joinmastodon.org/methods/follow_requests/
	GetFollowRequests           func(*http.Request, txn.GetFollowRequests) ([]object.Account, error)
	PostFollowRequest_Authorize func(*http.Request, txn.PostFollowRequest_Authorize) (object.Relationship, error)
	PostFollowRequest_Reject    func(*http.Request, txn.PostFollowRequest_Reject) (object.Relationship, error)

	// https://docs.joinmastodon.org/methods/followed_tags/
	GetFollowedTags func(*http.Request, txn.GetFollowedTags) ([]object.Tag, error)

	// https://docs.joinmastodon.org/methods/instance/
	GetInstance                     func(*http.Request, txn.GetInstance) (object.Instance, error)
	GetInstance_Peers               func(*http.Request, txn.GetInstance_Peers) ([]string, error)
	GetInstance_Activity            func(*http.Request, txn.GetInstance_Activity) (map[string]any, error)
	GetInstance_Rules               func(*http.Request, txn.GetInstance_Rules) ([]object.Rule, error)
	GetInstance_DomainBlocks        func(*http.Request, txn.GetInstance_DomainBlocks) ([]object.DomainBlock, error)
	GetInstance_ExtendedDescription func(*http.Request, txn.GetInstance_ExtendedDescription) (object.ExtendedDescription, error)

	// https://docs.joinmastodon.org/methods/lists/
	GetLists             func(*http.Request, txn.GetLists) ([]object.List, error)
	GetList              func(*http.Request, txn.GetList) (object.List, error)
	PostList             func(*http.Request, txn.PostList) (object.List, error)
	PutList              func(*http.Request, txn.PutList) (object.List, error)
	DeleteList           func(*http.Request, txn.DeleteList) (struct{}, error)
	GetList_Accounts     func(*http.Request, txn.GetList_Accounts) ([]object.Account, error)
	PostList_Accounts    func(*http.Request, txn.PostList_Accounts) (struct{}, error)
	DeleteLisst_Accounts func(*http.Request, txn.DeleteList_Accounts) (struct{}, error)

	// https://docs.joinmastodon.org/methods/markers/
	GetMarkers func(*http.Request, txn.GetMarkers) (object.Marker, error)
	PosMarker  func(*http.Request, txn.PostMarker) (object.Marker, error)

	// https://docs.joinmastodon.org/methods/media/
	PostMedia func(*http.Request, txn.PostMedia) (object.MediaAttachment, error)

	// https://docs.joinmastodon.org/methods/mutes/
	GetMutes func(*http.Request, txn.GetMutes) ([]object.Account, error)

	// https://docs.joinmastodon.org/methods/notifications/
	GetNotifications         func(*http.Request, txn.GetNotifications) ([]object.Notification, error)
	GetNotification          func(*http.Request, txn.GetNotification) (object.Notification, error)
	PostNotifications_Clear  func(*http.Request, txn.PostNotifications_Clear) (object.Notification, error)
	PostNotification_Dismiss func(*http.Request, txn.PostNotification_Dismiss) (object.Notification, error)

	// https://docs.joinmastodon.org/methods/oauth/
	GetOAuth_Authorize func(*http.Request, txn.GetOAuth_Authorize) (struct{}, error)
	PostOAuth_Token    func(*http.Request, txn.PostOAuth_Token) (object.Token, error)
	PostOAuth_Revoke   func(*http.Request, txn.PostOAuth_Revoke) (struct{}, error)

	// https://docs.joinmastodon.org/methods/oembed/
	GetOEmbed func(*http.Request, txn.GetOEmbed) (map[string]any, error)

	// https://docs.joinmastodon.org/methods/polls/
	GetPoll        func(*http.Request, txn.GetPoll) ([]object.Poll, error)
	PostPoll_Votes func(*http.Request, txn.PostPoll_Votes) ([]object.Poll, error)

	// https://docs.joinmastodon.org/methods/preferences/
	GetPreferences func(*http.Request, txn.GetPreferences) (map[string]any, error)

	// https://docs.joinmastodon.org/methods/profile/
	DeleteProfile_Avatar func(*http.Request, txn.DeleteProfile_Avatar) (object.Account, error)
	DeleteProfile_Header func(*http.Request, txn.DeleteProfile_Header) (object.Account, error)

	// https://docs.joinmastodon.org/methods/reports/
	PostReport func(*http.Request, txn.PostReport) (object.Report, error)

	// https://docs.joinmastodon.org/methods/scheduled_statuses/
	GetScheduledStatuses  func(*http.Request, txn.GetScheduledStatuses) ([]object.ScheduledStatus, error)
	GetScheduledStatus    func(*http.Request, txn.GetScheduledStatus) (object.ScheduledStatus, error)
	PutScheduledStatus    func(*http.Request, txn.PutScheduledStatus) (object.ScheduledStatus, error)
	DeleteScheduledStatus func(*http.Request, txn.DeleteScheduledStatus) (struct{}, error)

	// https://docs.joinmastodon.org/methods/search/
	GetSearch func(*http.Request, txn.GetSearch) (object.Search, error)

	// https://docs.joinmastodon.org/methods/statuses/#create
	PostStatus             func(*http.Request, txn.PostStatus) (object.Status, error)
	GetStatus              func(*http.Request, txn.GetStatus) (object.Status, error)
	DeleteStatus           func(*http.Request, txn.DeleteStatus) (struct{}, error)
	GetStatus_Context      func(*http.Request, txn.GetStatus_Context) (object.Context, error)
	PostStatus_Translate   func(*http.Request, txn.PostStatus_Translate) (object.Status, error)
	GetStatus_RebloggedBy  func(*http.Request, txn.GetStatus_RebloggedBy) ([]object.Account, error)
	GetStatus_FavouritedBy func(*http.Request, txn.GetStatus_FavouritedBy) ([]object.Account, error)
	PostStatus_Favourite   func(*http.Request, txn.PostStatus_Favourite) (object.Status, error)
	PostStatus_Unfavourite func(*http.Request, txn.PostStatus_Unfavourite) (object.Status, error)
	PostStatus_Reblog      func(*http.Request, txn.PostStatus_Reblog) (object.Status, error)
	PostStatus_Unreblog    func(*http.Request, txn.PostStatus_Unreblog) (object.Status, error)
	PostStatus_Bookmark    func(*http.Request, txn.PostStatus_Bookmark) (object.Status, error)
	PostStatus_Unbookmark  func(*http.Request, txn.PostStatus_Unbookmark) (object.Status, error)
	PostStatus_Mute        func(*http.Request, txn.PostStatus_Mute) (object.Status, error)
	PostStatus_Unmute      func(*http.Request, txn.PostStatus_Unmute) (object.Status, error)
	PostStatus_Pin         func(*http.Request, txn.PostStatus_Pin) (object.Status, error)
	PostStatus_Unpin       func(*http.Request, txn.PostStatus_Unpin) (object.Status, error)
	PutStatus              func(*http.Request, txn.PutStatus) (object.Status, error)
	GetStatus_History      func(*http.Request, txn.GetStatus_History) ([]object.StatusEdit, error)
	GetStatus_Source       func(*http.Request, txn.GetStatus_Source) (object.StatusSource, error)

	// https://docs.joinmastodon.org/methods/suggestions/
	GetSuggestions   func(*http.Request, txn.GetSuggestions) ([]object.Suggestion, error)
	DeleteSuggestion func(*http.Request, txn.DeleteSuggestion) (struct{}, error)

	// https://docs.joinmastodon.org/methods/tags/
	GetTag           func(*http.Request, txn.GetTag) (object.Tag, error)
	PostTag_Follow   func(*http.Request, txn.PostTag_Follow) (object.Tag, error)
	PostTag_Unfollow func(*http.Request, txn.PostTag_Unfollow) (object.Tag, error)

	// https://docs.joinmastodon.org/methods/timelines/
	GetTimeline_Public  func(*http.Request, txn.GetTimeline_Public) ([]object.Status, error)
	GetTimeline_Hashtag func(*http.Request, txn.GetTimeline_Hashtag) ([]object.Status, error)
	GetTimeline_Home    func(*http.Request, txn.GetTimeline_Home) ([]object.Status, error)
	GetTimeline_List    func(*http.Request, txn.GetTimeline_List) ([]object.Status, error)

	// https://docs.joinmastodon.org/methods/trends/
	GetTrends          func(*http.Request, txn.GetTrends) ([]object.Tag, error)
	GetTrends_Statuses func(*http.Request, txn.GetTrends_Statuses) ([]object.Status, error)
	GetTrends_Links    func(*http.Request, txn.GetTrends_Links) ([]object.PreviewCard, error)
}

func New() API {
	return API{}
}
