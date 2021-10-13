#

```
% kd cassdc dc1
Name:         dc1
Namespace:    default
Labels:       app.kubernetes.io/instance=k8ssandra-release
              app.kubernetes.io/managed-by=Helm
              app.kubernetes.io/name=k8ssandra
              app.kubernetes.io/part-of=k8ssandra-k8ssandra-release-default
              helm.sh/chart=k8ssandra-1.1.0
Annotations:  meta.helm.sh/release-name: k8ssandra-release
              meta.helm.sh/release-namespace: default
              reaper.cassandra-reaper.io/instance: k8ssandra-release-reaper
API Version:  cassandra.datastax.com/v1beta1
Kind:         CassandraDatacenter
```

```
% kd statefulsets k8ssandra-cluster-dc1-default-sts
Name:               k8ssandra-cluster-dc1-default-sts
Namespace:          default
CreationTimestamp:  Mon, 17 May 2021 18:45:27 -0400
Selector:           cassandra.datastax.com/cluster=k8ssandra-cluster,cassandra.datastax.com/datacenter=dc1,cassandra.datastax.com/rack=default
Labels:             app.kubernetes.io/managed-by=cass-operator
                    cassandra.datastax.com/cluster=k8ssandra-cluster
                    cassandra.datastax.com/datacenter=dc1
                    cassandra.datastax.com/rack=default
Annotations:        cassandra.datastax.com/resource-hash: O4SjTWWvzzd1sQxtySeNdaYsJxK/arWejIQJe4rQN8c=
```

```
% kd pod k8ssandra-cluster-dc1-default-sts-0
Name:         k8ssandra-cluster-dc1-default-sts-0
Namespace:    default
Priority:     0
Node:         gke-jdinoto-k8ssandra-api--small-pool-ebb55fdc-wm82/10.142.0.25
Start Time:   Mon, 17 May 2021 18:46:06 -0400
Labels:       app.kubernetes.io/managed-by=cass-operator
              cassandra.datastax.com/cluster=k8ssandra-cluster
              cassandra.datastax.com/datacenter=dc1
              cassandra.datastax.com/node-state=Started
              cassandra.datastax.com/rack=default
              cassandra.datastax.com/seed-node=true
              controller-revision-hash=k8ssandra-cluster-dc1-default-sts-749c949f5d
              statefulset.kubernetes.io/pod-name=k8ssandra-cluster-dc1-default-sts-0
Annotations:  <none>
Status:       Running
IP:           10.92.2.29
IPs:
  IP:           10.92.2.29
Controlled By:  StatefulSet/k8ssandra-cluster-dc1-default-sts
```

# Resources

## CassandraDatacenters

### cassdc dc1

```
Labels:       app.kubernetes.io/instance=k8ssandra
              app.kubernetes.io/managed-by=Helm
              app.kubernetes.io/name=k8ssandra
              app.kubernetes.io/part-of=k8ssandra-k8ssandra-default
              helm.sh/chart=k8ssandra-1.1.0
Annotations:  meta.helm.sh/release-name: k8ssandra
              meta.helm.sh/release-namespace: default
              reaper.cassandra-reaper.io/instance: k8ssandra-reaper
```

## StatefulSets

## statefulset k8ssandra-cassandra-cluster-dc1-default-sts

```
Labels:             app.kubernetes.io/managed-by=cass-operator
                    cassandra.datastax.com/cluster=k8ssandra-cassandra-cluster
                    cassandra.datastax.com/datacenter=dc1
                    cassandra.datastax.com/rack=default
Annotations:        cassandra.datastax.com/resource-hash: ks8SzAMrThj/HCYYZnAKxlSGY3gZn+96Rp4cCuyRuIg=
```

## Deployments

### deployment k8ssandra-cass-operator

```
Labels:                 app.kubernetes.io/instance=k8ssandra
                        app.kubernetes.io/managed-by=Helm
                        app.kubernetes.io/name=cass-operator
                        app.kubernetes.io/part-of=k8ssandra-k8ssandra-default
                        helm.sh/chart=cass-operator-0.29.1
Annotations:            deployment.kubernetes.io/revision: 1
                        meta.helm.sh/release-name: k8ssandra
                        meta.helm.sh/release-namespace: default
```

### deployment k8ssandra-dc1-stargate

```
Labels:                 app=k8ssandra-dc1-stargate
                        app.kubernetes.io/instance=k8ssandra
                        app.kubernetes.io/managed-by=Helm
                        app.kubernetes.io/name=k8ssandra
                        app.kubernetes.io/part-of=k8ssandra-k8ssandra-default
                        helm.sh/chart=k8ssandra-1.1.0
                        release=k8ssandra
Annotations:            deployment.kubernetes.io/revision: 1
                        meta.helm.sh/release-name: k8ssandra
                        meta.helm.sh/release-namespace: default
```

### deployment k8ssandra-grafana

```
Labels:                 app.kubernetes.io/instance=k8ssandra
                        app.kubernetes.io/managed-by=Helm
                        app.kubernetes.io/name=grafana
                        app.kubernetes.io/version=7.3.5
                        helm.sh/chart=grafana-6.1.17
Annotations:            deployment.kubernetes.io/revision: 1
                        meta.helm.sh/release-name: k8ssandra
                        meta.helm.sh/release-namespace: default
```

## Pods

### pod k8ssandra-cass-operator-7c4dc64969-bcrzp

```
Labels:       app.kubernetes.io/instance=k8ssandra
              app.kubernetes.io/managed-by=Helm
              app.kubernetes.io/name=cass-operator
              app.kubernetes.io/part-of=k8ssandra-k8ssandra-default
              helm.sh/chart=cass-operator-0.29.1
              pod-template-hash=7c4dc64969
```

### pod k8ssandra-cassandra-cluster-dc1-default-sts-0

```
Labels:       app.kubernetes.io/managed-by=cass-operator
              cassandra.datastax.com/cluster=k8ssandra-cassandra-cluster
              cassandra.datastax.com/datacenter=dc1
              cassandra.datastax.com/node-state=Started
              cassandra.datastax.com/rack=default
              cassandra.datastax.com/seed-node=true
              controller-revision-hash=k8ssandra-cassandra-cluster-dc1-default-sts-7f576657d9
              statefulset.kubernetes.io/pod-name=k8ssandra-cassandra-cluster-dc1-default-sts-0
```

```
Controlled By:  StatefulSet/k8ssandra-cassandra-cluster-dc1-default-sts
```

### pod k8ssandra-dc1-stargate-56b5d598f7-b2w5k

```
Labels:       app=k8ssandra-dc1-stargate
              pod-template-hash=56b5d598f7
```

```
Controlled By:  ReplicaSet/k8ssandra-dc1-stargate-56b5d598f7
```

### pod k8ssandra-reaper-d6bcd96c-zk927

```
Labels:       app.kubernetes.io/managed-by=reaper-operator
              pod-template-hash=d6bcd96c
              reaper.cassandra-reaper.io/reaper=k8ssandra-reaper
Annotations:  <none>
```

```
Controlled By:  ReplicaSet/k8ssandra-reaper-d6bcd96c
```

### pod k8ssandra-reaper-operator-f6bc9b77b-bfvkb

```
Labels:       app.kubernetes.io/instance=k8ssandra
              app.kubernetes.io/managed-by=Helm
              app.kubernetes.io/name=reaper-operator
              app.kubernetes.io/part-of=k8ssandra-k8ssandra-default
              helm.sh/chart=reaper-operator-0.30.0
              pod-template-hash=f6bc9b77b
Annotations:  <none>
```

### pod k8ssandra-kube-prometheus-operator-5556885bd6-hjjsx

```
Labels:       app=kube-prometheus-stack-operator
              chart=kube-prometheus-stack-12.11.3
              heritage=Helm
              pod-template-hash=5556885bd6
              release=k8ssandra
Annotations:  <none>
```

```
Controlled By:  ReplicaSet/k8ssandra-kube-prometheus-operator-5556885bd6
```

### pod prometheus-k8ssandra-kube-prometheus-prometheus-0

```
Labels:       app=prometheus
              controller-revision-hash=prometheus-k8ssandra-kube-prometheus-prometheus-6f98b87d4f
              operator.prometheus.io/name=k8ssandra-kube-prometheus-prometheus
              operator.prometheus.io/shard=0
              prometheus=k8ssandra-kube-prometheus-prometheus
              statefulset.kubernetes.io/pod-name=prometheus-k8ssandra-kube-prometheus-prometheus-0
Annotations:  <none>
```

```
Controlled By:  StatefulSet/prometheus-k8ssandra-kube-prometheus-prometheus
```

### pod k8ssandra-grafana-6858f6bbc-k6n4t

```
Labels:       app.kubernetes.io/instance=k8ssandra
              app.kubernetes.io/name=grafana
              pod-template-hash=6858f6bbc
Annotations:  checksum/config: 40bd7419006ff78ce88f60cea6235964381e95b4da50feec47ca97757efbba08
              checksum/dashboards-json-config: 01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b
              checksum/sc-dashboard-provider-config: 5d811572c6391b05043924ee8b72c217c7d4f56cb1e026d452c8ca303ea23549
              checksum/secret: 3d5b6452f6cebb714d32d1e0d503150632eeebfae12032ae3fd6eae11b1e7599
```

```
Controlled By:  ReplicaSet/k8ssandra-grafana-6858f6bbc
```
