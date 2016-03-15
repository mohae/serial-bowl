package shared

type ShRedditAccount struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Kind string      `json:"kind"`
	Data AccountData `json:"data"`
}

type AccountData struct {
	CommentKarma int64  `json:"comment_karma"`
	HasMail      bool   `json:"has_mail"`
	HasModMail   bool   `json:"has_mod_mail"`
	ID           string `json:"id"`
	InboxCount   int64  `json:"inbox_count"`
	IsFriend     bool   `json:"is_friend"`
	IsGold       bool   `json:"is_gold"`
	LinkKarma    int64  `json:"link_karma"`
	ModHash      string `json:"mod_hash"`
	Name         string `json:"name"`
	Over18       bool   `json:"over_18"`
}
