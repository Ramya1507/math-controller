/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	arithmeticopv1alpha1 "math-controller/pkg/apis/arithmeticop/v1alpha1"
	versioned "math-controller/pkg/client/clientset/versioned"
	internalinterfaces "math-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "math-controller/pkg/client/listers/arithmeticop/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MathResourceInformer provides access to a shared informer and lister for
// MathResources.
type MathResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MathResourceLister
}

type mathResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMathResourceInformer constructs a new informer for MathResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMathResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMathResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMathResourceInformer constructs a new informer for MathResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMathResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MathsV1alpha1().MathResources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MathsV1alpha1().MathResources(namespace).Watch(context.TODO(), options)
			},
		},
		&arithmeticopv1alpha1.MathResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *mathResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMathResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mathResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&arithmeticopv1alpha1.MathResource{}, f.defaultInformer)
}

func (f *mathResourceInformer) Lister() v1alpha1.MathResourceLister {
	return v1alpha1.NewMathResourceLister(f.Informer().GetIndexer())
}
