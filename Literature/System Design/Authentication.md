
JWT Token
[RFC7519](https://datatracker.ietf.org/doc/html/rfc7519)

### mainly used:
for authorization on server.

example:`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

decoded:
``{ alg: "HS256", typ: "JWT" }.{ iss: "my-site.auth.com", aud: "my-site.com", exp: 1435937883, id: 1234567890, username: "Denis Dekh", roles: ["Admin"] }.S9Zs/8/uEGGTVVtLggFTizCsMtwOJnRhjaQ2BMUQhcY`

- JSON Web Signature (JWS) 
- JSON Web Encryption (JWE) 
- Message Authentication Code (MAC) 
- Structure:
	1. Header(**Base64** format)
		1. the type of the token, which is JWT, 
		2. signing algorithm being used.
	2. Payload(**Base64** format)
	3. Signature
- 
