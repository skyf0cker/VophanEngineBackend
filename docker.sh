if [ `whoami` = "root" ];then
  /etc/init.d/mysql stop
  docker stop mysql
  docker stop es
  docker rm mysql
  docker rm es
  docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
  docker run --name es -p 9200:9200 -p 9300:9300 -v $(pwd)/data:/usr/share/elasticsearch/data peterzhang/elasticsearch-analysis-ik
else
  echo "你不是root用户，请登录后执行脚本"
fi