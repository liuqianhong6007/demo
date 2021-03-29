# minikube

- 当前 shell 将 minikube 的docker 环境映射出来： eval $(minikube docker-env)
- 此时 docker images 查看到的镜像就是 minikube 的私人镜像仓库
- docker build getip:latest  编译镜像，镜像会保存在 minikube 镜像仓库
- 这样 kubernetes job.yaml 就能拉取到镜像了

