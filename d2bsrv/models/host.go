package models

import "gopkg.in/mgo.v2/bson"

type Host struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Host       string        `json:"host" bson:"host"`
	Build      bool          `json:"build" bson:"build"`
	Debug      bool          `json:"debug" bson:"debug"`
	GPT        bool          `json:"gpt" bson:"gpt"`
	TagID      bson.ObjectId `json:"tagId" bson:"tagId"`
	TagRef     string        `json:"tagRef" bson:"tagRef"`
	Tag        *Tag          `json:"tag,omitempty"`
	Image      *Image        `json:"image,omitempty"`
	BootTag    *Tag          `json:"bootTag,omitempty"`
	BootImage  *Image        `json:"bootImage,omitempty"`
	KOpts      string        `json:"kOpts" bson:"kOpts"`
	TenantID   bson.ObjectId `json:"tenantId" bson:"tenantId"`
	TenantRef  string        `json:"tenantRef" bson:"tenantRef"`
	Tenant     *Tenant       `json:"tenant,omitempty"`
	Labels     []string      `json:"labels" bson:"labels"`
	SiteID     bson.ObjectId `json:"siteId" bson:"siteId"`
	SiteRef    string        `json:"siteRef" bson:"siteRef"`
	Site       *Site         `json:"site,omitempty"`
	Interfaces *[]Interface  `json:"interfaces,omitempty"`
}
