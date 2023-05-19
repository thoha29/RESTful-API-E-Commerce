# Delete User

Menghapus data user

**URL** : `/users/:id`

**Method** : `DELETE`

**Auth required** : YES

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

**Condition** : Jika tidak terauterisasi

**Code** : `401 Unauthorized`

**Content** :

```json
{
    "message": "invalid or expired jwt"
}
```

**Condition** : Jika user tidak ditemukan

**Code** : `404 StatusNotFound`

**Content** :

```json
{
    "status": 404,
    "message": "User not found",
    "data": null
}
```

Kembali ke halaman [awal](../README.md)
