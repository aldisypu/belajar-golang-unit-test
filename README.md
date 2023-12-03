# belajar-golang-unit-test
Learn Go unit test by Programmer Zaman Now

MEMBUAT UNIT TEST

    ATURAN FILE TEST
        - harus menggunakan akhiran_test
        - misal file hello_world.go => hello_world_test.go

    ATURAN FUNCTION UNIT TEST
        - nama function harus diawali dengan nama TEST
        - misal function HelloWorld => TestHelloWorld
        - harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value

    MENJALANKAN UNIT TEST
        - go test
        - go test -v jika ingin lohat lebih detail function test apa saja yang sudah running
        - go test -v -run TestNamaFunction jika ingin memilih function unit test

    MENJALANKAN SEMUA UNIT TEST
        - go test -v ./...
        - go test ./...



### MENGGAGALKAN TEST
    - menggunakan panic bukanlah hal yang bagus
    - untuk menggagalkan unit test menggunakan testing.t
    - terdapat function Fail(), FailNow(), Error(), dan Fatal()

    t.Fail() dan t.FailNow()
        - Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test
        - FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

    t.Error(args...) dan t.Fatal(args...)
        - Error() lebih seperti melakukan log(print) error, namun setelah melakukan log error,
        dia akan secara otomatis memanggil Fail()
        - Fatal() mirip dengan Error(), setelah melakukan log error, dia akan memanggil FailNow()



### ASSERTION
    - sangat disarankan menggunakan ASSERTION untuk melakukan pengecekan
    - Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh
    menambahkan library untuk melakukan assertion ini

    TESTIFY
        - library assertion yang paling populer di Go-Lang
        - https://github.com/stretchr/testify
        - go get github.com/stretchr/testify

    ASSERT VS REQUIRE
        - testify menyediakan dua package untuk assertion, yaitu assert dan require
        - assert, jika gagal, maka assert akan memanggil Fail()
        - require, jika gagal, maka require akan memanggil FailNow()



### SKIP TEST
    - membatalkan eksekusi unit test jika kita mau
    - untuk membatalkan unit test kita bisa menggunakan function Skip()

### BEFORE DAN AFTER TEST
    - di Go-Lang terdapat fitur yang bernama testing.M
    - fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga
    bisa kita gunakan untuk melakukan Before dan dan After di unit test

    testing.M
    - cukup membuat function bernama TestMain dengan parameter testing.M
    - ingat function TestMain itu dieksekusi hanya sekali per Go-Lang package,
    bukan pertiap function unit test



### SUB TEST
    - Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
    - membuat sub test, menggunakan function Run()

    MENJALANKAN HANYA SUB TEST
        - go test -run TestNamaFunction/NamaSubTest
        - untuk semua test semua sub test di semua function
        go test -run /NamaSubTest

### TABLE TEST
    - table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan
    ekspektasi hasil dari unit test
    - lalu slice tersebut kita iterasi menggunakan sub test

### MOCK
    - mock adalah object yang sudah kita program dengan ekspektasi terrentu sehingga ketika
    dipanggil, dia akan menghasilkan data yang suda kita program diawali
    - membuat mock object dari suatu object yang memang sulit untuk di testing
    - misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil
    API Call ke third party service

    TESTIFY MOCK
        - jika desain kode program kita jelek, akan sulit untuk melakukan mocking



### BENCHMARK
    - menghitung kecepatan performa kode aplikasi
    - testing.B

    testing.B
        - struck yang digunakan untuk melakukan benchmark
        - testing.B mirip testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dll
        - yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
        - salah satunya adalah attribute N, digunakan untuk melakukan total iterasi sebuah benchmark
    
    BENCHMARK FUNCTION
        - nama function BenchmarkHelloWorld
        - memiliki parameter (b *testing.B)
        - dan tidak boleh mengembalikan return value
        - nama file diakhiri dengan _test, misal hello_world_test.go

    MENJALANKAN BENCHMARK
        - Untuk menjalankan seluruh benchmark di module
        go test -v -bench=. 
        - hanya ingin menjalankan benchmark tanpa unit test
        go test -v -run=NotMathUnitTest -bench=.
        - menjalankan benchmark tertentu
        go test -v -crun=NotMathUnitTest -bench=BenchmarkTest
        - menjalankan benchmark di root module dan ingin semua module dijalankan
        go test -v -bench=. ./...

### SUB BENCHMARK
    -menggunakan function Run()

    MENJALANKAN HANYA SUB BENCHMARK
        - go test -v -bench=BenchmarkNama/NamaSub

### TABLE BENCHMARK
