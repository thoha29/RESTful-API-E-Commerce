# Create Data Cart

Menambahkan produk ke keranjang

**URL** : `/carts`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "product_id": "[product id in int]",
    "qty": "[quantity in int]"
}
      
```

**Data example**

```json
{
    "product_id": 1,
    "qty": 5
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
"Product added to cart"
```

## Error Response
**Condition** : Jika tidak ter-autentikasi

**Code** : `401 Unauthorized`

**Content** :

```json
{
    "message": "invalid or expired jwt"
}
```

**Condition** : Jika yang dimasukan ada yang kosong

**Code** : `400 Bad Request`

**Content** :

```json
"Invalid input"
```

Kembali ke halaman [awal](../README.md)
