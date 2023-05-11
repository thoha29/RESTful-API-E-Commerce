# Login

Login ini digunakan untuk membuat token user

**URL** : `/login`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "email": "[valid email address]",
    "pass": "[password in plain text]"
}
```

**Data example**

```json
{
    "email": "ino@email.com",
    "password": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "token": "93144b288eb1fdccbe46d6fc0f241a51766ecd3d"
}
```

## Error Response

**Condition** : Jika email dan pass salah

**Code** : `500 Internal server error`

**Content** :

```json
{
    "message": "sql: no rows in result set"
}
```
