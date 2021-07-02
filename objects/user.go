package objects

type (
	User struct {
		ID                             string   `json:"id"`
		Username                       string   `json:"username"`
		DisplayName                    string   `json:"displayName"`
		UserIcon                       string   `json:"userIcon"`
		Bio                            string   `json:"bio"`
		BioLinks                       []string `json:"bioLinks"`
		ProfilePicOverride             string   `json:"profilePicOverride"`
		StatusDescription              string   `json:"statusDescription"`
		CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
		FallbackAvatar                 string   `json:"fallbackAvatar"`
		State                          string   `json:"state"`
		Tags                           []string `json:"tags"`
		DeveloperType                  string   `json:"developerType"`
		LastLogin                      string   `json:"last_login"`
		LastPlatform                   string   `json:"last_platform"`
		AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
		Status                         string   `json:"status"`
		DateJoined                     string   `json:"date_joined"`
		IsFriend                       bool     `json:"isFriend"`
		FriendKey                      string   `json:"friendKey"`
		WorldID                        string   `json:"worldId"`
		InstanceID                     string   `json:"instanceId"`
		Location                       string   `json:"location"`
	}

	CurrentUser struct {
		ID                             string            `json:"id"`
		Username                       string            `json:"username"`
		DisplayName                    string            `json:"displayName"`
		UserIcon                       string            `json:"userIcon"`
		Bio                            string            `json:"bio"`
		BioLinks                       []string          `json:"bioLinks"`
		ProfilePicOverride             string            `json:"profilePicOverride"`
		StatusDescription              string            `json:"statusDescription"`
		PastDisplayNames               []PastDisplayName `json:"pastDisplayNames"`
		HasEmail                       bool              `json:"hasEmail"`
		HasPendingEmail                bool              `json:"hasPendingEmail"`
		ObfuscatedEmail                string            `json:"obfuscatedEmail"`
		ObfuscatedPendingEmail         string            `json:"obfuscatedPendingEmail"`
		EmailVerified                  bool              `json:"emailVerified"`
		HasBirthday                    bool              `json:"hasBirthday"`
		Unsubscribe                    bool              `json:"unsubscribe"`
		StatusHistory                  []string          `json:"statusHistory"`
		StatusFirstTime                bool              `json:"statusFirstTime"`
		Friends                        []string          `json:"friends"`
		FriendGroupNames               []string          `json:"friendGroupNames"`
		CurrentAvatarImageURL          string            `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string            `json:"currentAvatarThumbnailImageUrl"`
		FallbackAvatar                 string            `json:"fallbackAvatar"`
		CurrentAvatar                  string            `json:"currentAvatar"`
		CurrentAvatarAssetURL          string            `json:"currentAvatarAssetURL"`
		AccountDeletionDate            string            `json:"accountDeletionDate,omitempty"`
		AcceptedTOSVersion             int               `json:"acceptedTOSVersion"`
		SteamID                        string            `json:"steamId"`
		SteamDetails                   SteamDetails      `json:"steamDetails"`
		OculusID                       string            `json:"oculusId"`
		HasLoggedInFromClient          bool              `json:"hasLoggedInFromClient"`
		HomeLocation                   string            `json:"homeLocation"`
		TwoFactorAuthEnabled           bool              `json:"twoFactorAuthEnabled"`
		State                          string            `json:"state"`
		Tags                           []string          `json:"tags"`
		DeveloperType                  string            `json:"developerType"`
		LastLogin                      string            `json:"last_login"`
		LastPlatform                   string            `json:"last_platform"`
		AllowAvatarCopying             bool              `json:"allowAvatarCopying"`
		Status                         string            `json:"status"`
		DateJoined                     string            `json:"date_joined"`
		IsFriend                       bool              `json:"isFriend"`
		FriendKey                      string            `json:"friendKey"`
		OnlineFriends                  []string          `json:"onlineFriends"`
		ActiveFriends                  []string          `json:"activeFriends"`
		OfflineFriends                 []string          `json:"offlineFriends"`
		RequiresTwoFactorAuth          []string          `json:"requiresTwoFactorAuth,omitempty"`
	}

	LimitedUser struct {
		ID                             string   `json:"id"`
		Username                       string   `json:"username"`
		DisplayName                    string   `json:"displayName"`
		UserIcon                       string   `json:"userIcon"`
		Bio                            string   `json:"bio"`
		ProfilePicOverride             string   `json:"profilePicOverride"`
		StatusDescription              string   `json:"statusDescription"`
		CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
		FallbackAvatar                 string   `json:"fallbackAvatar"`
		Tags                           []string `json:"tags"`
		DeveloperType                  string   `json:"developerType"`
		LastPlatform                   string   `json:"last_platform"`
		Status                         string   `json:"status"`
		IsFriend                       bool     `json:"isFriend"`
	}

	PastDisplayName struct {
		DisplayName string `json:"displayName"`
		UpdatedAt   string `json:"updated_at"`
	}

	SteamDetails struct {
		SteamID                  string `json:"steamid"`
		CommunityVisibilityState int    `json:"communityvisibilitystate"`
		ProfileState             int    `json:"profilestate"`
		PersonaName              string `json:"personaname"`
		ProfileURL               string `json:"profileurl"`
		Avatar                   string `json:"avatar"`
		AvatarMedium             string `json:"avatarmedium"`
		AvatarFull               string `json:"avatarfull"`
		PersonaState             int    `json:"personastate"`
		RealName                 string `json:"realname"`
		PrimaryClanID            string `json:"primaryclanid"`
		TimeCreated              int    `json:"timecreated"`
		PersonaStateFlags        int    `json:"personastateflags"`
		AvatarHash               string `json:"avatarhash"`
		CommentPermission        int    `json:"commentpermission"`
		GameExtraInfo            string `json:"gameextrainfo"`
		GameID                   string `json:"gameid"`
	}
)

const (
	// DeveloperTypeNone Not a developer
	DeveloperTypeNone = "none"
	// DeveloperTypeTrusted Unknown
	DeveloperTypeTrusted = "trusted"
	// DeveloperTypeInternal Is a VRChat developer
	DeveloperTypeInternal = "internal"
	// DeveloperTypeModerator Is a VRChat moderator
	DeveloperTypeModerator = "moderator"

	// StateOnline User is online in VRChat
	StateOnline = "online"
	// StateActive User is online, but not in VRChat
	StateActive = "active"
	// StateOffline User is offline
	StateOffline = "offline"

	// StatusActive User can see requests
	StatusActive = "active"
	// StatusJoinMe User auto-accepts requests
	StatusJoinMe = "join me"
	// StatusAskMe People can't see user's location, but they can see requests
	StatusAskMe = "ask me"
	// StatusBusy User auto-declines requests
	StatusBusy = "busy"
	// StatusOffline User is offline
	StatusOffline = "offline"
)

func (obj CurrentUser) Requires2FA() bool {
	return obj.RequiresTwoFactorAuth != nil && len(obj.RequiresTwoFactorAuth) > 0
}
