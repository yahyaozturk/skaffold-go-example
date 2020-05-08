kubectl create namespace ingress
helm repo update
helm install ingress stable/nginx-ingress --namespace ingress