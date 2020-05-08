helm repo add stable https://kubernetes-charts.storage.googleapis.com/
helm repo update

helm install testdb-mongo stable/mongodb --set mongodbUsername=admin,mongodbPassword=password,mongodbDatabase=testdb