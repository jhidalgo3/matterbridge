// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ThumbnailSet undocumented
type ThumbnailSet struct {
	// Entity is the base model of ThumbnailSet
	Entity
	// Large undocumented
	Large *Thumbnail `json:"large,omitempty"`
	// Medium undocumented
	Medium *Thumbnail `json:"medium,omitempty"`
	// Small undocumented
	Small *Thumbnail `json:"small,omitempty"`
	// Source undocumented
	Source *Thumbnail `json:"source,omitempty"`
}
