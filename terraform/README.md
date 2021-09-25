### Провайдер

Начнем с манифеста, который будет описывать к какому облачному провайдеру мы планируем обратиться. Создадим файл provider.tf со следующим содержимым:

```
    provider "yandex" {
        token     = var.yc_token
        cloud_id  = var.yc_cloud
        folder_id = var.yc_folder
    }

    terraform {
        required_providers {
            yandex = {
            source = "yandex-cloud/yandex"
            }
        }
    }
```

Здесь мы указываем, что провайдер у нас Yandex.Cloud

```
    provider "yandex" {
    ...
```

И передаем необходимые переменные для использования данного провайдера, токен, ID облака и ID каталога. Файл со значениями этих переменных мы создадим позже. Далее, в этом же файле, мы указываем дополнительный блок, который будет использовать терраформ:

```
    terraform {
    required_providers {
        yandex = {
        source = "yandex-cloud/yandex"
        }
    }
```

Теперь нам надо создать файл variables.tf, в которым мы опишем какие переменные мы будем использовать в наших манифестах:

```
    variable "yc_cloud" {
        type = string
        description = "Yandex Cloud ID"
    }

    variable "yc_folder" {
        type = string
        description = "Yandex Cloud folder"
    }

    variable "yc_token" {
        type = string
        description = "Yandex Cloud OAuth token"
    }

    variable "db_password" {
        description = "PostgreSQL user pasword"
    }
```

А их значения мы укажем в файле wp.auto.tfvars. Содержимое этого файла следующее (!!! Указанные ниже значения приведены для примера !!!):

```
    yc_cloud  = "b1gf5768rgabjbptan7a"
    yc_folder = "b1gb8haadbndninaj928"
    yc_token = "jlkdsflgjoisgoskljgs"
    db_password = "password"
```

Напомню, что эти ID облака и каталога вы запомнили/записали при создании каталога в YC. Чтобы узнать вам токен, выполните команду:
```
    $ yc config list
```

Если вы корректно настроили интерфейс командной строки Yandex.Cloud (CLI), согласно ссылке https://cloud.yandex.ru/docs/cli/quickstart, то, в качестве бонуса, кроме значения токена эта утилита покажет вам и ID облака и каталога. И так, на текущий момент структура вашего репозитория должна быть следующей:

```
    .
    └── terraform
        ├── provider.tf
        ├── variables.tf
        └── wp.auto.tfvars
```

И уже сейчас мы можем проинициализировать терраформ при помощи следующей команды:

```
    $ terraform init

    Initializing the backend...

    Initializing provider plugins...
    - Finding latest version of yandex-cloud/yandex...
    - Installing yandex-cloud/yandex v0.63.0...
    - Installed yandex-cloud/yandex v0.63.0 (self-signed, key ID E40F590B50BB8E40)

    Terraform has been successfully initialized!
```

Теперь мы можем переходить к созданию манифестов для наших ресурсов.

### Виртуальные сети

Сначала создадим виртуальную сеть. Для этого мы будем использовать файл network.tf со следующим содержимым:

```
resource "yandex_vpc_network" "wp-network" {
  name = "wp-network"
}

resource "yandex_vpc_subnet" "wp-subnet-a" {
  name = "wp-subnet-a"
  v4_cidr_blocks = ["10.2.0.0/16"]
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.wp-network.id
}

resource "yandex_vpc_subnet" "wp-subnet-b" {
  name = "wp-subnet-b"
  v4_cidr_blocks = ["10.3.0.0/16"]
  zone           = "ru-central1-b"
  network_id     = yandex_vpc_network.wp-network.id
}

resource "yandex_vpc_subnet" "wp-subnet-c" {
  name = "wp-subnet-c"
  v4_cidr_blocks = ["10.4.0.0/16"]
  zone           = "ru-central1-c"
  network_id     = yandex_vpc_network.wp-network.id
}
```

Обратите внимание, что здесь мы создаем ресурс типа yandex_vpc_network и три ресурса с подсетями yandex_vpc_subnet, которые будут располагаться в разных зонах (помним про отказоустойчивость!). Соответственно, в ресурсе yandex_vpc_subnet мы указываем блок IP-адресов, зону, где будет расположена данная подсеть и ссылку на саму виртуальную сеть. Давайте проверим как работает данный манифест, выполним команду:

```
    $ terraform apply --auto-approve
```

Флаг --auto-approve мы используем, чтобы не тратить лишнее время на подтверждение запроса о применении манифеста. Если в манифесте мы не допустили ошибок, то после исполнения команды, последним сообщением мы увидим:

```
    Apply complete! Resources: 4 added, 0 changed, 0 destroyed.
```

Действительно, мы создали четыре новых ресурса - одну виртуальную сеть и три ее подсети.

### Виртуальные машины

Следующая задача - манифесты для хостов, где будет разворачиваться WordPress. Создадим файл wp-app.tf со следующим содержимым:

```
    resource "yandex_compute_instance" "wp-app-1" {
        name = "wp-app-1"
        zone = "ru-central1-a"

    resources {
        cores  = 2
        memory = 2
    }

    boot_disk {
        initialize_params {
            image_id = "fd80viupr3qjr5g6g9du"
        }
    }

    network_interface {
        # Указан id подсети default-ru-central1-a
        subnet_id = yandex_vpc_subnet.wp-subnet-a.id
        nat       = true
    }

    metadata = {
        ssh-keys = "ubuntu:${file("~/.ssh/yc.pub")}"
    }
    }

    resource "yandex_compute_instance" "wp-app-2" {
        name = "wp-app-2"
        zone = "ru-central1-b"

    resources {
        cores  = 2
        memory = 2
    }

    boot_disk {
        initialize_params {
        image_id = "fd80viupr3qjr5g6g9du"
        }
    }

    network_interface {
        # Указан id подсети default-ru-central1-b
        subnet_id = yandex_vpc_subnet.wp-subnet-b.id
        nat       = true
    }

    metadata = {
        ssh-keys = "ubuntu:${file("~/.ssh/yc.pub")}"
    }
    }
```

Здесь мы создаем два ресурса вида yandex_compute_instance - т.е. две виртуальные машины. Для них в блоке resources мы указываем кол-во ядер, памяти. В блоке boot_disk ID образа с операционной системы (в данном случае это Ubuntu 18.04). В блоке network_interface мы указываем подсеть, в которой будет располагаться виртуальная машина. И, наконец, в блоке metadata передается публичная часть ключа для подключения через SSH. Вы, разумеется, можете использовать имеющийся у вас ключ. Снова запустим команду применения манифестов и убедимся, что виртуальные машины созданы успешно:

```
    $ terraform apply --auto-approve
    ...

    Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
```

### Балансировщик трафика

Итак, виртуальные машины у нас есть, теперь мы можем создать балансировщик, который будет перенаправлять на них пользовательский трафик. Создадим манифест lb.tf со следующим содержимым:

```
    resource "yandex_lb_target_group" "wp_tg" {
    name      = "wp-target-group"

    target {
        subnet_id = yandex_vpc_subnet.wp-subnet-a.id
        address   = yandex_compute_instance.wp-app-1.network_interface.0.ip_address
    }

    target {
        subnet_id = yandex_vpc_subnet.wp-subnet-b.id
        address   = yandex_compute_instance.wp-app-2.network_interface.0.ip_address
    }
    }

    resource "yandex_lb_network_load_balancer" "wp_lb" {
    name = "wp-network-load-balancer"

    listener {
        name = "wp-listener"
        port = 80
        external_address_spec {
        ip_version = "ipv4"
        }
    }

    attached_target_group {
        target_group_id = yandex_lb_target_group.wp_tg.id

        healthcheck {
        name = "http"
        http_options {
            port = 80
            path = "/health"
        }
        }
    }
    }
```

В данном манифесте мы, во-первых, создаем группу хостов, куда будем направлять трафик, при помощи ресурса yandex_lb_target_group. В нем мы ссылаемся на IP-адреса созданных ранее виртуальных машин. Далее мы создаем сам балансировщик при помощи ресурса yandex_lb_network_load_balancer. В блоке listener мы указываем порт, который будет слушать балансировщик. А в блоке attached_target_group указывается ссылка на группу хостов yandex_lb_target_group. Запустим команду применения манифестов и убедимся, что виртуальные машины созданы успешно:

```
    $ terraform apply --auto-approve
    ...

    Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
```

В данном случае, два новых ресурса - это сам балансировщик и группа хостов, на которые он направляет трафик.

### База данных

Предполагается, что в качестве бекэнда для WordPress-а мы будем использовать PostgreSQL. Воспользуемся для этой цели возможностью создать в YC облачный кластер PostgreSQL и опишем для этого соответствующий манифест. Назовем его db.tf. Содержимое манифеста будет:

```
    locals {
        dbuser = tolist(yandex_mdb_postgresql_cluster.wp_postgresql.user.*.name)[0]
        dbpassword = tolist(yandex_mdb_postgresql_cluster.wp_postgresql.user.*.password)[0]
        dbhosts = yandex_mdb_postgresql_cluster.wp_postgresql.host.*.fqdn
        dbname = tolist(yandex_mdb_postgresql_cluster.wp_postgresql.database.*.name)[0]
    }

    resource "yandex_mdb_postgresql_cluster" "wp_postgresql" {
        name        = "wp-postgresql"
        folder_id   = var.yc_folder
        environment = "PRODUCTION"
        network_id  = yandex_vpc_network.wp-network.id

    config {
        version = 12
        resources {
        resource_preset_id = "s2.small"
        disk_type_id       = "network-ssd"
        disk_size          = 64
        }
    }

    database {
        name  = "db"
        owner = "user"
    }

    user {
        name     = "user"
        password = var.db_password
        permission {
        database_name = "db"
        }
    }

    host {
        zone      = "ru-central1-b"
        subnet_id = yandex_vpc_subnet.wp-subnet-b.id
        assign_public_ip = true
    }
    host {
        zone      = "ru-central1-c"
        subnet_id = yandex_vpc_subnet.wp-subnet-c.id
        assign_public_ip = true
    }
    }
```

Кластер баз данных мы создаем при помощи ресурса yandex_mdb_postgresql_cluster. Из основного в нем стоит обратить внимание на блок config, где задается версия PostgreSQL и указываются характеристики узлов, где будет развернут кластер. В блоках database и user мы задаем имя базы и имя с паролем для пользователя базы, соответственно. В блоках host указываются подсети, где будут размещены узлы кластера. Итак, запустим команду применения манифестов и убедимся, что кластер баз данных создан успешно. На этот раз придется подождать около семи минут, создание кластера - дело не быстрое.

```
    $ terraform apply --auto-approve
    ...

    yandex_mdb_postgresql_cluster.wp_postgresql: Creation complete after 7m17s [id=c9qgpnu68l0pjanrttlf]
    Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

### Output-переменные

И в завершении давайте создадим файл output.tf, где укажем вывод некоторой информации, которая может нам пригодится. Содержимое данного манифеста:

```
    output "load_balancer_public_ip" {
        description = "Public IP address of load balancer"
        value = yandex_lb_network_load_balancer.wp_lb.listener.*.external_address_spec[0].*.address
    }

    output "database_host_fqdn" {
        description = "DB hostname"
        value = local.dbhosts
    }
```

Как нетрудно заметить, вы запрашиваем вывод IP балансировщика и dns-имена баз данных. Мы можем запустить команду terraform apply еще раз и, хотя в этот раз никаких новых ресурсов мы не создали, но мы увидим вывод запрошенной информации (IP и имена баз данных у вас будут свои):

```
    $ terraform apply --auto-approve
    ...
    Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

    Outputs:

    database_host_fqdn = tolist([
        "rc1b-xveb6hht7i0a1luv.mdb.yandexcloud.net",
        "rc1c-u2r99z9ektqz4zq8.mdb.yandexcloud.net",
    ])
        load_balancer_public_ip = tolist([
        "84.201.161.168",
    ])
```
