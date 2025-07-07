// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apikey "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/apikey"
	cluster "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/cluster"
	clusterconfig "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/clusterconfig"
	environment "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/environment"
	kafkaacl "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/kafkaacl"
	kafkatopic "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/kafkatopic"
	rolebinding "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/rolebinding"
	schema "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/schema"
	serviceaccount "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/confluent/serviceaccount"
	providerconfig "github.com/rouzbehsedighi/provider-confluent-sap/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		cluster.Setup,
		clusterconfig.Setup,
		environment.Setup,
		kafkaacl.Setup,
		kafkatopic.Setup,
		rolebinding.Setup,
		schema.Setup,
		serviceaccount.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
