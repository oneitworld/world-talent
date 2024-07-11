# world-talent

Primera versión de la plataforma de adquisición de talentos con Inteligencia Artificial - AI World Talent con APIs REST y Arquitectura de Microservicios desplegadas en Kubernetes Cluster con infraestructuras de Nube:

1. OCI OKE (Oracle Kubernetes Engine) | Oracle Cloud Infrastructure
2. AWS EKS (Elastic Kubernetes Service) | Amazon Web Services

## Instrucciones para construir la imagen Docker

sudo docker build . -t oneitworld/world-talent-app
sudo docker images

## Instrucciones para correr imagenes Docker

sudo docker network create mi-red
sudo docker run -d -p 3307:3306 --name mysql-server --network mi-red -e MYSQL_ROOT_PASSWORD=Kalifornia2024$ -e MY_DATABASE=world mysql:latest
sudo docker -d -p 9090:9090 run --network mi-red oneitworld/world-talent-app
sudo docker ps -a