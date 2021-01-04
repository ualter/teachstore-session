SOURCE_YML="./config/config.yaml"
CONFIGMAP_TPL="./k8s/k8s/_configmap.tpl"
POD_TPL="./k8s/k8s/_pod.tpl"


configmap=""
while IFS= read -r line; do
	configmap="${configmap}\n$line"
done < "$CONFIGMAP_TPL"

idx=0
while IFS= read -r line; do
	lin=$(echo "${line}" | sed "s/ /#/g")
	if [ $idx = 0 ]; then
	    configmap="${configmap}$lin" 
	else
		configmap="${configmap}\n####$lin" 
	fi     	
	idx=$((idx+1))
done < "$SOURCE_YML"

printf "${configmap}" | sed "s/#/ /g" | sed "s/__VERSION__/$VERSION/g" > ./k8s/k8s/configmap.yaml
printf " \033[1;34m-->\033[0m Done!\n"

printf " \033[1;34m-->\033[0m Creating pod.yaml descriptor...\n"
printf " \033[1;34m---->\033[0m Placeholder version to \033[1;33m$VERSION\033[1;0m...\n"
cat k8s/k8s/_pod.tpl | sed "s/__VERSION__/$VERSION/g" > ./k8s/k8s/pod.yaml
printf " \033[1;34m---->\033[0m Placeholder IP_DOCKER_HOST to \033[1;33m$IP_DOCKER_HOST\033[1;0m...\n"
cat k8s/k8s/_pod.tpl | sed "s/__IP_DOCKER_HOST__/$IP_DOCKER_HOST/g" > ./k8s/k8s/pod.yaml
printf " \033[1;34m-->\033[0m Done!\n"
printf "\n"