#!/bin/bash

# Copyright 2020 The KubeEdge Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(unset CDPATH && cd $(dirname "${BASH_SOURCE[0]}")/../.. && pwd)

${SCRIPT_ROOT}/vendor/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
github.com/kubeedge/kubeedge/cloud/pkg/client github.com/kubeedge/kubeedge/cloud/pkg/apis \
"devices:v1alpha1 reliablesync:v1alpha1" \
--go-header-file ${SCRIPT_ROOT}/cloud/hack/boilerplate/boilerplate.txt
