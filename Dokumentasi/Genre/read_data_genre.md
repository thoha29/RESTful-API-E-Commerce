# Get All Genre

Melihat data genre secara keseluruhan

**URL** : `/genres`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content**

```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "genre_id": 1,
            "genre_name": "Action"
        },
        {
            "genre_id": 2,
            "genre_name": "Advanture"
        },
        {
            "genre_id": 3,
            "genre_name": "Yuri"
        },
        {
            "genre_id": 5,
            "genre_name": "Hentai"
        },
        {
            "genre_id": 6,
            "genre_name": "Game"
        },
        {
            "genre_id": 7,
            "genre_name": "Fantasy"
        }
    ]
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

Kembali ke halaman [awal](../README.md)
