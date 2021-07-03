# User

#### MySQL
```bash
docker run --name go-news --restart=always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

#### Database
```sql
CREATE DATABASE user DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE `ts_account` (
                              `user_id` varchar(256) COLLATE utf8mb4_bin NOT NULL,
                              `account` varchar(256) COLLATE utf8mb4_bin NOT NULL,
                              `password` varchar(256) COLLATE utf8mb4_bin NOT NULL,
                              `create_time` int NOT NULL,
                              `update_time` int NOT NULL,
                              PRIMARY KEY (`user_id`),
                              UNIQUE KEY `ts_account_account_uindex` (`account`),
                              UNIQUE KEY `ts_account_user_id_uindex` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `ts_user` (
                           `user_id` varchar(64) COLLATE utf8mb4_bin NOT NULL,
                           `user_name` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
                           `telephone` varchar(256) COLLATE utf8mb4_bin DEFAULT '',
                           `email` varchar(256) COLLATE utf8mb4_bin DEFAULT '',
                           `age` int NOT NULL DEFAULT '0',
                           `create_time` int NOT NULL,
                           `update_time` int NOT NULL,
                           PRIMARY KEY (`user_id`),
                           UNIQUE KEY `ts_user_user_id_uindex` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
```

#### User Server
```bash
docker build . -t go-news-user

docker run -d --name user -p 8888:8888 go-news-user
```
