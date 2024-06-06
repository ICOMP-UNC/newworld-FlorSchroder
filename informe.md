```
my-project/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── analytics/
│   │   ├── database.go
│   │   ├── market.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── user.go
│   │
│   ├── controllers/
│   │   └── controller.go
│   └── models/
│       └── model.go
│
│
├── docs/
│   └── swagger_files.go
│
├── go.mod
└── go.sum
```

## Postgress
In your own console:
```
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo service postgresql start
// to check status
sudo service postgresql status
```
Create database:
<pre>(base) <font color="#8AE234"><b>florxha@florxha-Inspiron-7375</b></font>:<font color="#729FCF"><b>~</b></font>$ sudo -i -u postgres
postgres@florxha-Inspiron-7375:~$ psql
psql (14.12 (Ubuntu 14.12-0ubuntu0.22.04.1))
Type &quot;help&quot; for help.

postgres=# CREATE USER florxha WITH PASSWORD &apos;mydb123&apos;;
CREATE ROLE
postgres=# CREATE DATABASE florxha_tp3;
CREATE DATABASE
postgres=# GRANT ALL PRIVILEGES ON DATABASE florxha_tp3 TO florxha;
GRANT
postgres=# \l+ florxha_tp3
                                                     List of databases
    Name     |  Owner   | Encoding |   Collate   |    Ctype    |   Access privileges   |  Size   | Tablespace | Description
-------------+----------+----------+-------------+-------------+-----------------------+---------+------------+-------------
 florxha_tp3 | postgres | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =Tc/postgres         +| 8577 kB | pg_default |
             |          |          |             |             | postgres=CTc/postgres+|         |            |
             |          |          |             |             | florxha=CTc/postgres  |         |            |
(1 row)
postgres=# exit
postgres@florxha-Inspiron-7375:~$ exit
logout
</pre>

In internal/analytics/database.go change the database data:
```
	databaseUrl := "postgres://username:password@localhost:5432/dbname" // Update with your DB details
	databaseUrl := "postgres://florxha:mydb123@localhost:5432/florxha_tp3"
```

To verify that the users are being added correctly when making a post from Swagger:
<pre>(base) <font color="#8AE234"><b>florxha@florxha-Inspiron-7375</b></font>:<font color="#729FCF"><b>~</b></font>$ psql postgres://florxha:mydb123@localhost:5432/florxha_tp3
psql (14.12 (Ubuntu 14.12-0ubuntu0.22.04.1))
SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, bits: 256, compression: off)
Type &quot;help&quot; for help.

florxha_tp3=&gt; \dt
        List of relations
 Schema | Name  | Type  |  Owner
--------+-------+-------+---------
 public | users | table | florxha
(1 row)

florxha_tp3=&gt; SELECT * FROM users;
 id | username | email | password
----+----------+-------+----------
(0 rows)

florxha_tp3=&gt; SELECT * FROM users;
 id | username |        email        |  password
----+----------+---------------------+-------------
  1 | johndoe  | example@example.com | password123
(1 row)

florxha_tp3=&gt; SELECT * FROM users;
 id | username |        email        |  password
----+----------+---------------------+-------------
  1 | johndoe  | example@example.com | password123
  2 | flor     | flor@example.com    | flor123
(2 rows)

florxha_tp3=&gt;
</pre>
