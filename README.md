# Traefik simplecontainer provider
Traefik simplecontainer provider exposes http provider for the Traefik to consume. It can handle distributed access to containers via single traefik deployment.

## Quick start
Deploy traefik provider with next static configuration:

```yaml
prefix: simplecontainer.io/v1
kind: containers
meta:
  name: traefik-provider
  group: traefik
  labels:
    test: "testing"
spec:
  image: "quay.io/simplecontainer/traefik-provider"
  tag: "latest"
  replicas: 1
  ports:
    - container: "7431"
```

Afterwards, deploy traefik using traefik provider as main provider:

```yaml
providers:
  http:
    endpoint: "http://bridge.traefik.traefik-provider.private"    
    pollInterval: "1s"
    
entryPoints:
  web:
    address: ":80"
  websecure:
    address: ":443"
      
api:
  insecure: true
  dashboard: true
```

Now definining single or cross node ingress is simple using custom resource:

```yaml
prefix: traefik.io/v1
kind: custom
meta:
  name: dynamic
  group: traefik
spec:
  traefik:
    enable: "true"
    http:
      routers:
        my-router:
          rule: "Host(`example.local`)"
          entrypoints:
          - "web"
          service: "my-service"
      services:
        my-service:
          loadBalancer:
            servers:
              - url: http://bridge.nginx.nginx.private
```

Using this custom resource you can apply traefik dynamic configuration. It follows the same syntax as [file provider](https://doc.traefik.io/traefik/providers/file/).

This example will load balance requests to docker containers on the all nodes if domain is from cluster network eg. `cluster.nginx.nginx.private`. 
`cluster.group.name.private` is headless domain (return IP addresses of all replicas).

Traefik will use simplecontainer node DNS and forward requests correctly even if container is running on another node.