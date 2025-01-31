// Copyright The OpenTelemetry Authors
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

package sdkapi // import "go.opentelemetry.io/otel/metric/sdkapi"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/number"
)

type noopInstrument struct{}
type noopBoundInstrument struct{}
type noopSyncInstrument struct{ noopInstrument }
type noopAsyncInstrument struct{ noopInstrument }

var _ SyncImpl = noopSyncInstrument{}
var _ BoundSyncImpl = noopBoundInstrument{}
var _ AsyncImpl = noopAsyncInstrument{}

// NewNoopSyncInstrument returns a No-op implementation of the
// synchronous instrument interface.
func NewNoopSyncInstrument() SyncImpl {
	return noopSyncInstrument{}
}

// NewNoopSyncInstrument returns a No-op implementation of the
// asynchronous instrument interface.
func NewNoopAsyncInstrument() SyncImpl {
	return noopSyncInstrument{}
}

func (noopInstrument) Implementation() interface{} {
	return nil
}

func (noopInstrument) Descriptor() Descriptor {
	return Descriptor{}
}

func (noopBoundInstrument) RecordOne(context.Context, number.Number) {
}

func (noopBoundInstrument) Unbind() {
}

func (noopSyncInstrument) Bind([]attribute.KeyValue) BoundSyncImpl {
	return noopBoundInstrument{}
}

func (noopSyncInstrument) RecordOne(context.Context, number.Number, []attribute.KeyValue) {
}
