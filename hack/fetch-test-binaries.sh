#!/usr/bin/env bash
set -euo pipefail
readonly tmpdir="${TMPDIR:-/tmp}"
readonly tarball="${tmpdir}/kubebuilder-tools.tar.gz"
curl -o "${tarball}" https://storage.googleapis.com/kubebuilder-tools/kubebuilder-tools-1.14.1-darwin-amd64.tar.gz
tar -C "${tmpdir}" -zvxf "${tarball}"
