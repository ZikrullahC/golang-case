# Go Case Öğrenci Yapılacaklar Listesi

Bu proje, Go dilinde yazılmış bir web uygulamasıdır. Öğrenci ve plan bilgilerini yönetmek için bir RESTful API sağlar. Projede, Echo web framework'ü ve GORM ORM kullanılmıştır.

## Başlarken

### Gereksinimler

- [Go 1.16+](https://golang.org/doc/install)
- [MySQL](https://www.mysql.com/downloads/)

### Kurulum

1. Bu projeyi yerel makinenize klonlayın:

    ```sh
    git clone (https://github.com/ZikrullahC/golang-case)
    cd golang-case
    ```

2. Gerekli Go bağımlılıklarını yükleyin:

    ```sh
    go mod download
    ```
   

3. MySQL veritabanınızı yapılandırın ve `dbBaglantisi` değişkenini `db/connect.go` dosyasında güncelleyin:

    ```go
    dbBaglantisi := "username:password@tcp(127.0.0.1:3306)/vatan_soft_go_staj_case?charset=utf8mb4&parseTime=True&loc=Local"
    ```
4. Echo Framework kurulumu için:
    ```sh
    go get -u github.com/labstack/echo/v4
    ```
5. Gorm Framework Kurulumu için:
   ```sh
   go get -u gorm.io/gorm
   ```
   MySQL sürücüsünü kullanmak için:
   ```sh
   go get -u gorm.io/driver/mysql
   ```
   

## Kullanım

### Sunucuyu Başlatma

Sunucuyu başlatmak için aşağıdaki komutu kullanın:

```sh
go run main.go
```

Sunucu http://localhost:8080 adresinde çalışacaktır.

# API Endpoints

## Öğrenci Endpoints
- **POST /ogrenci**: Yeni bir öğrenci oluşturur.
- **GET /ogrenci/:id**: Belirtilen ID'ye sahip öğrencinin bilgilerini alır.
- **PUT /ogrenci/:id**: Belirtilen ID'ye sahip öğrencinin bilgilerini günceller.
- **DELETE /ogrenci/:id**: Belirtilen ID'ye sahip öğrenciyi siler.
- **GET /ogrenci**: Tüm öğrencilerin bilgilerini alır.

## Plan Endpoints
- **POST /plan**: Yeni bir plan oluşturur.
- **GET /plan/:id**: Belirtilen ID'ye sahip planın bilgilerini alır.
- **PUT /plan/:id**: Belirtilen ID'ye sahip planın bilgilerini günceller.
- **DELETE /plan/:id**: Belirtilen ID'ye sahip planı siler.
- **GET /plan/ogrenci/:id**: Belirtilen öğrenci ID'sine ait tüm planları listeler.

## Middleware
- **Logger Middleware**: Gelen istekler ve sunucu tarafından gönderilen yanıtlar hakkında bilgi içeren günlük kayıtları oluşturur.
- **Recover Middleware**: Bir hata oluştuğunda sunucuyu çökmekten korur ve kullanıcılara uygun bir hata mesajı gösterir.

## Proje Yapısı
- **db**: Veritabanı bağlantısı ve migrasyonu.
- **handler**: API endpoint fonksiyonları.
- **model**: Veritabanı modelleri.
