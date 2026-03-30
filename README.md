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

## Hati-Hati Saat Membuat Array
- Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice


### contoh di file:
- [slice.go](slice.go)

--- 


# Tipe Data Map
- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
- Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

## Function Map
- len(map)                    | Untuk mendapatkan jumlah data di map
- map[key]                    | Mengambil data di map dengan key
- map[key] = value            | Mengubah data di map dengan key
- make(map[TypeKey]TypeValue) | Membuat map baru
- delete(map, key)            | Menghapus data di map dengan key


### contoh di file:
- [map.go](map.go)

--- 

# If Expression
- If adalah salah satu kata kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir di semua bahasa pemrograman mendukung if expression

## Else Expression
- Blok if akan dieksekusi ketika kondisi if bernilai true
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
- Hal ini bisa dilakukan menggunakan else expression

## Else If Expression 
- Kadang dalam If, kita butuh membuat beberapa kondisi
- Kasus seperti ini, kita bisa menggunakan Else If expression


## If Short Statement
- If mendukung short statement sebelum kondisi
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi



### contoh di file:
- [if_expression.go](if_expression.go)

--- 

# Switch Expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch  expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

## Swicth Short Statement
- Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

## Swicth tanpa kondisi
- Kondisi di switch expression tidak wajib
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya


### contoh di file:
- [switch.go](switch.go)

--- 

# For Loops
- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
- Salah satu fitur perulangan adalah for loops

## For dengan Statement
- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
- Init statement, yaitu statement sebelum for di eksekusi
- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

## For Range
- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
- Data collection contohnya Array, Slice dan Map


### contoh di file:
- [for.go](for.go)

--- 

# Break & Continue
- Break & continue adalah kata kunci yang bisa digunakan dalam perulangan 
- Break digunakan untuk menghentikan seluruh perulangan
- Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya


### contoh di file:
- [break.go](break.go)
- [continue.go](continue.go)

--- 

# Function
- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main
- Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
- Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti dengan nama function nya dan blok kode isi function nya
- Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya diikuti tanda kurung buka, kurung tutup


### contoh di file:
- [function.go](function.go)


--- 

# Function Parameter
- Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
- Kita bisa menambahkan parameter di function, bisa lebih dari satu
- Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya


### contoh di file:
- [function_parameter.go](function_parameter.go)

--- 

# Function Return Value
- Function bisa mengembalikan data
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
- Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya



### contoh di file:
- [function_return_value.go](function_return_value.go)

--- 

# Returning Multiple Values
- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function


### contoh di file:
- [return_multiple_values.go](return_multiple_values.go)

--- 

# Named Return Values
- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variable secara langsung di tipe data return function nya


### contoh di file:
- [named_return_value.go](named_return_value.go)

--- 

# Variadic Function
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array?
      - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
      - JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma



### contoh di file:
- [variadic_function.go](variadic_function.go)

--- 





