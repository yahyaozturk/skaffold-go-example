kubectl create secret generic mongodb --from-literal=mongoHost="testdb-mongo-mongodb.default.svc.cluster.local" --from-literal=mongoUser="admin" --from-literal=mongoPassword="admin"