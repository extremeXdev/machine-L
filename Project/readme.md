=======

### initalisation du projet golang 

### information > pas besoin de faire car projet déjà crée

> go mod init github.com/extremeXdev/machine-L/Project/api


### pour tester  
> cd github.com/extremeXdev/machine-L/Project/api


> go run cmd/main.go

``
2023/07/05 08:10:55 No .env file found

2023/07/05 08:10:55 You must set your 'MONGODB_URI' environmental variable. 
See https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable
exit status 1
``

####  

https://blog.logrocket.com/integrating-mongodb-go-applications/


## 1 - Todo

#### a - connection mongodb

tester la connection a une base de données avec mongodb et golang

#### b - definir une structure de collection mongodb

apres definir une collection structure de collection pour une base de données

### c - regarder go-swagger

lien d'exemple fait 


https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-fiber-version-4la0

> lien de golang detaillée

https://goswagger.io/generate/client.html

> add dockerfile

sudo docker build -t eylndoweb:v1 .
sudo docker run -p 192.168.1.66:2502:2501 -it eylndoweb:v1