# playing_minio

POC for using minio which installed as standalone server from docker

Installing Minio using docker:

docker run -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"

credentials: minioadmin:minioadmin

POC Details:

Get the buckets list

List the objects in each bucket

Read the available object from minio

convert the content to struct

write part of above content as another object