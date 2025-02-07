/*
 * Copyright 2023 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha3

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	servicebindingv1beta1 "github.com/servicebinding/runtime/apis/v1beta1"
)

var _ conversion.Convertible = (*ServiceBinding)(nil)

func (src *ServiceBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*servicebindingv1beta1.ServiceBinding)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = src.Spec
	dst.Status = src.Status

	return nil
}

func (dst *ServiceBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*servicebindingv1beta1.ServiceBinding)

	dst.ObjectMeta = src.ObjectMeta
	dst.Spec = src.Spec
	dst.Status = src.Status

	return nil
}
