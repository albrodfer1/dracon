
ιôgosec�
6file:///code/pkg/sessions/redis/redis_store.go:191-191G402TLS MinVersion too low. 0:D190: 			/* #nosec */
191: 			opt.TLSConfig = &tls.Config{}
192: 		}
Bunknownb j��
6file:///code/pkg/sessions/redis/redis_store.go:165-165G402TLS MinVersion too low. 0:D164: 			/* #nosec */
165: 			opt.TLSConfig = &tls.Config{}
166: 		}
Bunknownb j��
,file:///code/pkg/validation/options.go:36-36G402 TLS InsecureSkipVerify set true. 0:s35: 		insecureTransport := &http.Transport{
36: 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
37: 		}
Bunknownb j��
2file:///code/pkg/middleware/jwt_session.go:119-119G101Potential hardcoded credentials 0:l118: 		/* #nosec G101 */
119: 		if password == "" || password == "x-oauth-basic" {
120: 			return user, nil
Bunknownb j��
9file:///code/pkg/authentication/basic/htpasswd.go:163-163G401#Use of weak cryptographic primitive 0:|162: 		// We support SHA1 HTPasswd entries
163: 		d := sha1.New() // #nosec G401
164: 		_, err := d.Write([]byte(password))
Bunknownb j��
5file:///code/pkg/app/pagewriter/sign_in_page.go:79-79G203�The used method does not auto-escape HTML. This can potentially lead to 'Cross-site Scripting' vulnerabilities, in case the attacker controls the input. 0:e78: 		Footer:        template.HTML(s.footer),
79: 		LogoData:      template.HTML(s.logoData),
80: 	}
Bunknownb jO�
5file:///code/pkg/app/pagewriter/sign_in_page.go:78-78G203�The used method does not auto-escape HTML. This can potentially lead to 'Cross-site Scripting' vulnerabilities, in case the attacker controls the input. 0:�77: 		ProxyPrefix:   s.proxyPrefix,
78: 		Footer:        template.HTML(s.footer),
79: 		LogoData:      template.HTML(s.logoData),
Bunknownb jO�
5file:///code/pkg/app/pagewriter/sign_in_page.go:72-72G203�The used method does not auto-escape HTML. This can potentially lead to 'Cross-site Scripting' vulnerabilities, in case the attacker controls the input. 0:{71: 		ProviderName:  s.providerName,
72: 		SignInMessage: template.HTML(s.signInMessage),
73: 		StatusCode:    statusCode,
Bunknownb jO�
3file:///code/pkg/app/pagewriter/error_page.go:79-79G203�The used method does not auto-escape HTML. This can potentially lead to 'Cross-site Scripting' vulnerabilities, in case the attacker controls the input. 0:m78: 		RequestID:   opts.RequestID,
79: 		Footer:      template.HTML(e.footer),
80: 		Version:     e.version,
Bunknownb jO�
#file:///code/pkg/util/util.go:24-24G304%Potential file inclusion via variable 0:x23: 		// Cert paths are a configurable option
24: 		data, err := os.ReadFile(path) // #nosec G304
25: 		if err != nil {
Bunknownb j�
7file:///code/pkg/authentication/basic/htpasswd.go:57-57G304%Potential file inclusion via variable 0:{56: 	// We allow HTPasswd location via config options
57: 	r, err := os.Open(filename) // #nosec G304
58: 	if err != nil {
Bunknownb j�
5file:///code/pkg/app/pagewriter/static_pages.go:95-95G304%Potential file inclusion via variable 0:l94: 	if p.dir != "" && isFile(filePath) {
95: 		content, err := os.ReadFile(filePath)
96: 		if err != nil {
Bunknownb j�
7file:///code/pkg/app/pagewriter/sign_in_page.go:115-115G304%Potential file inclusion via variable 0:I114: 
115: 	logoData, err := os.ReadFile(logoPath)
116: 	if err != nil {
Bunknownb j�
-file:///code/pkg/apis/options/load.go:152-152G304%Potential file inclusion via variable 0:K151: 
152: 	data, err := os.ReadFile(configFileName)
153: 	if err != nil {
Bunknownb j�
5file:///code/pkg/authentication/basic/htpasswd.go:5-5G505<Blocklisted import crypto/sha1: weak cryptographic primitive 0:c4: 	// We support SHA1 & bcrypt in HTPasswd
5: 	"crypto/sha1" // #nosec G505
6: 	"encoding/base64"
Bunknownb j��
file:///code/validator.go:73-73G103%Use of unsafe calls should be audited 0:V72: 	}
73: 	atomic.StorePointer(&um.m, unsafe.Pointer(&updated)) // #nosec G103
74: }
Bunknownb j��
file:///code/validator.go:27-27G103%Use of unsafe calls should be audited 0:}26: 	m := make(map[string]bool)
27: 	atomic.StorePointer(&um.m, unsafe.Pointer(&m)) // #nosec G103
28: 	if usersFile != "" {
Bunknownb j��
$file:///code/oauthproxy.go:1232-1232G104Errors unhandled. 0:A1231: 	// application/json
1232: 	rw.Write([]byte("{}"))
1233: }
Bunknownb j�