package objects

import (
	"time"
)

type RemoteConfig struct {
	VoiceEnableDegradation      bool    `json:"VoiceEnableDegradation"`
	VoiceEnableReceiverLimiting bool    `json:"VoiceEnableReceiverLimiting"`
	AccountPosition             float64 `json:"accountPosition"`
	Address                     string  `json:"address"`
	Announcements               []struct {
		Name string `json:"name"`
		Text string `json:"text"`
	} `json:"announcements"`
	APIKey                           string    `json:"apiKey"`
	AppName                          string    `json:"appName"`
	BonesSocketPipeline              float64   `json:"bonesSocketPipeline"`
	BuildVersionTag                  string    `json:"buildVersionTag"`
	ClientAPIKey                     string    `json:"clientApiKey"`
	ClientBPSCeiling                 int       `json:"clientBPSCeiling"`
	ClientDisconnectTimeout          int       `json:"clientDisconnectTimeout"`
	ClientReservedPlayerBPS          int       `json:"clientReservedPlayerBPS"`
	ClientSentCountAllowance         int       `json:"clientSentCountAllowance"`
	ClipDefaultBuild                 bool      `json:"clipDefaultBuild"`
	ContactEmail                     string    `json:"contactEmail"`
	CopyrightEmail                   string    `json:"copyrightEmail"`
	CurrentTOSVersion                int       `json:"currentTOSVersion"`
	DefaultAvatar                    string    `json:"defaultAvatar"`
	DeploymentGroup                  string    `json:"deploymentGroup"`
	DevAppVersionStandalone          string    `json:"devAppVersionStandalone"`
	DevDownloadLinkWindows           string    `json:"devDownloadLinkWindows"`
	DevSdkURL                        string    `json:"devSdkUrl"`
	DevSdkVersion                    string    `json:"devSdkVersion"`
	DevServerVersionStandalone       string    `json:"devServerVersionStandalone"`
	DisCountdown                     time.Time `json:"dis-countdown"`
	DisableAvatarCopying             bool      `json:"disableAvatarCopying"`
	DisableAvatarGating              bool      `json:"disableAvatarGating"`
	DisableCommunityLabs             bool      `json:"disableCommunityLabs"`
	DisableCommunityLabsPromotion    bool      `json:"disableCommunityLabsPromotion"`
	DisableEmail                     bool      `json:"disableEmail"`
	DisableEventStream               bool      `json:"disableEventStream"`
	DisableFeedbackGating            bool      `json:"disableFeedbackGating"`
	DisableHello                     bool      `json:"disableHello"`
	DisableRegistration              bool      `json:"disableRegistration"`
	DisableSteamNetworking           bool      `json:"disableSteamNetworking"`
	DisableTwoFactorAuth             bool      `json:"disableTwoFactorAuth"`
	DisableUdon                      bool      `json:"disableUdon"`
	DisableUpgradeAccount            bool      `json:"disableUpgradeAccount"`
	DistanceCrossAccountScrollMobile int       `json:"distanceCrossAccountScrollMobile"`
	DownloadLinkWindows              string    `json:"downloadLinkWindows"`
	DownloadUrls                     struct {
		Sdk2        string `json:"sdk2"`
		Sdk3Worlds  string `json:"sdk3-worlds"`
		Sdk3Avatars string `json:"sdk3-avatars"`
	} `json:"downloadUrls"`
	DynamicWorldRows []struct {
		Name          string `json:"name"`
		SortHeading   string `json:"sortHeading"`
		SortOwnership string `json:"sortOwnership"`
		SortOrder     string `json:"sortOrder"`
		Platform      string `json:"platform"`
		Index         int    `json:"index"`
		Tag           string `json:"tag,omitempty"`
	} `json:"dynamicWorldRows"`
	EmbedSegmentNameNuisanceHidden int `json:"embedSegmentNameNuisanceHidden"`
	Events                         struct {
		DistanceClose             int `json:"distanceClose"`
		DistanceFactor            int `json:"distanceFactor"`
		DistanceFar               int `json:"distanceFar"`
		GroupDistance             int `json:"groupDistance"`
		MaximumBunchSize          int `json:"maximumBunchSize"`
		NotVisibleFactor          int `json:"notVisibleFactor"`
		PlayerOrderBucketSize     int `json:"playerOrderBucketSize"`
		PlayerOrderFactor         int `json:"playerOrderFactor"`
		SlowUpdateFactorThreshold int `json:"slowUpdateFactorThreshold"`
		ViewSegmentLength         int `json:"viewSegmentLength"`
	} `json:"events"`
	FavoriteSystem                  int     `json:"favoriteSystem"`
	FileFarLengthLocalizationFormat float64 `json:"fileFarLengthLocalizationFormat"`
	FileOnly                        struct {
	} `json:"fileOnly"`
	GearDemoRoomID               string `json:"gearDemoRoomId"`
	GeneratorHiddenPacketHandler struct {
	} `json:"generatorHiddenPacketHandler"`
	HeadpatPermanentPipeline                      bool   `json:"headpatPermanentPipeline"`
	HomeWorldID                                   string `json:"homeWorldId"`
	HomepageRedirectTarget                        string `json:"homepageRedirectTarget"`
	HubWorldID                                    string `json:"hubWorldId"`
	JobsEmail                                     string `json:"jobsEmail"`
	MessageOfTheDay                               string `json:"messageOfTheDay"`
	ModerationEmail                               string `json:"moderationEmail"`
	ModerationQueryPeriod                         int    `json:"moderationQueryPeriod"`
	NotAllowedToSelectAvatarInPrivateWorldMessage string `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	PacketAlignmentSdkNetworkBurrito              bool   `json:"packetAlignmentSdkNetworkBurrito"`
	Plugin                                        string `json:"plugin"`
	ReleaseAppVersionStandalone                   string `json:"releaseAppVersionStandalone"`
	ReleaseSdkURL                                 string `json:"releaseSdkUrl"`
	ReleaseSdkVersion                             string `json:"releaseSdkVersion"`
	ReleaseServerVersionStandalone                string `json:"releaseServerVersionStandalone"`
	RoomAlignmentKey                              int    `json:"roomAlignmentKey"`
	SdkDeveloperFaqURL                            string `json:"sdkDeveloperFaqUrl"`
	SdkDiscordURL                                 string `json:"sdkDiscordUrl"`
	SdkNotAllowedToPublishMessage                 string `json:"sdkNotAllowedToPublishMessage"`
	SdkUnityVersion                               string `json:"sdkUnityVersion"`
	ServerName                                    string `json:"serverName"`
	//StackAccountLocalization                      []interface{} `json:"stackAccountLocalization"`
	SupportEmail            string   `json:"supportEmail"`
	ThresholdPluginEmbed    int      `json:"thresholdPluginEmbed"`
	TimeOutWorldID          string   `json:"timeOutWorldId"`
	TutorialWorldID         string   `json:"tutorialWorldId"`
	UpdateRateMsMaximum     int      `json:"updateRateMsMaximum"`
	UpdateRateMsMinimum     int      `json:"updateRateMsMinimum"`
	UpdateRateMsNormal      int      `json:"updateRateMsNormal"`
	UpdateRateMsUdonManual  int      `json:"updateRateMsUdonManual"`
	UploadAnalysisPercent   int      `json:"uploadAnalysisPercent"`
	URLList                 []string `json:"urlList"`
	UseReliableUDPForVoice  bool     `json:"useReliableUdpForVoice"`
	UserUpdatePeriod        int      `json:"userUpdatePeriod"`
	UserVerificationDelay   int      `json:"userVerificationDelay"`
	UserVerificationRetry   int      `json:"userVerificationRetry"`
	UserVerificationTimeout int      `json:"userVerificationTimeout"`
	ViveWindowsURL          string   `json:"viveWindowsUrl"`
	WhiteListedAssetUrls    []string `json:"whiteListedAssetUrls"`
	WorldUpdatePeriod       int      `json:"worldUpdatePeriod"`
	YoutubeDLHash           string   `json:"youtubedl-hash"`
	YoutubeDLVersion        string   `json:"youtubedl-version"`
}
