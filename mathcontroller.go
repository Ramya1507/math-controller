package main

import (
	"fmt"
	"time"
  "reflect"
  "context"


	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

  maths          "math-controller/pkg/apis/maths/v1alpha1"
	clientset      "math-controller/pkg/client/clientset/versioned"
	mathresourcescheme   "math-controller/pkg/client/clientset/versioned/scheme"
	informers      "math-controller/pkg/client/informers/externalversions/maths/v1alpha1"
	listers        "math-controller/pkg/client/listers/maths/v1alpha1"
)
const statemessage = "SUCCEED"
const statusmessage = "updated successfully"
const controllerAgentName = "math-controller"



type Controller struct {
	kubeclientset kubernetes.Interface

	mathclientset clientset.Interface

	mathresourcesLister listers.MathResourceLister
	mathresourcesSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	informer cache.SharedIndexInformer

	recorder record.EventRecorder
}
func NewController(
	kubeclientset kubernetes.Interface,mathclientset clientset.Interface,
	mathResourceInformer informers.MathResourceInformer) *Controller {

	utilruntime.Must(mathresourcescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:    kubeclientset,
		mathclientset:    mathclientset,
		mathresourcesLister:   mathResourceInformer.Lister(),
		mathresourcesSynced:   mathResourceInformer.Informer().HasSynced,
		workqueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "mathresource"),
		recorder:         recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when Student resources change
	mathResourceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueMathResource,
		UpdateFunc: func(old, new interface{}) {
        newMath := new.(*maths.MathResource)
        oldMath := old.(*maths.MathResource)
       // klog.Info("status:",newMath.Status.State)
        //klog.Info("Message:",newMath.Status.Message)
        if reflect.DeepEqual(newMath.Spec, oldMath.Spec) {
        
        klog.Info("Specs not modified. Ignoring update event")
        return
       }
    controller.enqueueMathResource(new)
    },
		DeleteFunc: controller.enqueueMathResourceForDelete,
	})

	return controller
}

func (c *Controller) processNextItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		if key, ok = obj.(string); !ok {

			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		c.workqueue.Forget(obj)
		//klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}


func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	cmath, err := c.mathresourcesLister.MathResources(namespace).Get(name)
	if err != nil {
		klog.Errorf("Fetching CRD  with key %s from store failed with %v", key, err)
		return err
	}
  
	if cmath.Spec.Operation != "" {

		switch cmath.Spec.Operation {

		case ("add"):
			{
				klog.Infof("Operation Addition  value= %d \n", cmath.Spec.FirstNum + cmath.Spec.SecondNum)

			}

		case ("sub"):
			{
				klog.Infof("Operation subtraction value= %d \n", cmath.Spec.FirstNum - cmath.Spec.SecondNum)

			}
		case ("mul"):
			{
				klog.Infof("Operation multiplication  value= %d \n", cmath.Spec.FirstNum * cmath.Spec.SecondNum)

			}

		case ("div"):
			{
        if(cmath.Spec.SecondNum > 0){
				klog.Infof("Operation division value= %d \n", cmath.Spec.FirstNum / cmath.Spec.SecondNum)
        }else {
        klog.Errorf(" second num should be greater than zero")
        }

			}

		}

	} else {

		klog.Errorf("Fetching object cmath.Spec.Operation with  key %s from store failed with %v", key, err)
		return err

	}
 
  err = c.updateMathStatus(cmath)
  if err != nil {
		klog.Fatal(err)
	}
  c.recorder.Event(cmath, corev1.EventTypeNormal, "objecthandled", "objec is handled by custom controller")
	return nil

}



func (c *Controller) updateMathStatus(cmath *maths.MathResource) error {

  mathCopy := cmath.DeepCopy()
  mathCopy.Status.State = statemessage
  mathCopy.Status.Message = statusmessage
  _, err := c.mathclientset.MathsV1alpha1().MathResources(cmath.Namespace).UpdateStatus(context.TODO(), mathCopy, metav1.UpdateOptions{})
  
  return err
}


func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	// Let the workers stop when we are done
	defer c.workqueue.ShutDown()
	klog.Info("start controller Business, start a cache data synchronization")
	if ok := cache.WaitForCacheSync(stopCh, c.mathresourcesSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Infof("worker start-up")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("worker Already started")
	<-stopCh
	klog.Info("worker It's already over.")

	return nil

}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) enqueueMathResource(obj interface{}) {
	var key string
	var err error
	// Cache objects
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}

	// Queue key s
	c.workqueue.AddRateLimited(key)
}

// Delete operation
func (c *Controller) enqueueMathResourceForDelete(obj interface{}) {
	var key string
	var err error
	// Delete the specified object from the cache
	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	//Queue the key again
	c.workqueue.AddRateLimited(key)
}
