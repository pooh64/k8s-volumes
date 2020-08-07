package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	//"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	//apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//typev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	//"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	//"k8s.io/client-go/util/homedir"
	//"k8s.io/client-go/util/retry"
)

func main() {
	client, err := getClientDynamic(getConfigPath())
	if err != nil {
		panic(err.Error())
	}

	snapshotResource := schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshots"}

	snapshots, err := client.Resource(snapshotResource).Namespace("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("There are %d VolumeSnapshots in the cluster\n", len(snapshots.Items))
	for indx, entry := range snapshots.Items {
		//fmt.Println(indx)
		s, _ := json.MarshalIndent(entry, "", "\t")
		fmt.Printf("Snapshot[%d]:\n%s\n", indx, string(s))
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getConfigPath() string {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	return *kubeconfig
}

func getClientDynamic(configLocation string) (dynamic.Interface, error) {
	kubeconfig := filepath.Clean(configLocation)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
