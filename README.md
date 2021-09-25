```
$ packer build template.json 

    yandex: output will be in this color.

    ==> yandex: Creating temporary ssh key for instance...
    ==> yandex: Using as source image: fd8nr7rmmtu5mkmkag3h (name: "ubuntu-20-04-lts-v20210924", family: "ubuntu-2004-lts")
    ==> yandex: Creating network...
    ==> yandex: Creating subnet in zone "ru-central1-a"...
    ==> yandex: Creating disk...

        ...

    ==> Wait completed after 3 minutes 16 seconds

    ==> Builds finished. The artifacts of successful builds are:
    --> yandex: A disk image was created: ubuntu-2004-lts-nginx-2021-09-25t10-10-43z (id: fd85jq9dc1ls2ou545br) with family name ubuntu-web-server

$ yc compute image list

    +----------------------+--------------------------------------------+-------------------+----------------------+--------+
    |          ID          |                    NAME                    |      FAMILY       |     PRODUCT IDS      | STATUS |
    +----------------------+--------------------------------------------+-------------------+----------------------+--------+
    | fd85jq9dc1ls2ou545br | ubuntu-2004-lts-nginx-2021-09-25t10-10-43z | ubuntu-web-server | f2eitbc7hc8l5ino1sim | READY  |
    +----------------------+--------------------------------------------+-------------------+----------------------+--------+

$ yc compute image delete fd85jq9dc1ls2ou545br

    done (8s)
```
