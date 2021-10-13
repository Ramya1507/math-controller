package main

import (
	"flag"
	"path/filepath"
	"time"

	clientset "math-controller/pkg/client/clientset/versioned"

	informers "math-controller/pkg/client/informers/externalversions"
	"math-controller/pkg/signals"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

var kubeconfig *string

func main() {

	stopCh := signals.SetupSignalHandler()

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// creates the connection
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		klog.Fatal(err)
	}
	

	// creates the clientset
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	mathClient, err := clientset.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}


	mathInformerFactory := informers.NewSharedInformerFactory(mathClient, time.Second*30)

	controller := NewController(kubeClient,mathClient, mathInformerFactory.Maths().V1alpha1().MathResources())



	// Now let's start the controller



	mathInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}


}
