# Update Qyt Product in Cart

Mengubah quantity produk yang diinginkan didalam keranjang

**URL** : `/carts/:id`

**Method** : `PUT`

**Auth required** : YES

**Data example**

```json
{
    "product_id": 1
    "qty": 11
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "status": 200,
    "message": "Success",
    "data": {
        "rowsAffected": 1
    }
}
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

**Condition** : Jika data yang dimasukan kosong

**Code** : `400 Bad Request`

**Content** :

```json
"Invalid input"
```
**Condition** : Jika cart tidak ditemukan

**Code** : `404 StatusNotFound`

**Content** :

```json
{
    "status": 404,
    "message": "Cart not found",
    "data": null
}
```

Kembali ke halaman [awal](../README.md)
