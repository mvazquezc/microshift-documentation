/*
Copyright © 2021 Microshift Contributors

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
package controllers

import (
	"github.com/sirupsen/logrus"

	kubescheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func KubeScheduler(ready chan bool) {
	command := kubescheduler.NewSchedulerCommand()
	args := []string{
		"--config=/etc/kubernetes/ushift-resources/kube-scheduler/config/config.yaml",
		//"--cert-dir=/var/run/kubernetes",
		//"--port=0",
		//"--authentication-kubeconfig=" + constant.KubeSchedulerKubeconfigPath,
		//"--authorization-kubeconfig=" + constant.KubeSchedulerKubeconfigPath,
		//"--feature-gates=APIPriorityAndFairness=true,LegacyNodeRoleBehavior=false,NodeDisruptionExclusion=true,RemoveSelfLink=false,RotateKubeletServerCertificate=true,SCTPSupport=true,ServiceNodeExclusion=true,SupportPodPidsLimit=true",
		"--v=2",
		"--tls-cert-file=/etc/kubernetes/ushift-resources/kube-scheduler/secrets/tls.crt",
		"--tls-private-key-file=/etc/kubernetes/ushift-resources/kube-scheduler/secrets/tls.key",
	}
	if err := command.ParseFlags(args); err != nil {
		logrus.Fatalf("failed to parse flags:%v", err)
	}
	go func() {
		command.Run(command, nil)
		logrus.Fatalf("kube-scheduler exited")
	}()

	ready <- true
}
