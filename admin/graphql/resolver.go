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

type QueryResolver struct{}

func (r *QueryResolver) Clusters(args struct{ Name *string }) []*clusterResolver {
	var crs []*clusterResolver

	for _, c := range mockClusters {
		if args.Name == nil || *args.Name == "" || *args.Name == c.name {
			crs = append(crs, &clusterResolver{c: c})
		}
	}

	return crs
}

type clusterResolver struct {
	c *cluster
}

func (r *clusterResolver) Name() string {
	return r.c.name
}

/*
func (r *clusterResolver) Stats() *clusterStatsResolver {
	return nil
}

func (r *clusterResolver) Nodes() []*clusterNodeResolver {
	return nil
}

type clusterStatsResolver struct {
}

type clusterNodeResolver struct {
}
*/
