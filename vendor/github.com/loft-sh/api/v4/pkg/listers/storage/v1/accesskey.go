// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AccessKeyLister helps list AccessKeys.
// All objects returned here must be treated as read-only.
type AccessKeyLister interface {
	// List lists all AccessKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1.AccessKey, err error)
	// Get retrieves the AccessKey from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storagev1.AccessKey, error)
	AccessKeyListerExpansion
}

// accessKeyLister implements the AccessKeyLister interface.
type accessKeyLister struct {
	listers.ResourceIndexer[*storagev1.AccessKey]
}

// NewAccessKeyLister returns a new AccessKeyLister.
func NewAccessKeyLister(indexer cache.Indexer) AccessKeyLister {
	return &accessKeyLister{listers.New[*storagev1.AccessKey](indexer, storagev1.Resource("accesskey"))}
}
