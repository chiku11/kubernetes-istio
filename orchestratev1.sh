kubectl create namespace istio-demo
kubectl label namespace istio-demo istio-injection=enabled
istioctl install --set profile=demo -y

cd clientserver/v1
docker build -t clientserver:1.0 .  
kubectl apply -f clientserver.deploy.yml
kubectl apply -f clientserver.svc.yml

cd ..
cd weatherserver/v1
docker build -t weatherserver:1.0 .  
kubectl apply -f weatherserver.deploy.yml
kubectl apply -f weatherserver.svc.yml

cd ..
cd stockpriceserver/v1
docker build -t stockpriceserver:1.0 .  
kubectl apply -f stockpriceserver.deploy.yml
kubectl apply -f stockpriceserver.svc.yml

cd ../..
cd istio
kubectl apply -f istio-gateway.yml
kubectl apply -f istio-virtualservice.yml
