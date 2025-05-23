// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterAccessesGetter has a method to return a ClusterAccessInterface.
// A group's client should implement this interface.
type ClusterAccessesGetter interface {
	ClusterAccesses() ClusterAccessInterface
}

// ClusterAccessInterface has methods to work with ClusterAccess resources.
type ClusterAccessInterface interface {
	Create(ctx context.Context, clusterAccess *storagev1.ClusterAccess, opts metav1.CreateOptions) (*storagev1.ClusterAccess, error)
	Update(ctx context.Context, clusterAccess *storagev1.ClusterAccess, opts metav1.UpdateOptions) (*storagev1.ClusterAccess, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterAccess *storagev1.ClusterAccess, opts metav1.UpdateOptions) (*storagev1.ClusterAccess, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.ClusterAccess, error)
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1.ClusterAccessList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *storagev1.ClusterAccess, err error)
	ClusterAccessExpansion
}

// clusterAccesses implements ClusterAccessInterface
type clusterAccesses struct {
	*gentype.ClientWithList[*storagev1.ClusterAccess, *storagev1.ClusterAccessList]
}

// newClusterAccesses returns a ClusterAccesses
func newClusterAccesses(c *StorageV1Client) *clusterAccesses {
	return &clusterAccesses{
		gentype.NewClientWithList[*storagev1.ClusterAccess, *storagev1.ClusterAccessList](
			"clusteraccesses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *storagev1.ClusterAccess { return &storagev1.ClusterAccess{} },
			func() *storagev1.ClusterAccessList { return &storagev1.ClusterAccessList{} },
		),
	}
}
