package tmsgp

//go:generate msgp
type RedditAccount struct {
	ID   string      `msg:"id"`
	Name string      `msg:"name"`
	Kind string      `msg:"kind"`
	Data AccountData `msg:"data"`
}

type AccountData struct {
	CommentKarma int64  `msg:"comment_karma"`
	HasMail      bool   `msg:"has_mail"`
	HasModMail   bool   `msg:"has_mod_mail"`
	ID           string `msg:"id"`
	InboxCount   int64  `msg:"inbox_count"`
	IsFriend     bool   `msg:"is_friend"`
	IsGold       bool   `msg:"is_gold"`
	LinkKarma    int64  `msg:"link_karma"`
	ModHash      string `msg:"mod_hash"`
	Name         string `msg:"name"`
	Over18       bool   `msg:"over_18"`
}
