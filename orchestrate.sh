cd clientserver
docker build -t clientserver:1.0 .  
kubectl apply -f clientserver.deploy.yml
kubectl apply -f clientserver.svc.yml
cd ..
cd weatherserver
docker build -t weatherserver:1.0 .  
kubectl apply -f weatherserver.deploy.yml
kubectl apply -f weatherserver.svc.yml
cd ..
cd stockpriceserver
docker build -t stockpriceserver:1.0 .  
kubectl apply -f stockpriceserver.deploy.yml
kubectl apply -f stockpriceserver.svc.yml