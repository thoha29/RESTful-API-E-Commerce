# Create Order

Menambahkan produk dari keranjang ke order.

**URL** : `/orders`

**Method** : `POST`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Cart items successfully moved to order"
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

**Condition** : Jika tidak ada produk didalam cart

**Code** : `400 Bad Request`

**Content** :

```json
"no items found in cart"
```

Kembali ke halaman [awal](../README.md)
