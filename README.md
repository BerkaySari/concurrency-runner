# concurrency-runner

Runtime anında özellikle hataları işlerken stabil çalışmayan **go func** concurrent kullanımı yerine background context'inden bir errgroup nesnesi oluşturup sonuçları ve hataları işleyen kütüphanedir.

https://pkg.go.dev/golang.org/x/sync/errgroup


### Çalıştırmak için ###
```
go mod tidy
```
```
go run main.go examples.go
```