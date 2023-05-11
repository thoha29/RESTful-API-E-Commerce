# RESTful-API-E-Commerce

## Deskripsi
Pada kesemoatan ini saya membuat Backend untuk e-commerce yang menjual komik menggunakan Golang dengan framework Echo. Didalamnya terdapat CRUD serta dilengkapi dengan fitur register dan login.

## Database
Disini saya menggunakan database MySQL. Nama database e_commerce dengan susunan tabel sebagai berikut:
#### ERD
![Golang](https://github.com/thoha29/RESTful-API-E-Commerce/assets/113443657/7bde615a-5cd7-4933-853a-235907f832d7)

Relasi antara tabel adalah One-to-Many:
- Tabel genres memiliki relasi one-to-many dengan tabel products, karena satu genre dapat memiliki banyak produk, namun satu produk hanya dapat memiliki satu genre.
- Tabel users memiliki relasi one-to-many dengan tabel carts, karena satu pengguna dapat memiliki banyak keranjang belanja, namun satu keranjang belanja hanya dapat dimiliki oleh satu pengguna.
- Tabel products memiliki relasi one-to-many dengan tabel carts, karena satu produk dapat ada di banyak keranjang belanja, namun satu keranjang belanja hanya dapat berisi satu produk yang sama.
- Tabel products memiliki relasi one-to-many dengan tabel orders, karena satu produk dapat ada di banyak pesanan, namun satu pesanan hanya dapat berisi satu produk yang sama.
- Tabel users memiliki relasi one-to-many dengan tabel orders, karena satu pengguna dapat membuat banyak pesanan, namun satu pesanan hanya dapat dibuat oleh satu pengguna.

## Dokumentasi
Untuk dokumentasi dapat dilihat [disini](Dokumentasi)
