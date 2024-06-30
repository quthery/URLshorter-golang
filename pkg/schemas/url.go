package url

type URL struct {
	Id    int    `json:"-"`
	Alias string `json:"alias" binding:"required"`
	URL   string `json:"url" binding:"required"`
}

type AliasGet struct {
	Alias string `form:"alias" json:"alias" binding:"required"`
}
