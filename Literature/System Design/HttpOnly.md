[RFC6265](https://datatracker.ietf.org/doc/html/rfc6265)
當 cookie 有設定 HttpOnly flag 時，瀏覽器會限制 cookie 只能經由 HTTP(S) 協定來存取。因此當網站有 XSS 弱點時，若 cookie 含有 HttpOnly flag，則攻擊者無法直接經由 JavaScript 存取使用者的 session cookie
因此 HttpOnly 可有效降低 XSS 的影響並提升攻擊難度。