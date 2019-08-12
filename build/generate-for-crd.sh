#!/bin/bash -e

BUILD_DIR=$( dirname "${BASH_SOURCE[0]}" )
GENERATOR_REPO=https://github.com/kubernetes/code-generator.git
GENERATOR_VERSION=$(go list -m all | grep k8s.io/code-generator | rev | cut -d"-" -f1 | cut -d" " -f1 | rev)
GENERATOR_DIR="${BUILD_DIR}/../_local/code-generator/${GENERATOR_VERSION}"

mkdir -p "${GENERATOR_DIR}"
git clone "${GENERATOR_REPO}" "${GENERATOR_DIR}" 2> /dev/null || git -C "${GENERATOR_DIR}" pull
(cd ${GENERATOR_DIR} && git reset --hard ${GENERATOR_VERSION})

${GENERATOR_DIR}/generate-groups.sh \
   all \
   github.com/kheer/kheer/pkg/crd/api/generated \
   github.com/kheer/kheer/pkg/crd/api \
   kheer:v1alpha1