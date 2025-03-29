# Test Auth API

🚀 **Test Auth API** is a simple authentication application using JWT (JSON Web Token) with **Golang** and **Echo**.

## 🛠 Teknologi Used

- [Golang](https://go.dev/)
- [Echo](https://echo.labstack.com/) (Framework HTTP)
- [JWT](https://jwt.io/) (JSON Web Token)
- [Validator](https://github.com/go-playground/validator) (Validasi Input)

---

## 📌 **Endpoint**
POST localhost:8085/login
raw data:
```
{
    "username": "test",
    "password": "password"
}
```
Note: The username and password can be any values, but both fields are required.

GET localhost:8085/profile
```
set header
Authorization : token
```
Note: To make sure the authorization can work to other service
---

## 📦 **How To Run**

### 1️⃣ Run with MakeFile
```
make tidy
make run
```
