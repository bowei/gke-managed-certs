/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"sort"

	compute "google.golang.org/api/compute/v0.alpha"

	api "managed-certs-gke/pkg/apis/alpha.cloud.google.com/v1alpha1"
)

func Equals(mcert *api.ManagedCertificate, sslCert *compute.SslCertificate) bool {
	if len(mcert.Spec.Domains) != len(sslCert.Managed.Domains) {
		return false
	}

	mDomains := make([]string, len(mcert.Spec.Domains))
	copy(mDomains, mcert.Spec.Domains)
	sort.Strings(mDomains)

	sDomains := make([]string, len(sslCert.Managed.Domains))
	copy(sDomains, sslCert.Managed.Domains)
	sort.Strings(sDomains)

	for i, v := range mDomains {
		if v != sDomains[i] {
			return false
		}
	}

	return true
}
