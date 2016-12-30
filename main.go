package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	flag "github.com/spf13/pflag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig    string
	namespace     string
	labels        string
	version       bool
	podColumns    = []string{"NAME", "IMAGE", "STATUS", "RESTARTS", "START", "NAMESPACE"}
	deployColumns = []string{"NAME", "IMAGE", "NAMESPACE"}
)

func main() {

	flags := flag.NewFlagSet("kubeps", flag.ExitOnError)

	flags.Usage = func() {
		flags.PrintDefaults()
	}

	flags.StringVar(&kubeconfig, "kubeconfig", "", "Path of kubeconfig")
	flags.StringVar(&labels, "labels", "", "Label filter query")
	flags.StringVar(&namespace, "namespace", "", "Kubernetes namespace")
	flags.BoolVarP(&version, "version", "v", false, "Print version")

	// uses the current context in kubeconfig
	if kubeconfig == "" {
		if os.Getenv("KUBECONFIG") != "" {
			kubeconfig = os.Getenv("KUBECONFIG")
		} else {
			kubeconfig = clientcmd.RecommendedHomeFile
		}
	}

	if err := flags.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if version {
		printVersion()
		os.Exit(0)
	}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{})

	config, err := clientConfig.ClientConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	deployments, err := clientset.Deployments(namespace).List(v1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	deploymentPrint := new(tabwriter.Writer)
	deploymentPrint.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(deploymentPrint, strings.Join(deployColumns, "\t"))

	for _, deployment := range deployments.Items {
		for _, containers := range deployment.Spec.Template.Spec.Containers {
			fmt.Fprintln(deploymentPrint, strings.Join(
				[]string{deployment.Name, containers.Image, deployment.Namespace}, "\t",
			))
		}
	}
	fmt.Println("=== Deployment ===")
	deploymentPrint.Flush()
	fmt.Println()

	podList, err := clientset.Core().Pods(namespace).List(v1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	podPrint := new(tabwriter.Writer)
	podPrint.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(podPrint, strings.Join(podColumns, "\t"))

	for _, pod := range podList.Items {
		for _, container := range pod.Spec.Containers {
			if pod.Status.ContainerStatuses != nil {
				fmt.Fprintln(podPrint, strings.Join(
					[]string{
						pod.Name,
						container.Image,
						string(pod.Status.Phase),
						strconv.FormatInt(int64(pod.Status.ContainerStatuses[0].RestartCount), 10),
						pod.Status.StartTime.String(),
						pod.Namespace,
					}, "\t",
				))
			} else {
				fmt.Fprintln(podPrint, strings.Join(
					[]string{
						pod.Name,
						container.Image,
						string(pod.Status.Phase),
						"<none>",
						"<none>",
						pod.Namespace,
					}, "\t",
				))
			}
		}
	}
	fmt.Println("=== Pod ===")
	podPrint.Flush()
	fmt.Println()

}
