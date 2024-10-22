# Cubs Coin API Documentation

---

## Public Routes

### 1. Get All Transactions
**GET** `/public/transactions`  
**Description:** Fetches all transactions.  
**Response:** JSON array of transactions.

---

### 2. Get Transaction by ID
**GET** `/public/transaction/{id}`  
**Description:** Fetches a transaction by its ID.  
**Parameters:**  
- `id`: Transaction ID (integer)  
**Response:** Transaction details.

---

### 3. Sign Up
**POST** `/public/signup`  
**Description:** Registers a new user.  
**Request Body:**
```json
{
  "student_id": "string",
  "name": "string",
  "password": "string"
}
```
**Response:** User creation confirmation or error message.
**Note:** default Balance=0, Role=member

---

### 4. Sign In
**POST** `/public/signin`  
**Description:** Logs in a user and returns a JWT token.  
**Request Body:**
```json
{
  "student_id": "string",
  "password": "string"
}
```
**Response:** JWT token for authenticated access.

---

## Private Routes

### 1. Get My Info
**GET** `/private/user`  
**Description:** Fetches information about the authenticated user.  
**Authentication:** Requires Bearer token.  
**Response:** User details.

---

### 2. Transfer Cubs Coins
**POST** `/public/transaction/create`  
**Description:** Transfers Cubs Coins to another user.  
**Request Body:**
```json
{
  "receiver": "string",
  "amount": number
}
```
**Response:** Transfer confirmation or error message.

---

## Admin Routes

### 1. Create Transaction
**POST** `/public/transaction/create`  
**Description:** Creates a new transaction between two users.  
**Request Body:**
```json
{
  "sender": "string",
  "receiver": "string",
  "amount": number
}
```
**Response:** Transaction creation confirmation or error message.

---

### 2. Get All Users
**GET** `/public/users`  
**Description:** Fetches a list of all users.  
**Response:** JSON array of user details.

---

### 3. Get User by ID
**GET** `/api/user/{id}`  
**Description:** Fetches a user's details by their ID.  
**Parameters:**  
- `id`: User ID (integer)  
**Response:** User details.

---

### 4. Create User
**POST** `/public/user/create`  
**Description:** Creates a new user.  
**Request Body:**
```json
{
  "student_id": "string",
  "name": "string",
  "password": "string"
}
```
**Response:** User creation confirmation or error message.

---

### 5. Change Role to Admin
**PATCH** `/api/changetoadmin/{id}`  
**Description:** Changes a user's role to "admin".  
**Authentication:** Requires Bearer token (admin privileges).  
**Parameters:**  
- `id`: User ID (integer)  
**Response:** Role change confirmation or error message.

---

### 6. Change Role to Member
**PATCH** `/api/changetomember/{id}`  
**Description:** Changes a user's role to "member".  
**Authentication:** Requires Bearer token (admin privileges).  
**Parameters:**  
- `id`: User ID (integer)  
**Response:** Role change confirmation or error message.

---

## Contributors
- Chayoot Kositwanich