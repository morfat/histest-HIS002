# HIS002 

### The db design is exported SQL due to time contraints 



### Installlation


1. Create a .env file in the root folder and add the below encrinment variables replacing the data respectively
i.e <db-user> with your postgresl user, <db-password> , and <db-name>

```
DEBUG=True
DATABASE_URL = 'postgresql://<db-user>:<db-password>@localhost:5432/<db-name>'
SERVER_ADDRESS='localhost:8000'

```

2. Install necessary libraries as below
   
   ``` 
   go get -u github.com/gin-gonic/gin
   go get -u github.com/jackc/pgx/v5
   go get -u github.com/jackc/pgx/v5/pgxpool
   go get -u github.com/joho/godotenv
   ```



3. To run the code , do:
   `go run .`