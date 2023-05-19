# Delete Data Genre

Menghapus data genre

**URL** : `/genres/:id`

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

**Condition** : Jika genre tidak ditemukan

**Code** : `404 StatusNotFound`

**Content** :

```json
{
    "status": 404,
    "message": "genre not found",
    "data": null
}
```

Kembali ke halaman [awal](../README.md)
