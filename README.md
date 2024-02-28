# OTP API

API to generate TOTP from a secret key.

## API Documentation

### Generate TOTP

```http
GET /totp?secret=${secret}
```

- Sample Request:

    ```bash
    curl -X GET "http://otp.bitgarden.tech/totp?secret=JBSWY3DPEHPK3PXP"
    ```

- Sample Response:

    ```txt
    123456
    ```