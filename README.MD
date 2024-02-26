```markdown
myclass/
├── cmd/
│   └── biodata/
│       └── main.go
├── internal/
│   ├── entities/
│   │   └── friend.go
│   ├── usecase/
│   │   └── friend_usecase.go
│   ├── presenter/
│   │   └── friend_presenter.go
│   └── repositories/
│       └── friend_repository.go
└── go.mod

# Friend CLI

Friend CLI adalah sebuah aplikasi command-line interface (CLI) sederhana yang ditulis menggunakan bahasa pemrograman Go (Golang) yang memungkinkan Anda untuk menampilkan informasi tentang teman-teman Anda dalam format tabel.

## Instalasi

Sebelum menggunakan aplikasi ini, pastikan Anda telah menginstal Go di komputer Anda. Anda dapat mengunduh dan menginstal Go dari situs web resmi Go: [https://golang.org/dl/](https://golang.org/dl/)

Untuk menggunakan Friend CLI, Anda perlu memiliki Go terinstal di komputer Anda. Kemudian, Anda dapat menginstalnya menggunakan perintah berikut:

```bash
go get -u github.com/usernameanda/friend-cli
```

## Penggunaan

Setelah instalasi, Anda dapat menggunakan CLI untuk menampilkan informasi tentang teman-teman Anda dengan memberikan nomor absen mereka sebagai argumen. Berikut adalah cara menggunakannya:

```bash
friend-cli <absen>
```

Gantilah `<absen>` dengan nomor absen teman Anda.

## Contoh

Misalkan Anda ingin menampilkan informasi tentang seorang teman dengan nomor absen 1. Anda akan menjalankan perintah berikut:

```bash
 go run main.go 1
```

Ini akan menghasilkan informasi teman dalam format tabel:

```
+-------+--------+----------------------+-----------------+---------------------------------------+
| Absen |  Nama  |        Alamat        |    Pekerjaan    |          Alasan Memilih Golang        |
+-------+--------+----------------------+-----------------+---------------------------------------+
|   1   | Calman | Jalan Kehidupan No. 1| Software Engineer| Ingin belajar bahasa pemrograman Go  |
+-------+--------+----------------------+-----------------+---------------------------------------+
```# myclass
