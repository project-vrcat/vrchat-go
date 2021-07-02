package objects

import (
	"time"
)

type (
	World struct {
		ID                string          `json:"id"`
		Name              string          `json:"name"`
		Description       string          `json:"description"`
		Featured          bool            `json:"featured"`
		AuthorID          string          `json:"authorId"`
		AuthorName        string          `json:"authorName"`
		Capacity          int             `json:"capacity"`
		Tags              []string        `json:"tags"`
		ReleaseStatus     string          `json:"releaseStatus"`
		ImageURL          string          `json:"imageUrl"`
		ThumbnailImageURL string          `json:"thumbnailImageUrl"`
		AssetURL          string          `json:"assetUrl"`
		Namespace         string          `json:"namespace"`
		UnityPackages     []UnityPackages `json:"unityPackages"`
		Version           int             `json:"version"`
		Organization      string          `json:"organization"`
		//PreviewYoutubeID    interface{}   `json:"previewYoutubeId"`
		Favorites           int       `json:"favorites"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		PublicationDate     time.Time `json:"publicationDate"`
		LabsPublicationDate string    `json:"labsPublicationDate"`
		Visits              int       `json:"visits"`
		Popularity          int       `json:"popularity"`
		Heat                int       `json:"heat"`
		PublicOccupants     int       `json:"publicOccupants"`
		PrivateOccupants    int       `json:"privateOccupants"`
		Occupants           int       `json:"occupants"`
		//Instances           []interface{} `json:"instances"`
	}

	LimitedWorld struct {
		ID                  string          `json:"id"`
		Name                string          `json:"name"`
		AuthorID            string          `json:"authorId"`
		AuthorName          string          `json:"authorName"`
		Capacity            int             `json:"capacity"`
		ImageURL            string          `json:"imageUrl"`
		ThumbnailImageURL   string          `json:"thumbnailImageUrl"`
		ReleaseStatus       string          `json:"releaseStatus"`
		Organization        string          `json:"organization"`
		Tags                []string        `json:"tags"`
		Favorites           int             `json:"favorites"`
		CreatedAt           time.Time       `json:"created_at"`
		UpdatedAt           time.Time       `json:"updated_at"`
		PublicationDate     time.Time       `json:"publicationDate"`
		LabsPublicationDate time.Time       `json:"labsPublicationDate"`
		UnityPackages       []UnityPackages `json:"unityPackages"`
		Popularity          int             `json:"popularity"`
		Heat                int             `json:"heat"`
		Occupants           int             `json:"occupants"`
	}

	UnityPackages struct {
		ID              string    `json:"id"`
		AssetURL        string    `json:"assetUrl"`
		PluginURL       string    `json:"pluginUrl"`
		UnityVersion    string    `json:"unityVersion"`
		UnitySortNumber int64     `json:"unitySortNumber"`
		AssetVersion    int       `json:"assetVersion"`
		Platform        string    `json:"platform"`
		CreatedAt       time.Time `json:"created_at"`
	}
)
