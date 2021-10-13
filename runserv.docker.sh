docker stop block7serv
docker rm block7serv
docker run -d --name block7serv -p 3723:3723 -v $PWD/logs:/app/block7serv/logs -v $PWD/data:/app/block7serv/data -v $PWD/cfg:/app/block7serv/cfg block7serv