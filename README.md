# JMIND

# Task:

✅ Implementation in Golang.

✅ Readable and structured code.

✅ Use of standard and widely adopted libs, tools, and Go practices as much as possible.

✅ Runtime-specific parameters (for example, external API key) can be passed to the service as a
  configuration without code changes. The configuration approach is up to you.
  
✅ Unit tests of internal application logic.

✅ If the service has dependencies (for example, database), provide a docker-compose file
so anyone can start the dependencies and test the implementation.

# Bonus points / nice to have:
✅ Testing of the implemented endpoint on HTTP level. (file JMIND/TestAPI.http)

✅ As accessing external data sources can be time consuming, consider some caching or persistent
storage for already requested blocks. 

🔲 Integration testing of the implemented 3rd-party API client.

✅ Graceful shutdown (if implemented right)



### To start app:

##### Need to:
1) Pull this repository

2) Download docker - https://www.docker.com/products/docker-desktop/

#### After:
1) In cmd enter "docker-compose up" to initialize mongodb (port 27017 must be unused or change in docker-compose.yml file)
2) API :

localhost:8000/api/block/<block_number>/total


### screenshots of working program
## API:
![img](https://user-images.githubusercontent.com/57154344/164715619-2fe2e01b-9966-4698-88dd-c4e9f2a84410.png)
## Mongo:
![img_1](https://user-images.githubusercontent.com/57154344/164715680-b6bda4c6-66af-4368-b859-56fd2b5b28a2.png)

