docker stop block7serv
docker rm block7serv
docker run -d --name block7serv -p 3723:3723 block7serv