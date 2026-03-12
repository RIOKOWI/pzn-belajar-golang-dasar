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
