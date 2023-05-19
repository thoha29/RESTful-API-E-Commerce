# Delete User

Menghapus data product

**URL** : `/products/:id`

**Method** : `DELETE`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content**

```json
{
    "status": 200,
    "message": "Success",
    "data": {
        "product_id": 6,
        "rows_affected": 1
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
