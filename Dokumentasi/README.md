# Dokumentasi RESTAPI

Terdapat deberapa fitur-fitur yang diperioritaskan kepada user biasa/ pelanggan (bukan admin) yaitu:
- Pelanggan dapat melihat daftar produk berdasarkan kategori produk
- Pelanggan dapat menambahkan produk ke keranjang belanja
- Pelanggan dapat melihat daftar produk yang telah ditambahkan ke keranjang belanja
- Pelanggan dapat menghapus daftar produk di keranjang belanja
- Pelanggan dapat checkout dan melakukan transaksi pembayaran
- Login dan daftar pelanggan

Selain fungsi-fungsi diatas saya juga menambahkan fungsi-fungsi lain mengikuti CRUD untuk melengkapi keseluruhan fungsi yang dibutuhkan.   

## RESTful API
API dikelompokan menjadi 5 sesuai dari banyaknya tabel yang ada di database.
Untuk mengakses API digunakan `http://localhost:1234/(API)`

## USER
register dan login tidak menggunakan Authentication sisanya menggunakan auth.
- [Login](User/login.md): `POST /login`
- [Register](User/register.md): `POST /users`

Dari sini samapai API berikutnya kebawah semuannya menggunakan auth.
- [Add all user](User/show_data_user.md): `GET /users`
- [Update user](User/update_data_user.md): `PUT /users/:id`
- [Delete user](User/delete_data_user.md): `DELETE /users/:id`

***NOTE*** id disini merupakan user_id

## PRODUCT
API Product disini pengaksesnya lebih diperuntukan untuk Admin yang memegang CMS
- [Create product](Products/create_data_product): `POST /products`
- [Add all product](Products/read_data_product): `GET /products`
- [Update product](Products/update_data_product): `PUT /products/:id`
- [Delete product](Products/delete_data_product): `DELETE /products/:id`

***NOTE*** id disini merupakan product_id

## GENRE
API Genre disini pengaksesnya lebih diperuntukan untuk Admin yang memegang CMS
- [Create genre](Genre/create_data_genre): `POST /genres`
- [Add all genre](Genre/read_data_genre): `GET /genres`
- [Update genre](Genre/update_data_genre): `PUT /genres/:id`
- [Delete genre](Genre/delete_data_genre): `DELETE /genres/:id`

***NOTE*** id disini merupakan genre_id

## CART
API Cart disini pengaksesnya lebih diperuntukan untuk User/pelanggan
- [Create cart](Cart/create_data_product): `POST /carts`
- [Add all cart sesuai user yang login](Cart/read_data_product): `GET /carts`
- [Update qyt product di cart](Cart/update_data_product): `PUT /carts/:id`  
- [Delete cart](Cart/delete_data_product): `DELETE /carts/:id`

***NOTE*** id disini merupakan cart_id


## ORDER
API Order disini ada yang dipakai untuk User dan Admin
- [Create order oleh user](Cart/create_data_product): `POST /orders`
- [Add all order sesuai user yang login](Cart/read_data_product): `GET /orders`
- [Update status order oleh User/Admin](Cart/update_data_product): `PUT /ordes/:id`
- [Delete order oleh user](Cart/delete_data_product): `DELETE /ordes/:id`

***NOTE*** id disini merupakan order_id
