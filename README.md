# Webtoon REST API
![Cover](https://raw.githubusercontent.com/LuigiEnzoFerrari/Webtoon-API/master/.github/images/webtoons_background.jpg)  
## About  

This is currently a REST API to get the most useful information about any webtoon avaliable.

## Requirements  

### Install golang  

https://go.dev/doc/install  

### Create DATABASE (Postgres)

* Install postgres

* Create a database (optional name)

* Populate the database with the dataset: https://www.kaggle.com/datasets/victorsoeiro/webtoons-dataset?resource=download  

## How to use  

Clone the repository  
Inside the repository use the following commands to install the modules
```sh
go mod init <name-module>
go mod tidy
```
Set the environment variables as follow
```sh
DB_USER=<your-database-username>
DB_NAME=<your-database-name>
DB_PASS=<your-database-passoword>
```

Now  execute the main file  

```sh
go run main.go
```  
The application will be running on localhost:8080  

To list all the Webtoons available go on localhost:8080/webtoons 

To search an specific webtoon go on localhost:8080/search?title=webtoon-title  
