apiVersion: v1
kind: Pod
metadata:
  name: session
  namespace: develop
spec:
  containers:
  - name: teachstore-session
    image: teachstore-session
    volumeMounts:
    - name: config-volume
      #mountPath: /etc/config
      mountPath: /go/src/teachstore-session/config
    ports:
      - containerPort: 9480
    env:
    - name: version
      value: "__VERSION__"
    - name: IP_DOCKER_HOST
      value: "__IP_DOCKER_HOST__"  
    - name: MY_NODE_NAME
      valueFrom:
        fieldRef:
          fieldPath: spec.nodeName   
    - name: MY_POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: MY_POD_NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
    - name: MY_POD_IP
      valueFrom:
        fieldRef:
          fieldPath: status.podIP                 
  volumes:
    - name: config-volume
      configMap:
        name: teachstore-session-__VERSION__