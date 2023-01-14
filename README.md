# Enrd-System
- Estimation of network resource status for Dataflow platform management

# Try
```
$ tinet upconf | sudo sh -x
```

- `docker exec -it compute1 bash`
    - Execute on all compute nodes
    ```
    root@compute1:/go/enrd# ./bin/agent net0 fc00:1::/64 
    ```
    > Note: Excute after the database is built.

- `docker exec -it controller bash`
    ```
    root@controller:/go/enrd# service mysql start
    root@controller:/go/enrd# mysql
    MariaDB [(none)]> CREATE DATABASE enrd;
    MariaDB [(none)]> GRANT ALL PRIVILEGES ON enrd.* TO 'enrd'@'localhost' IDENTIFIED BY 'PASSWORD';
    MariaDB [(none)]> GRANT ALL PRIVILEGES ON enrd.* TO 'enrd'@'%' IDENTIFIED BY 'PASSWORD';
    MariaDB [(none)]> flush privileges;
    MariaDB [(none)]> exit
    root@controller:/go/enrd# ./bin/controller init
    root@controller:/go/enrd# ./bin/controller config -c templates/config.yaml 
    root@controller:/go/enrd# ./bin/controller estimate -c temptates/param.yaml
    ```
