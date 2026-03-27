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

# Tipe Data String
- String adalah tipe data kumpulan karakter
- Jumlah karakter di dalam String bisa nol sampai tak terhingga
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
- Nilai data String di Go-Lang selalu diawali dengan karakter " dan di akhiri dengan "

# Function untuk String
- len("string") = menghitung jumlah karakter di String
- "string"[number] = mengambil karakter pada posisi yang di tentukan

### contoh di file:
- [string.go](string.go)

---

# Variable
- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

# Tipe Data Variable
- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

# Kata kunci var
- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib. 
- Asalkan saat membuat variable kita langsung menginisialisasi datanya
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut

# Deklarasi multiple variable
- Di Go-Lang kita bisa membuat variable secara sekaligus banyak
- Code yang dibuat akan lebih bagus dan mudah dibaca


### contoh di file:
- [variable.go](variable.go)

---

# Constant
- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

# Deklarasi Multiple Constant
- Sama seperti variable, di Go-Lang  juga kita bisa membuat constant secara sekaligus banyak


### contoh di file:
- [constant.go](constant.go)

---

# Konversi Tipe Data
- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
- Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

### contoh di file:
- [convertion.go](convertion.go)

---


# Type Declarations
- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti


### contoh di file:
- [type_declarations.go](type_declarations.go)

---

# Operasi Matematika

- (+) pertambahan
- (-) pengurangan
- (*) perkalian
- (/) pembagian
- (%) modulus atau sisa pembagian

## Augmented Assignments
- a = a + 10   |   a += 10
- a = a - 10   |   a -= 10
- a = a * 10   |   a *= 10
- a = a / 10   |   a /= 10
- a = a % 10   |   a %= 10

## Unary Operator
- (++) | a = a + 1
- (--) | a = a - 1
- (+) | Postive
- (-) | Negative
- (!) | kebalikan boolean

### contoh di file:
- [matematika.go](matematika.go)

---

# Operasi Perbandingan
- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
- Jika hasil operasinya adalah benar, maka nilainya adalah true
- Jika hasil operasinya adalah salah, maka nilainya adalah false


### contoh di file:
- [perbandingan.go](perbandingan.go)

---

# Operasi Boolean

- && = Dan
- || = Atau
- ! = Kebalikan


### contoh di file:
- [operasi_boolean.go](operasi_boolean.go)

---

# Tipe Data Array

- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
- Daya tampung Array tidak bisa bertambah setelah Array dibuat

## Function Array
- len(array) : Untuk mendapatkan panjang Array
- array[index] : Mendapat data di posisi index
- array[index] = value : Mengubah data di posisi index

### contoh di file:
- [array.go](array.go)

--- 

# Tipe Data Slice

- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

## Detail Tipe Slice
- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity


## Membuat Slice dari Array
- array[low:high] = Membuat slice dari array dimulai index low sampai index sebelum high
- array[low:] = Membuat slide dari array dimulai index low sampai index akhir di array
- array[:high] = Membuat slice dari array dimulai index 0 sampai index sebelum high
- array[:] = Membuat slice dari array dimulai index 0 sampai index akhir di array

## Function Slice
- len(slice) = Untuk mendapatkan panjang 
- cap(slice) = Untuk mendapat kapasitas
- append(slice, data) = Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
- make([]TypeData, length, capacity) = Membuat slice baru
- copy(destination, source) = Menyalin slice dari source ke destination


### contoh di file:
- [slice.go](slice.go)

--- 