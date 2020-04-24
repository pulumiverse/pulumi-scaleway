#!/bin/bash

set -o errexit
set -o pipefail

if [ "$#" -ne 1 ]; then
    >&2 echo "usage: $0 <path-to-repository-root>"
    exit 1
fi

SCRIPT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
SOURCE_ROOT="$1"

if [ -n "${NUGET_TOKEN}" ]; then
    find "${SOURCE_ROOT}/sdk/dotnet/bin/Debug/" -name 'Pulumi.*.nupkg' \
        -exec dotnet nuget push -k ${NUGET_TOKEN} -s https://api.nuget.org/v3/index.json {} ';'
else
    echo "No nuget token!"
    exit 1
fi
