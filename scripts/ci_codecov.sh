#!/bin/bash -eu
set -x

cd $(dirname ${BASH_SOURCE})/..

bash <(curl -s https://codecov.io/bash)
