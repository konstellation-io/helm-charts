#!/bin/sh

BUILD_DOCKER_IMAGES=0
cmd_deploy() {
  case $* in
    --build)
      BUILD_DOCKER_IMAGES=1
    ;;
  esac

  deploy
}

show_deploy_help() {
  echo "$(help_global_header "deploy")

    options:
      --build  re-build all docker images before deploying on minikube.
      --clean  sends a prune command to remove old docker images and containers.

    $(help_global_options)
"
}

deploy() {
  minikube_start
  prepare_helm

  if [ "$BUILD_DOCKER_IMAGES" = "1" ]; then
    build_docker_images
  fi

  replace_env_vars
  create_namespace

  if [ "$ENABLE_TLS" != "false" ]; then
    ./scripts/create_self_signed_cert.sh $NAMESPACE $DOMAIN
  fi

  deploy_helm_chart
}

HELM_READY=""

prepare_helm() {
  if [ "$HELM_READY" = "1" ]; then
    return
  fi

  if [ "$HELM_VERSION" = "2" ]; then
    # Helm v2 needs to be initiated first
    echo_wait "Init helm v2..."
    run helm init --upgrade --wait
  fi

  if [ "$MINIKUBE_RESET" = "1" ]; then
    clean_helm_deps
  fi

  HELM_READY=1
}

clean_helm_deps() {
  rm -rf helm/kdl-server/charts/*
  helm dep update helm/kdl-server
}

get_kubectl_dry_run() {
    act_version="$(kubectl version --client=true --short | sed 's/[^0-9.]*\([0-9.]*\).*/\1/')"
    req_version="1.18.0"

    # get the lowest version of the two compared
    lowest_version=$(printf '%s\n' "${act_version}" "${req_version}" | sort -V | head -n1)

    # if minimum required is met, use newer parameter
    if [ "$lowest_version" = "$req_version" ]; then
      echo "--dry-run=client"
      return
    fi

    echo "--dry-run"
}

create_namespace() {
  DRY_RUN=$(get_kubectl_dry_run)
  echo_info "📚️ Create Namespace if not exist..."
  NS=$(kubectl create ns "${NAMESPACE}" ${DRY_RUN} -o yaml)
  if [ "$VERBOSE" = "1" ]; then
    # NOTE: there is no way to call run() with pipe commands
    echo_run "kubectl create ns \"${NAMESPACE}\" ${DRY_RUN} -o yaml | kubectl apply -f -"
    echo "$NS" | kubectl apply -f -
  else
    echo "$NS" | kubectl apply > /dev/null -f - 2>&1
  fi
}

deploy_helm_chart() {
  echo_info "📦 Applying helm chart..."
  run helm dep update helm/kdl-server
  run helm upgrade \
    --wait \
    --install "${RELEASE_NAME}" \
    --namespace "${NAMESPACE}" \
    --set domain=$DOMAIN \
    --set science-toolkit.domain=$DOMAIN \
    --set science-toolkit.tls.enabled=$ENABLE_TLS \
    --set science-toolkit.minio.securityContext.runAsUser=0 \
    --set science-toolkit.gitea.admin.username=$GITEA_ADMIN_USER \
    --set science-toolkit.gitea.admin.password=$GITEA_ADMIN_PASSWORD \
    --timeout 60m \
    helm/kdl-server
}
