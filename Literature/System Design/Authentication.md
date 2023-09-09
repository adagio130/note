
JWT(Json Web Token)
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
	2. Payload(**Base64** format, also called Access token):
		1. user defined
		2. Do not put secret information in a JWT payload unless it is encrypted
	3. Signature
		1. can be generated using both symmetric encryption algorithms and asymmetric ones.
		2. is used to verify the message wasn’t changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.
- Types:
	- Access token
		- is used to authorize requests and store additional information about the user
		- To make the application more secure, you need to store the access token in the memory of the client application, not in LocalStorage or any other place where a hacker can easily steal our token.**
	- Refresh token
		- issued by the server upon successful authentication and used to obtain a new pair of **access** and/or **refresh** tokens.
		- Since we are trying to protect our application as much as possible from hacking, we must store our refresh token exclusively in an HttpOnly Cookie.
		- Each token has its own lifetime, for example access: 30 min, refresh: 14 days. This means that after 30 minutes the access token will expire and the client will try to renew the token, but if the user does not interact with our application within 14 days, the user will be automatically logged out.