# tryinghttps


### Production
```
docker-compose -f production.yaml up

```

### Development
```
docker-compose -f development.yaml up

```



### Build the API Docker image
```
docker image build -t hello .
docker container run -p 8090:8090 hello
```



#### Previous...
```bigquery
# Change permissions of acme.json in /configs dir
chmod 600 acme.json
```