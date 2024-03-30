VERSION=0.1.2
docker build -t lister:$VERSION .

docker tag lister:$VERSION theluckiestsoul/lister:$VERSION

docker push theluckiestsoul/lister:$VERSION

kubectl apply -f deployment.yaml