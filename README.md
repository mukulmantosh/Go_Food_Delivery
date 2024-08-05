# Go Food Delivery

![Coverage](https://img.shields.io/badge/Coverage-55.5%25-yellow)
![Workflow](https://github.com/mukulmantosh/Go_Food_Delivery/actions/workflows/test.yaml/badge.svg)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/04452c54468446dfbcc566604e69f379)](https://app.codacy.com?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

![background](./misc/images/background.png)

This project has just begun.


Run Postgres db

```bash
docker run --name food-delivery -p 5432:5432 -e POSTGRES_PASSWORD=****** -d postgres
```


Update env

```
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=mukul123
DB_NAME=food_delivery
DB_PORT=5432
STORAGE_TYPE=local
STORAGE_DIRECTORY=uploads
LOCAL_STORAGE_PATH=C:\Users\win10\GolandProjects\Go_Food_Delivery\uploads
```