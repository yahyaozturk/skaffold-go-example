for i in {1..10000}
do
   curl -X PUT -H "Content-Type: application/json" -d '{ "dateOfBirth": "2000-01-01" }' http://localhost:8080/api/v1/hello/$i
   echo ""
   curl http://localhost:8080/api/v1/hello/$i
   echo ""
done



for i in {1..10000}
do
   curl http://localhost:8080/api/v1/hello/$i
   echo ""
done


for i in {7000..10000}
do
   curl http://10.0.1.240:8080/api/v1/hello/$i
   echo ""
done



for i in {7000..10000}
do
   curl -X PUT -H "Content-Type: application/json" -d '{ "dateOfBirth": "2000-01-01" }' http://10.0.1.5:8080/api/v1/hello/$i
done