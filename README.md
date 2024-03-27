# SayaKaya Technical Test - Backend Engineer
Program Language: Go (Echo Framework) \
Database: PostgreSQL

## Flowchart Diagram
![sayakaya test drawio](https://github.com/kevinlau2102/sayakaya-test/assets/71743231/b09c0f8a-3d8f-44c0-9ce2-0e5939ddba21)

## Database Design
![sayakaya db design](https://github.com/kevinlau2102/sayakaya-test/assets/71743231/327e9c6e-d3df-4afd-92c5-7d5b6d325209)

### Happy Birthday Email (Using SMTP)

![sayakaya email smtp](https://github.com/kevinlau2102/sayakaya-test/assets/71743231/f67592a2-65f9-4959-9281-db3c312bba34)

## API Documentation

### 1. Fetch Verified Users
```http
GET /api/v1/users
```
#### Response
```json
{
    "code": "200",
    "message": "success",
    "data": [
        {
            <user_data>
        },
        {
            <user_data>
        }
    ]
}
```

### 2. Check Promo Code
```http
GET /api/v1/promo?code=<code>&user_id=<user_id>
```

#### Request
| Parameter | Type | Description |
| :--- | :--- | :--- |
| `code` | `string` | **Required**. Promo Code |
| `user_id` | `int` | **Required**. User ID |

#### Response
```json
{
    "code": "200",
    "message": "success",
    "data": {
        <promo_data>
    }
}
```
