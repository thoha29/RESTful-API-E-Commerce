# Update Status Dari Order

Mengubah status order. Ini dilakukan oleh admin untuk mengubah status pesanan pelanggan

**URL** : `/orders/:id`

**Method** : `PUT`

**Auth required** : YES

**Data example**

```json
{
    "status": "berhasil"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Order status updated successfully"
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
{
    "message": "Status cannot be empty"
}
```
**Condition** : Jika order_id tidak ditemukan

**Code** : `500 Internal Server Error`

**Content** :

```json
"order not found"
```

Kembali ke halaman [awal](../README.md)
