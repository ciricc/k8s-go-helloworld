# K8S Go (Привет, мир!)

Простое приложение с одним Docker образом, которое запускается в Kubernetes.

`./build.sh` - Создать образ
`./deploy.sh` - Задеплоить в k8s
`./connect.sh` - Настроить проброс порта, чтобы иметь возможность открыть сервис локально

После запуска сервис доступен по адресу [localhost:8080](https://localhost:8080/)