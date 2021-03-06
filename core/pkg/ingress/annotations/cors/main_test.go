/*
Copyright 2017 The Kubernetes Authors.

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

package cors

import (
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

const (
	notCorsAnnotation = "ingress.kubernetes.io/enable-not-cors"
)

func TestParse(t *testing.T) {
	ap := NewParser()
	if ap == nil {
		t.Fatalf("expected a parser.IngressAnnotation but returned nil")
	}

	testCases := []struct {
		annotations map[string]string
		expected    bool
	}{
		{map[string]string{annotation: "true"}, true},
		{map[string]string{annotation: "false"}, false},
		{map[string]string{notCorsAnnotation: "true"}, false},
		{map[string]string{}, false},
		{nil, false},
	}

	ing := &extensions.Ingress{
		ObjectMeta: api.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: extensions.IngressSpec{},
	}

	for _, testCase := range testCases {
		ing.SetAnnotations(testCase.annotations)
		result, _ := ap.Parse(ing)
		if result != testCase.expected {
			t.Errorf("expected %t but returned %t, annotations: %s", testCase.expected, result, testCase.annotations)
		}
	}
}
