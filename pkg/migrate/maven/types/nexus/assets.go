package nexus

import (
	"reflect"
	"time"
)

type GetAssetsResponse struct {
	Items             []Item `json:"items"`
	ContinuationToken string `json:"continuationToken"`
}

type Item struct {
	DownloadUrl string `json:"downloadUrl"`
	Path        string `json:"path"`
	Id          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
	Checksum    struct {
		Sha1   string `json:"sha1"`
		Sha512 string `json:"sha512"`
		Sha256 string `json:"sha256"`
		Md5    string `json:"md5"`
	} `json:"checksum"`
	ContentType    string     `json:"contentType"`
	LastModified   time.Time  `json:"lastModified"`
	BlobCreated    time.Time  `json:"blobCreated"`
	LastDownloaded *time.Time `json:"lastDownloaded"`
	Maven2         Maven2     `json:"maven2,omitempty"`
}

type Maven2 struct {
	Extension  string `json:"extension"`
	GroupId    string `json:"groupId"`
	ArtifactId string `json:"artifactId"`
	Version    string `json:"version"`
}

func (m Maven2) IsEmpty() bool {
	return reflect.DeepEqual(m, Maven2{})
}
