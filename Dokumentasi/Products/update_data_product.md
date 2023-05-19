# Update Product

Mengubah data product

**URL** : `/products/:id`

**Method** : `PUT`

**Auth required** : YES

**Data example**

```json
{
    "product_name": "One Pice",
    "price": 50000,
    "stock": 20,
    "genre_id": 1
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
"Data can't be empty"
```
**Condition** : Jika product tidak ditemukan

**Code** : `404 StatusNotFound`

**Content** :

```json
{
    "status": 404,
    "message": "Product not found",
    "data": null
}
```

Kembali ke halaman [awal](../README.md)
