// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package graphql

var mockClusters = []*cluster{
	&cluster{name: "cluster1"},
	&cluster{name: "cluster2"},
}

func mockQueryCluster(name *string) []*cluster {
	var cs []*cluster

	for _, c := range mockClusters {
		if name == nil || *name == "" || *name == c.name {
			cs = append(cs, c)
		}
	}

	return cs
}
