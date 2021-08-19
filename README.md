# playing_minio

POC for using minio which installed as standalone server

Installing Minio using docker:
docker run -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
credentials: minioadmin:minioadmin
