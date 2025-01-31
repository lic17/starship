#!/bin/bash -e

if [[ $# -gt 1 ]]; then
  echo "Extra arguments, should be '$0 [tag]', exit ..."
  exit 1
fi

tag=""

if [[ $# == 1 ]]; then
  tag="$1"
fi

if [[ ${tag} == "" ]]; then
  commit=$(git rev-parse HEAD)
  branch=$(git rev-parse --abbrev-ref HEAD)
  tag="${commit}_${branch}"
  if [ -n "$(git status --porcelain)" ]; then
    tag="${tag}_dirty"
    echo "Current branch is dirty"
    echo "Use tag ${tag} ..."
  fi
fi

echo "Using tag: ${tag}"

echo "Logging into docker ..."
aws ecr-public get-login-password --region us-east-1 |\
    docker login --username AWS --password-stdin public.ecr.aws/tricorder

bazel run -c opt --define=TAG=${tag} --define=REGISTRY=public.ecr.aws/tricorder //src/api-server/cmd:push_api-server_image
bazel run -c opt --define=TAG=${tag} --define=REGISTRY=public.ecr.aws/tricorder //src/agent/cmd:push_agent_image

ToT=$(git rev-parse --show-toplevel)
${ToT}/devops/release/build_and_push_mgmt_ui_image.sh ${tag}
