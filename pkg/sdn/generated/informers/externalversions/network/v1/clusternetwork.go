// This file was automatically generated by informer-gen

package v1

import (
	network_v1 "github.com/openshift/origin/pkg/sdn/apis/network/v1"
	clientset "github.com/openshift/origin/pkg/sdn/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/sdn/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/sdn/generated/listers/network/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterNetworkInformer provides access to a shared informer and lister for
// ClusterNetworks.
type ClusterNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterNetworkLister
}

type clusterNetworkInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterNetworkInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.NetworkV1().ClusterNetworks().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.NetworkV1().ClusterNetworks().Watch(options)
			},
		},
		&network_v1.ClusterNetwork{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network_v1.ClusterNetwork{}, newClusterNetworkInformer)
}

func (f *clusterNetworkInformer) Lister() v1.ClusterNetworkLister {
	return v1.NewClusterNetworkLister(f.Informer().GetIndexer())
}
