/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	appv1beta1 "sigs.k8s.io/application/pkg/apis/app/v1beta1"
	versioned "sigs.k8s.io/application/pkg/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/application/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "sigs.k8s.io/application/pkg/client/listers/app/v1beta1"
)

// ApplicationInformer provides access to a shared informer and lister for
// Applications.
type ApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ApplicationLister
}

type applicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationInformer constructs a new informer for Application type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationInformer constructs a new informer for Application type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1beta1().Applications(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1beta1().Applications(namespace).Watch(options)
			},
		},
		&appv1beta1.Application{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appv1beta1.Application{}, f.defaultInformer)
}

func (f *applicationInformer) Lister() v1beta1.ApplicationLister {
	return v1beta1.NewApplicationLister(f.Informer().GetIndexer())
}
