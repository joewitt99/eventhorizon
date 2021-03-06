// Copyright (c) 2014 - Max Ekman <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

import "context"

// DefaultNamespace is the namespace to use if not set in the context.
const DefaultNamespace = "default"

type contextKey int

const (
	// namespaceKey is the context key for the namespace value.
	namespaceKey contextKey = iota
)

// Namespace returns the namespace from the context, or the default namespace.
func Namespace(ctx context.Context) string {
	if ns, ok := ctx.Value(namespaceKey).(string); ok {
		return ns
	}
	return DefaultNamespace
}

// WithNamespace sets the namespace to use in the context. The namespace is
// used to determine which database
func WithNamespace(ctx context.Context, namespace string) context.Context {
	return context.WithValue(ctx, namespaceKey, namespace)
}
