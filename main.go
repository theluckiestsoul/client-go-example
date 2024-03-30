package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	path.Join()
	defaultKubeConfigPath := filepath.Join(homeDir, ".kube", "config")
	cfgFile := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig file")
	cfg, err := clientcmd.BuildConfigFromFlags("", *cfgFile)
	if err != nil {
		fmt.Println("Using in-cluster config")
		cfg, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("failed to build in-cluster config: %v", err)
		}
	}

	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create clientset: %v", err)
	}

	pl, err := cs.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to list pods: %v", err)
	}

	fmt.Println("Pods:")
	for _, p := range pl.Items {
		fmt.Println(p.Name, p.Status.Phase)
	}

	ds, err := cs.AppsV1().Deployments("default").List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to list deployments: %v", err)
	}

	fmt.Println("Deployments:")
	for _, d := range ds.Items {
		fmt.Println(d.Name, d.Status.AvailableReplicas)
	}
}
