/*
Copyright 2018 The Rook Authors. All rights reserved.

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

package v1

func (p *PoolSpec) IsReplicated() bool {
	return p.Replicated.Size > 0
}

func (p *PoolSpec) IsErasureCoded() bool {
	return p.ErasureCoded.CodingChunks > 0 || p.ErasureCoded.DataChunks > 0
}

func (p *PoolSpec) IsCompressionEnabled() bool {
	return p.CompressionMode != ""
}

func (p *ReplicatedSpec) IsTargetRatioEnabled() bool {
	return p.TargetSizeRatio != 0
}
