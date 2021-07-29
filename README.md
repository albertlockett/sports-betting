# sports-betting

## build
It's build using GCP cloud build


## TODO automate some of this
log into gcp
```
gcloud auth application-default login
```

setup k8s cluster
```
cd terraform/k8s
terraform apply
```
TODO - need to grant the service account permissions to pull from gcr


setup k8s to talk to cluster
```
gcloud container clusters get-credentials
gcloud container clusters get-credentials my-gke-cluster --region us-east1
```

install es
```
helm install elasticsearch elastic/elasticsearch -f helm/elasticsearch/values.yaml
```

create indices (port forward first)
```
curl -XPUT localhost:9200/handicaps
curl -XPUT localhost:9200/lines
curl -XPUT localhost:9200/expected-values
```