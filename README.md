# world-talent

Primera versión de la plataforma de adquisición de talentos con Inteligencia Artificial - AI World Talent con APIs REST y Arquitectura de Microservicios desplegadas en Kubernetes Cluster con infraestructuras de Nube:

1. OCI OKE (Oracle Kubernetes Engine) | Oracle Cloud Infrastructure
2. AWS EKS (Elastic Kubernetes Service) | Amazon Web Services

## Instrucciones para construir la imagen Docker
1. sudo docker build . -t oneitworld/world-talent-app
2. sudo docker images

## Instrucciones para correr imagenes Docker
1. sudo docker volume create mysql-data
2. sudo docker network create mi-red
3. sudo docker run -d -p 3307:3306 --name mysql-server --network mi-red -e MYSQL_ROOT_PASSWORD=Kalifornia2024$ -e MY_DATABASE=sys -v mysql-data:/var/lib/mysql mysql:latest
4. sudo docker run -d -p 9091:9090 --network mi-red oneitworld/world-talent-app
5. sudo docker ps -a
6. sudo docker volume ls
7. sudo docker inspect mysql-server