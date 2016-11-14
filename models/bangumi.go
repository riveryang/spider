package models

type Bangumi struct {
	Id         uint64         `json:"id,omitempty"`
	Title      string         `json:"title,omitempty" orm:"size(128);null"`
	AliasTitle string         `json:"aliasTitle,omitempty" orm:"size(128);null"`
	Type       string         `json:"type,omitempty" orm:"size(64);null"`
	Score      string         `json:"score,omitempty" orm:"size(8);null"`
	ScoreDesc  string         `json:"scoreDesc,omitempty" orm:"size(32);null"`
	Rank       string         `json:"rank,omitempty" orm:"size(16);null"`
	Image      string         `json:"image,omitempty" orm:"size(256);null"`
	Info       []*BangumiInfo `json:"info,omitempty" orm:"reverse(many)"`
	Summary    string         `json:"summary,omitempty" orm:"type(text);null"`
	Characters []*Character   `json:"characters,omitempty" orm:"reverse(many)"`
}

type BangumiInfo struct {
	Id     uint64 `json:"id,omitempty" orm:"auto"`
	Tag    string `json:"tag,omitempty" orm:"size(128);null"`
	Values string `json:"values,omitempty" orm:"size(2048);null"`

	Bangumi *Bangumi `json:"-" orm:"rel(fk)"`
}

type Character struct {
	Id        uint64 `json:"id,omitempty" orm:"auto"`
	Name      string `json:"name,omitempty" orm:"size(512);null"`
	AliasName string `json:"aliasName,omitempty" orm:"size(512);null"`
	BadgeJob  string `json:"badgeJob,omitempty" orm:"size(64);null"`
	Seiyuu    string `json:"seiyuu,omitempty" orm:"size(1024);null"`

	Bangumi *Bangumi `json:"-" orm:"rel(fk)"`
}
