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

package schemes

import (
	"k8s.io/perf-tests/benchmark/pkg/metricsfetcher/util"
)

// GetLastNJobRuns returns a list of run numbers of last 'n' completed runs of
// 'job'. If it cannot find n runs, returns as many runs as it could collect.
func GetLastNJobRuns(job string, n int, utils util.JobLogUtils) ([]int, error) {
	latestRunNumber, err := utils.GetLatestBuildNumberForJob(job)
	if err != nil {
		return nil, err
	}

	var runs []int
	for runNumber := latestRunNumber; runNumber > 0 && len(runs) < n; runNumber-- {
		if _, err := utils.GetJobRunFinishedStatus(job, runNumber); err == nil {
			runs = append(runs, runNumber)
		}
	}
	return runs, nil
}
