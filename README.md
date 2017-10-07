# Binar Assessment
#### By Reza Andriyunanto
### 1. Stack Backend Design
![Stack Backend Design](https://raw.githubusercontent.com/rezandry/binarassessment/master/image/Arc.jpg)
* #### Arsitektur 
* [x] Alasan dari arsitektur tersebut adalah, dimana proses client dan server terpisah, sehingga proses client akan lebih ringan dan beban lebih berada di server karena proses akan dilakukan diserver dengan media perantara data dengan bentuk json.
* #### Tools
* [x] Bahasa menggunakan golang karena dia memiliki proses yang lebih cepat, dan pada case tertentu bisa diimplementasikan concurrency yang bisa mempercepat proses yang tidak memiliki dependensi.
* [x] IDE menggunakan VSCode karena dia menyediakan plugins yang cukup lengkap, terutama go, otomatis import package dan sangat dimudahkan dengan adanya terminal built-in.
* [x] Database menggunakan postgres karena relational database dan katanya lebih baik dalam beberapa case daripada mysql

### 2. Keamanan Pengiriman data
* #### Mobile Apps
* [x] Diperlukan adanya sistematika, dimana ketika data dikirim, dilakukan encrypt terlebih dahulu dan ketika proses yang membutuhkan authorisasi, perlu ditambahkan semacam token untuk memastikan user tersebut memiliki wewenang dari data tersebut
* #### Server
* [x] Diperlukan pengecekan token untuk mengetahui hak akses dari pengirim request bisa berupa token, serta pengecekan data sesuai dengan yang dikoordinasikan dengan mobile developer sehingga data yang diterima server adalah data asli, bukan data yang bisa jadi dirubah oleh man in the middle

### 3. Code
* Menggunakan Gin Framework, Postgres dan Go language
* _Unit Testing Belum Terimplementasi_
