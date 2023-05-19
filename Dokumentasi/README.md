# Dokumentasi RESTAPI

Terdapat deberapa fitur-fitur yang diperioritaskan kepada user biasa/ pelanggan (bukan admin) yaitu:
- Pelanggan dapat melihat daftar produk berdasarkan kategori produk
- Pelanggan dapat menambahkan produk ke keranjang belanja
- Pelanggan dapat melihat daftar produk yang telah ditambahkan ke keranjang belanja
- Pelanggan dapat menghapus daftar produk di keranjang belanja
- Pelanggan dapat checkout dan melakukan transaksi pembayaran
- Login dan daftar pelanggan

Selain fungsi-fungsi diatas saya juga menambahkan fungsi-fungsi CRUD yang lain untuk melengkapi keseluruhan fungsi yang dibutuhkan.   

## RESTful API
API dikelompokan menjadi 5 sesuai dari banyaknya tabel yang ada di database.
Untuk mengakses API digunakan `http://localhost:1234/(API)`

## USER
register dan login tidak menggunakan Authentication sisanya menggunakan auth.
- [Login](User/login.md): `POST /login`
- [Register](User/registrasi.md): `POST /users`

Dari sini samapai API berikutnya kebawah semuannya menggunakan auth.
- [Add all user](User/read_data_user.md): `GET /users`
- [Update user](User/update_data_user.md): `PUT /users/:id`
- [Delete user](User/delete_data_user.md): `DELETE /users/:id`

***NOTE*** id disini merupakan user_id

## PRODUCT
API Product disini pengaksesnya lebih diperuntukan untuk Admin yang memegang CMS
- [Create product](Products/create_data_product.md): `POST /products`
- [Add all product](Products/read_data_product.md): `GET /products`
- [Update product](Products/update_data_product.md): `PUT /products/:id`
- [Delete product](Products/delete_data_product.md): `DELETE /products/:id`

***NOTE*** id disini merupakan product_id

## GENRE
API Genre disini pengaksesnya lebih diperuntukan untuk Admin yang memegang CMS
- [Create genre](Genre/create_data_genre.md): `POST /genres`
- [Add all genre](Genre/read_data_genre.md): `GET /genres`
- [Update genre](Genre/update_data_genre.md): `PUT /genres/:id`
- [Delete genre](Genre/delete_data_genre.md): `DELETE /genres/:id`

***NOTE*** id disini merupakan genre_id

## CART
API Cart disini pengaksesnya lebih diperuntukan untuk User/pelanggan
- [Create cart](Cart/create_data_cart.md): `POST /carts`
- [Add all cart sesuai user yang login](Cart/read_data_cart.md): `GET /carts`
- [Update qyt product di cart](Cart/update_data_cart.md): `PUT /carts/:id`  
- [Delete cart](Cart/delete_data_cart.md): `DELETE /carts/:id`

***NOTE*** id disini merupakan cart_id


## ORDER
API Order disini ada yang dipakai untuk User dan Admin
- [Create order oleh user](Order/create_data_order.md): `POST /orders`
- [Add all order sesuai user yang login](Order/read_data_order.md): `GET /orders`
- [Update status order oleh User/Admin](Order/update_data_order.md): `PUT /ordes/:id`

***NOTE*** id disini merupakan order_id
