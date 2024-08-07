package unsplash

import "time"

type UnSplash struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		Id               string `json:"id"`
		Slug             string `json:"slug"`
		AlternativeSlugs struct {
			En string `json:"en"`
			Es string `json:"es"`
			Ja string `json:"ja"`
			Fr string `json:"fr"`
			It string `json:"it"`
			Ko string `json:"ko"`
			De string `json:"de"`
			Pt string `json:"pt"`
		} `json:"alternative_slugs"`
		CreatedAt      time.Time     `json:"created_at"`
		UpdatedAt      time.Time     `json:"updated_at"`
		PromotedAt     time.Time     `json:"promoted_at"`
		Width          int           `json:"width"`
		Height         int           `json:"height"`
		Color          string        `json:"color"`
		BlurHash       string        `json:"blur_hash"`
		Description    string        `json:"description"`
		AltDescription string        `json:"alt_description"`
		Breadcrumbs    []interface{} `json:"breadcrumbs"`
		Urls           struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
			SmallS3 string `json:"small_s3"`
		} `json:"urls"`
		Links struct {
			Self             string `json:"self"`
			Html             string `json:"html"`
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
		} `json:"links"`
		Likes                  int           `json:"likes"`
		LikedByUser            bool          `json:"liked_by_user"`
		CurrentUserCollections []interface{} `json:"current_user_collections"`
		Sponsorship            interface{}   `json:"sponsorship"`
		TopicSubmissions       struct {
			FoodDrink struct {
				Status     string    `json:"status"`
				ApprovedOn time.Time `json:"approved_on"`
			} `json:"food-drink"`
		} `json:"topic_submissions"`
		AssetType string `json:"asset_type"`
		User      struct {
			Id              string    `json:"id"`
			UpdatedAt       time.Time `json:"updated_at"`
			Username        string    `json:"username"`
			Name            string    `json:"name"`
			FirstName       string    `json:"first_name"`
			LastName        string    `json:"last_name"`
			TwitterUsername string    `json:"twitter_username"`
			PortfolioUrl    string    `json:"portfolio_url"`
			Bio             string    `json:"bio"`
			Location        string    `json:"location"`
			Links           struct {
				Self      string `json:"self"`
				Html      string `json:"html"`
				Photos    string `json:"photos"`
				Likes     string `json:"likes"`
				Portfolio string `json:"portfolio"`
				Following string `json:"following"`
				Followers string `json:"followers"`
			} `json:"links"`
			ProfileImage struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Large  string `json:"large"`
			} `json:"profile_image"`
			InstagramUsername          string `json:"instagram_username"`
			TotalCollections           int    `json:"total_collections"`
			TotalLikes                 int    `json:"total_likes"`
			TotalPhotos                int    `json:"total_photos"`
			TotalPromotedPhotos        int    `json:"total_promoted_photos"`
			TotalIllustrations         int    `json:"total_illustrations"`
			TotalPromotedIllustrations int    `json:"total_promoted_illustrations"`
			AcceptedTos                bool   `json:"accepted_tos"`
			ForHire                    bool   `json:"for_hire"`
			Social                     struct {
				InstagramUsername string      `json:"instagram_username"`
				PortfolioUrl      string      `json:"portfolio_url"`
				TwitterUsername   string      `json:"twitter_username"`
				PaypalEmail       interface{} `json:"paypal_email"`
			} `json:"social"`
		} `json:"user"`
		Tags []struct {
			Type   string `json:"type"`
			Title  string `json:"title"`
			Source struct {
				Ancestry struct {
					Type struct {
						Slug       string `json:"slug"`
						PrettySlug string `json:"pretty_slug"`
					} `json:"type"`
					Category struct {
						Slug       string `json:"slug"`
						PrettySlug string `json:"pretty_slug"`
					} `json:"category"`
				} `json:"ancestry"`
				Title           string `json:"title"`
				Subtitle        string `json:"subtitle"`
				Description     string `json:"description"`
				MetaTitle       string `json:"meta_title"`
				MetaDescription string `json:"meta_description"`
				CoverPhoto      struct {
					Id               string `json:"id"`
					Slug             string `json:"slug"`
					AlternativeSlugs struct {
						En string `json:"en"`
						Es string `json:"es"`
						Ja string `json:"ja"`
						Fr string `json:"fr"`
						It string `json:"it"`
						Ko string `json:"ko"`
						De string `json:"de"`
						Pt string `json:"pt"`
					} `json:"alternative_slugs"`
					CreatedAt      time.Time     `json:"created_at"`
					UpdatedAt      time.Time     `json:"updated_at"`
					PromotedAt     time.Time     `json:"promoted_at"`
					Width          int           `json:"width"`
					Height         int           `json:"height"`
					Color          string        `json:"color"`
					BlurHash       string        `json:"blur_hash"`
					Description    string        `json:"description"`
					AltDescription string        `json:"alt_description"`
					Breadcrumbs    []interface{} `json:"breadcrumbs"`
					Urls           struct {
						Raw     string `json:"raw"`
						Full    string `json:"full"`
						Regular string `json:"regular"`
						Small   string `json:"small"`
						Thumb   string `json:"thumb"`
						SmallS3 string `json:"small_s3"`
					} `json:"urls"`
					Links struct {
						Self             string `json:"self"`
						Html             string `json:"html"`
						Download         string `json:"download"`
						DownloadLocation string `json:"download_location"`
					} `json:"links"`
					Likes                  int           `json:"likes"`
					LikedByUser            bool          `json:"liked_by_user"`
					CurrentUserCollections []interface{} `json:"current_user_collections"`
					Sponsorship            interface{}   `json:"sponsorship"`
					TopicSubmissions       struct {
						Health struct {
							Status     string    `json:"status"`
							ApprovedOn time.Time `json:"approved_on"`
						} `json:"health"`
					} `json:"topic_submissions"`
					AssetType string `json:"asset_type"`
					User      struct {
						Id              string      `json:"id"`
						UpdatedAt       time.Time   `json:"updated_at"`
						Username        string      `json:"username"`
						Name            string      `json:"name"`
						FirstName       string      `json:"first_name"`
						LastName        string      `json:"last_name"`
						TwitterUsername interface{} `json:"twitter_username"`
						PortfolioUrl    string      `json:"portfolio_url"`
						Bio             string      `json:"bio"`
						Location        string      `json:"location"`
						Links           struct {
							Self      string `json:"self"`
							Html      string `json:"html"`
							Photos    string `json:"photos"`
							Likes     string `json:"likes"`
							Portfolio string `json:"portfolio"`
							Following string `json:"following"`
							Followers string `json:"followers"`
						} `json:"links"`
						ProfileImage struct {
							Small  string `json:"small"`
							Medium string `json:"medium"`
							Large  string `json:"large"`
						} `json:"profile_image"`
						InstagramUsername          string `json:"instagram_username"`
						TotalCollections           int    `json:"total_collections"`
						TotalLikes                 int    `json:"total_likes"`
						TotalPhotos                int    `json:"total_photos"`
						TotalPromotedPhotos        int    `json:"total_promoted_photos"`
						TotalIllustrations         int    `json:"total_illustrations"`
						TotalPromotedIllustrations int    `json:"total_promoted_illustrations"`
						AcceptedTos                bool   `json:"accepted_tos"`
						ForHire                    bool   `json:"for_hire"`
						Social                     struct {
							InstagramUsername string      `json:"instagram_username"`
							PortfolioUrl      string      `json:"portfolio_url"`
							TwitterUsername   interface{} `json:"twitter_username"`
							PaypalEmail       interface{} `json:"paypal_email"`
						} `json:"social"`
					} `json:"user"`
				} `json:"cover_photo"`
			} `json:"source,omitempty"`
		} `json:"tags"`
	} `json:"results"`
}
