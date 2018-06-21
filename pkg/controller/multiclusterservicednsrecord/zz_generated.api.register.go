/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package multiclusterservicednsrecord

import (
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"github.com/kubernetes-sigs/federation-v2/pkg/controller/sharedinformers"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// MultiClusterServiceDNSRecordController implements the controller.MultiClusterServiceDNSRecordController interface
type MultiClusterServiceDNSRecordController struct {
	queue *controller.QueueWorker

	// Handles messages
	controller *MultiClusterServiceDNSRecordControllerImpl

	Name string

	BeforeReconcile func(key string)
	AfterReconcile  func(key string, err error)

	Informers *sharedinformers.SharedInformers
}

// NewController returns a new MultiClusterServiceDNSRecordController for responding to MultiClusterServiceDNSRecord events
func NewMultiClusterServiceDNSRecordController(config *rest.Config, si *sharedinformers.SharedInformers) *MultiClusterServiceDNSRecordController {
	q := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "MultiClusterServiceDNSRecord")

	queue := &controller.QueueWorker{q, 10, "MultiClusterServiceDNSRecord", nil}
	c := &MultiClusterServiceDNSRecordController{queue, nil, "MultiClusterServiceDNSRecord", nil, nil, si}

	// For non-generated code to add events
	uc := &MultiClusterServiceDNSRecordControllerImpl{}
	var ci sharedinformers.Controller = uc

	// Call the Init method that is implemented.
	// Support multiple Init methods for backwards compatibility
	if i, ok := ci.(sharedinformers.LegacyControllerInit); ok {
		i.Init(config, si, c.LookupAndReconcile)
	} else if i, ok := ci.(sharedinformers.ControllerInit); ok {
		i.Init(&sharedinformers.ControllerInitArgumentsImpl{si, config, c.LookupAndReconcile})
	}

	c.controller = uc

	queue.Reconcile = c.reconcile
	if c.Informers.WorkerQueues == nil {
		c.Informers.WorkerQueues = map[string]*controller.QueueWorker{}
	}
	c.Informers.WorkerQueues["MultiClusterServiceDNSRecord"] = queue
	si.Factory.Multiclusterdns().V1alpha1().MultiClusterServiceDNSRecords().Informer().
		AddEventHandler(&controller.QueueingEventHandler{q, nil, false})
	return c
}

func (c *MultiClusterServiceDNSRecordController) GetName() string {
	return c.Name
}

func (c *MultiClusterServiceDNSRecordController) LookupAndReconcile(key string) (err error) {
	return c.reconcile(key)
}

func (c *MultiClusterServiceDNSRecordController) reconcile(key string) (err error) {
	var namespace, name string

	if c.BeforeReconcile != nil {
		c.BeforeReconcile(key)
	}
	if c.AfterReconcile != nil {
		// Wrap in a function so err is evaluated after it is set
		defer func() { c.AfterReconcile(key, err) }()
	}

	namespace, name, err = cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	u, err := c.controller.Get(namespace, name)
	if errors.IsNotFound(err) {
		glog.Infof("Not doing work for MultiClusterServiceDNSRecord %v because it has been deleted", key)
		// Set error so it is picked up by AfterReconcile and the return function
		err = nil
		return
	}
	if err != nil {
		glog.Errorf("Unable to retrieve MultiClusterServiceDNSRecord %v from store: %v", key, err)
		return
	}

	// Set error so it is picked up by AfterReconcile and the return function
	err = c.controller.Reconcile(u)

	return
}

func (c *MultiClusterServiceDNSRecordController) Run(stopCh <-chan struct{}) {
	for _, q := range c.Informers.WorkerQueues {
		q.Run(stopCh)
	}
	controller.GetDefaults(c.controller).Run(stopCh)
}
