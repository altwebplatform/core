package k8s

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/pkg/api/v1"
	"log"
	"strings"
)

//docker run --rm -p 9000:9000 -e MINIO_ACCESS_KEY -e MINIO_SECRET_KEY minio/minio server /export &

func main() {
	kubeconfig := flag.String("kubeconfig", "./.kubeconfig", "absolute path to the kubeconfig file")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	namespace, err := clientset.Namespaces().Get(NAMESPACE)
	if (err != nil && strings.Contains(err.Error(), "not found")) || namespace == nil {
		namespace, err = clientset.Namespaces().Create(&v1.Namespace{
			ObjectMeta: v1.ObjectMeta{Name: NAMESPACE},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
