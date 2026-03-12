# Slide
https://docs.google.com/presentation/d/1J0DbqyuLQVnGnkbL7bX3jL6iQc6RdXy8zQkfH8rbE0Q/edit

# Source code
https://github.com/ProgrammerZamanNow/belajar-golang-dasar-2023


# Membuat Project
- Project di Go-lang, biasanya disebut sebagai module
- Untuk membuat module, kita bisa menggunakan perintah berikut di folder tempat kita akan membuat module :
```bash
go mod init nama-module
```

### contoh di file:
- [go.mod](go.mod)

---

# Main Function
- Go-Lang, itu mirip seperti bahasa pemrograman C/C++, dimana perlu ada yang namanya main function
- Main function adalah sebuah fungsi yang akan dijalankan ketika program berjalan
- Untuk membuat function, kita bisa menggunan kata kunci func
- Main function harus terdapat di dalam main package
- Titik koma di Golang, tidaklah wajib, artinya kita bisa menambahkan titik koma atau tidak, diakhir kode program kita


# Println
- Untuk menulis tulisan, kita perlu melakukan import module fmt terlebih dahulu

# Perintah Compile project
```bash
go build
```

# Run project
```bash
.\pzn-belajar-golang-dasar.exe
```

# Menjalankan tanpa kompilasi
```bash
go run hello-world.go
```

### contoh di file:
- [hello-world.go](hello-world.go)
- [pzn-belajar-golang-dasar.exe](pzn-belajar-golang-dasar.exe)

---

# Multiple Main Function
- Di Golang, function dalam module / project adalah unik, artinya kita tidak boleh membuat nama function yang sama
- Oleh karena itu, jika kita membuat file baru, misal sample.go, lau membuat nama function yang sama yaitu main
- Maka kita tidak bisa melakukan build module, karena main functin tersebut duplikat dengan yang ada di main function hello-world.go

Jika di:
```bash
go build
```

muncul error:
```bash
PS C:\Users\owi\pzn-belajar-golang-dasar> go build
# pzn-belajar-golang-dasar
.\sample.go:5:6: main redeclared in this block
        .\hello-world.go:5:6: other declaration of main
```

# Solusinya?
- Karena Sekarang kita masih di dalam fase belajar, oleh karena itu kita tidak akan melakukan build project module terbelih dahulu
- Sekarang kita akan fokus menjalankan file Golang satu persatu, sehingga tidak akan terjadi error jika dijalankan file nya satu persatu
- Tapi INGAT, pada kenyataannya nanti, saat kita membuat project, kita hanya akan membuat satu main function saja

### contoh di file:
- [sample.go](sample.go)
- [hello-world.go](hello-world.go)

---

# Tipe Data Number
- Ada 2 jenis tipe data Number yaitu:
  - Integer (bilangan bulat)
  - Floating Point (bilangan desimal)

# Tipe Data Integer 
  - int8,16,32,64 (jika butuh negatif)
  - uint8,16,32,64 (jika tidak butuh negatif)

# Alias
 - byte = uint8
 - rune = int32
 - int = minimal int32
 - uint = mininaml uint32

### contoh di file:
- [number.go](number.go)

---

# Tipe Data Boolean
- tipe data boolean adalah tipe data yang memiliki dua nilai, yaitu benar atau salah
- Di Go-Lang, tipe data boolean direpresentasikan menggunakan kata kunci bool

### contoh di file:
- [boolean.go](boolean.go)

---

### Tipe Data String
- String adalah tipe data kumpulan karakter
- Jumlah karakter di dalam String bisa nol sampai tak terhingga
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
- Nilai data String di Go-Lang selalu diawali dengan karakter " dan di akhiri dengan "

### contoh di file:
- [string.go](string.go)