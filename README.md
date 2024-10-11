# Go Food Delivery

![Coverage](https://img.shields.io/badge/Coverage-50.4%25-yellow)
![Workflow](https://github.com/mukulmantosh/Go_Food_Delivery/actions/workflows/test.yaml/badge.svg)

![background](./misc/images/background.png)


### Prerequisites

Before starting up this project, make sure you have the necessary dependencies installed in your machine.

###  Installation

- [x] [Go](https://go.dev/) - Go is an open source programming language that makes it simple to build secure, scalable systems.

- [x] [Docker](https://www.docker.com/) - Docker helps developers bring their ideas to life by conquering the complexity of app development.

- [x] [PostgreSQL](https://www.postgresql.org/) - The World's Most Advanced Open Source Relational Database

- [x] [NATS](https://nats.io/) - NATS is an open-source messaging system. The NATS server is written in the Go programming language



#### Running Postgres Database

```bash
docker run --name food-delivery -p 5432:5432 -e POSTGRES_PASSWORD=****** -d postgres
```

#### Running NATS

```bash
docker network create nats
docker run --name nats -d --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222 --cluster_name NATS --cluster nats://0.0.0.0:6222
```


### Environment Variables

Be sure to place the `.env` file in the project root and update the information according to your settings. Refer to the example below.

```
APP_ENV=dev
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=*************
DB_NAME=food_delivery
DB_PORT=5432
STORAGE_TYPE=local
STORAGE_DIRECTORY=uploads
LOCAL_STORAGE_PATH=C:\Users\win10\GolandProjects\Go_Food_Delivery\uploads
UNSPLASH_API_KEY=*******************
JWT_SECRET_KEY=********************
PASSWORD_SALT=********************
```

### External APIs

We are using [UnSplash](https://unsplash.com/) to generate images. So, you need to have an API key to work with the application. 


### Frontend (UI)

To configure the application's frontend UI, be sure to follow the instructions in this [repository][repo].


[repo]: https://github.com/mukulmantosh/food_delivery_frontend

### HTTP Client

We have covered the APIs which you can directly test out from the IDE. 

![httpclient](./misc/images/httpclient.png)

### Get Set Go 🚀

Once you’re ready, clone this repo in [GoLand,](https://www.jetbrains.com/go/) and you're good to go.

![go_run_config](./misc/images/go_run_config.png)
