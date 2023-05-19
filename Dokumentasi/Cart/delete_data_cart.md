# Delete Data Cart

Menghapus data cart

**URL** : `/carts/:id`

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
