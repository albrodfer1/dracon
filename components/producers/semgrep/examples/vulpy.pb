
��Ŵsemgrep�
file://bad/api_list.py:10-10�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.Zpython.lang.security.audit.insecure-transport.requests.request-with-http.request-with-http 0:�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.
 extra lines:     r = requests.get('http://127.0.1.1:5000/api/post/{}'.format(username))Bunknownb �
file://bad/api_post.py:16-16�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.Zpython.lang.security.audit.insecure-transport.requests.request-with-http.request-with-http 0:�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.
 extra lines:         r = requests.post('http://127.0.1.1:5000/api/key', json={'username':username, 'password':password})Bunknownb �
file://bad/api_post.py:30-30�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.Zpython.lang.security.audit.insecure-transport.requests.request-with-http.request-with-http 0:�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.
 extra lines:     r = requests.post('http://127.0.1.1:5000/api/post', json={'text':message}, headers={'X-APIKEY': api_key})Bunknownb �
file://bad/db.py:19-19IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:         c.execute("INSERT INTO users (user, password, failures) VALUES ('%s', '%s', '%d')" %(u, p, 0))Bunknownb �
file://bad/db.py:19-19�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:         c.execute("INSERT INTO users (user, password, failures) VALUES ('%s', '%s', '%d')" %(u, p, 0))Bunknownb �
file://bad/db_init.py:20-20IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:         c.execute("INSERT INTO users (username, password, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%d', '%d', '%s')" %(u, p, 0, 0, ''))Bunknownb �
file://bad/db_init.py:20-20�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:         c.execute("INSERT INTO users (username, password, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%d', '%d', '%s')" %(u, p, 0, 0, ''))Bunknownb �
file://bad/libuser.py:12-12IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:     user = c.execute("SELECT * FROM users WHERE username = '{}' and password = '{}'".format(username, password)).fetchone()Bunknownb �
file://bad/libuser.py:12-12�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:     user = c.execute("SELECT * FROM users WHERE username = '{}' and password = '{}'".format(username, password)).fetchone()Bunknownb �
file://bad/libuser.py:25-25IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:     c.execute("INSERT INTO users (username, password, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%d', '%d', '%s')" %(username, password, 0, 0, ''))Bunknownb �
file://bad/libuser.py:25-25�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:     c.execute("INSERT INTO users (username, password, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%d', '%d', '%s')" %(username, password, 0, 0, ''))Bunknownb �
file://bad/libuser.py:53-53IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:     c.execute("UPDATE users SET password = '{}' WHERE username = '{}'".format(password, username))Bunknownb �
file://bad/libuser.py:53-53�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:     c.execute("UPDATE users SET password = '{}' WHERE username = '{}'".format(password, username))Bunknownb �

#file://bad/templates/csp.html:29-29�This tag is missing an 'integrity' subresource integrity attribute. The 'integrity' attribute allows for the browser to verify that externally hosted files (for example from a CDN) are delivered without unexpected manipulation. Without this attribute, if an attacker can modify the externally hosted resource, this could lead to XSS and other types of attacks. To prevent this, include the base64-encoded cryptographic hash of the resource (file) you’re telling the browser to fetch in the 'integrity' attribute for all externally hosted files.7html.security.audit.missing-integrity.missing-integrity 0:�This tag is missing an 'integrity' subresource integrity attribute. The 'integrity' attribute allows for the browser to verify that externally hosted files (for example from a CDN) are delivered without unexpected manipulation. Without this attribute, if an attacker can modify the externally hosted resource, this could lead to XSS and other types of attacks. To prevent this, include the base64-encoded cryptographic hash of the resource (file) you’re telling the browser to fetch in the 'integrity' attribute for all externally hosted files.
 extra lines:             <td><script src="https://apis.google.com/js/platform.js" async defer></script><g:plusone></g:plusone></td>Bunknownb �
)file://bad/templates/mfa.enable.html:8-14_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>OTP</b></label>
        <input class="w3-input w3-border" type="text" name="otp" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Enable</button>
      </div>
    </form>Bunknownb �
*file://bad/templates/posts.view.html:13-17_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:       <form method="POST" action="/posts/" class="w3-padding-16">
        <div class="w3-rest">
          <input class="w3-input w3-border" type="text" name="text" placeholder="What are you thinking?">
        </div>
      </form>Bunknownb �
,file://bad/templates/user.chpasswd.html:8-16_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>New Password</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="password" name="password" required>
        <label><b>Again</b></label>
        <input class="w3-input w3-border" type="password" name="password_again" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Change</button>
      </div>
    </form>Bunknownb �
*file://bad/templates/user.create.html:8-18_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Username</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Email</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="email" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border" type="password" name="password" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Create</button>
      </div>
    </form>Bunknownb �
)file://bad/templates/user.login.html:8-16_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Email</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border" type="password" name="password" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Login</button>
      </div>
    </form>Bunknownb �
-file://bad/templates/user.login.mfa.html:8-18_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Username</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="password" name="password" required>
        <label><b>OTP (if enabled in your account)</b></label>
        <input class="w3-input w3-border" type="text" name="otp">
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Login</button>
      </div>
    </form>Bunknownb �
&file://bad/templates/welcome.html:7-11_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" action="/post" class="w3-padding-16">
      <div class="w3-rest">
        <input class="w3-input w3-border" type="text" name="text" placeholder="What are you thinking?">
      </div>
    </form>Bunknownb �
file://bad/vulpy-ssl.py:13-13[Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files insteadNpython.flask.security.audit.hardcoded-config.avoid_hardcoded_config_SECRET_KEY 0:�Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files instead
 extra lines: app.config['SECRET_KEY'] = 'aaaaaaa'Bunknownb �
file://bad/vulpy-ssl.py:29-29ptop-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a functionPpython.flask.security.audit.app-run-security-config.avoid_using_app_run_directly 0:�top-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a function
 extra lines: app.run(debug=True, host='127.0.1.1', ssl_context=('/tmp/acme.cert', '/tmp/acme.key'))Bunknownb �
file://bad/vulpy-ssl.py:29-29�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.7python.flask.security.audit.debug-enabled.debug-enabled 0:�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.
 extra lines: app.run(debug=True, host='127.0.1.1', ssl_context=('/tmp/acme.cert', '/tmp/acme.key'))Bunknownb �
file://bad/vulpy.py:16-16[Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files insteadNpython.flask.security.audit.hardcoded-config.avoid_hardcoded_config_SECRET_KEY 0:�Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files instead
 extra lines: app.config['SECRET_KEY'] = 'aaaaaaa'Bunknownb �
file://bad/vulpy.py:55-55ptop-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a functionPpython.flask.security.audit.app-run-security-config.avoid_using_app_run_directly 0:�top-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a function
 extra lines: app.run(debug=True, host='127.0.1.1', port=5000, extra_files='csp.txt')Bunknownb �
file://bad/vulpy.py:55-55�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.7python.flask.security.audit.debug-enabled.debug-enabled 0:�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.
 extra lines: app.run(debug=True, host='127.0.1.1', port=5000, extra_files='csp.txt')Bunknownb �
file://good/httpbrute.py:22-22�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.Zpython.lang.security.audit.insecure-transport.requests.request-with-http.request-with-http 0:�Detected a request using 'http://'. This request will be unencrypted, and attackers could listen into traffic on the network and be able to obtain sensitive information. Use 'https://' instead.
 extra lines:     response = requests.post(URL, data = {'username': username, 'password': password})Bunknownb �
file://good/libapi.py:24-24�Hardcoded JWT secret or private key is used. This is a Insufficiently Protected Credentials weakness: https://cwe.mitre.org/data/definitions/522.html Consider using an appropriate security mechanism to protect the credentials (e.g. keeping secrets in environment variables)<python.jwt.security.jwt-hardcode.jwt-python-hardcoded-secret 0:�Hardcoded JWT secret or private key is used. This is a Insufficiently Protected Credentials weakness: https://cwe.mitre.org/data/definitions/522.html Consider using an appropriate security mechanism to protect the credentials (e.g. keeping secrets in environment variables)
 extra lines:         }, secret, algorithm='HS256').decode()Bunknownb �
file://good/libuser.py:61-61IDetected possible formatted SQL query. Use parameterized queries instead.Bpython.lang.security.audit.formatted-sql-query.formatted-sql-query 0:�Detected possible formatted SQL query. Use parameterized queries instead.
 extra lines:     c.execute("INSERT INTO users (username, password, salt, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%s', '%d', '%d', '%s')" %(username, '', '', 0, 0, ''))Bunknownb �	
file://good/libuser.py:61-61�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.Tpython.sqlalchemy.security.sqlalchemy-execute-raw-query.sqlalchemy-execute-raw-query 0:�Avoiding SQL string concatenation: untrusted input concatenated with raw SQL query can result in SQL Injection. In order to execute raw query safely, prepared statement should be used. SQLAlchemy provides TextualSQL to easily used prepared statement with named parameters. For complex SQL composition, use SQL Expression Language or Schema Definition Language. In most cases, SQLAlchemy ORM will be a better option.
 extra lines:     c.execute("INSERT INTO users (username, password, salt, failures, mfa_enabled, mfa_secret) VALUES ('%s', '%s', '%s', '%d', '%d', '%s')" %(username, '', '', 0, 0, ''))Bunknownb �

$file://good/templates/csp.html:29-29�This tag is missing an 'integrity' subresource integrity attribute. The 'integrity' attribute allows for the browser to verify that externally hosted files (for example from a CDN) are delivered without unexpected manipulation. Without this attribute, if an attacker can modify the externally hosted resource, this could lead to XSS and other types of attacks. To prevent this, include the base64-encoded cryptographic hash of the resource (file) you’re telling the browser to fetch in the 'integrity' attribute for all externally hosted files.7html.security.audit.missing-integrity.missing-integrity 0:�This tag is missing an 'integrity' subresource integrity attribute. The 'integrity' attribute allows for the browser to verify that externally hosted files (for example from a CDN) are delivered without unexpected manipulation. Without this attribute, if an attacker can modify the externally hosted resource, this could lead to XSS and other types of attacks. To prevent this, include the base64-encoded cryptographic hash of the resource (file) you’re telling the browser to fetch in the 'integrity' attribute for all externally hosted files.
 extra lines:             <td><script src="https://apis.google.com/js/platform.js" async defer></script><g:plusone></g:plusone></td>Bunknownb �
*file://good/templates/mfa.enable.html:8-14_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>OTP</b></label>
        <input class="w3-input w3-border" type="text" name="otp" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Enable</button>
      </div>
    </form>Bunknownb �
+file://good/templates/posts.view.html:13-17_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:       <form method="POST" action="/posts/" class="w3-padding-16">
        <div class="w3-rest">
          <input class="w3-input w3-border" type="text" name="text" placeholder="What are you thinking?">
        </div>
      </form>Bunknownb �
-file://good/templates/user.chpasswd.html:8-18_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Current Password</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="password" name="current_password" required>
        <label><b>New Password</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="password" name="new_password" required>
        <label><b>New Password (again)</b></label>
        <input class="w3-input w3-border" type="password" name="new_password_again" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Change</button>
      </div>
    </form>Bunknownb �
+file://good/templates/user.create.html:8-18_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Username</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Email</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="email" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border" type="password" name="password" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Create</button>
      </div>
    </form>Bunknownb �
*file://good/templates/user.login.html:8-16_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Email</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border" type="password" name="password" required>
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Login</button>
      </div>
    </form>Bunknownb �
.file://good/templates/user.login.mfa.html:8-18_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" class="w3-padding-32">
      <div class="w3-section">
        <label><b>Username</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="text" name="username" required>
        <label><b>Password</b></label>
        <input class="w3-input w3-border w3-margin-bottom" type="password" name="password" required>
        <label><b>OTP (if enabled in your account)</b></label>
        <input class="w3-input w3-border" type="text" name="otp">
        <button class="w3-button w3-block w3-black w3-section w3-padding" type="submit">Login</button>
      </div>
    </form>Bunknownb �
'file://good/templates/welcome.html:7-11_Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.@python.django.security.django-no-csrf-token.django-no-csrf-token 0:�Manually-created forms in django templates should specify a csrf_token to prevent CSRF attacks.
 extra lines:     <form method="POST" action="/post" class="w3-padding-16">
      <div class="w3-rest">
        <input class="w3-input w3-border" type="text" name="text" placeholder="What are you thinking?">
      </div>
    </form>Bunknownb �
file://good/vulpy-ssl.py:13-13[Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files insteadNpython.flask.security.audit.hardcoded-config.avoid_hardcoded_config_SECRET_KEY 0:�Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files instead
 extra lines: app.config['SECRET_KEY'] = 'aaaaaaa'Bunknownb �
file://good/vulpy-ssl.py:29-29ptop-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a functionPpython.flask.security.audit.app-run-security-config.avoid_using_app_run_directly 0:�top-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a function
 extra lines: app.run(debug=True, host='127.0.1.1', ssl_context=('/tmp/acme.cert', '/tmp/acme.key'))Bunknownb �
file://good/vulpy-ssl.py:29-29�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.7python.flask.security.audit.debug-enabled.debug-enabled 0:�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.
 extra lines: app.run(debug=True, host='127.0.1.1', ssl_context=('/tmp/acme.cert', '/tmp/acme.key'))Bunknownb �
file://good/vulpy.py:17-17[Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files insteadNpython.flask.security.audit.hardcoded-config.avoid_hardcoded_config_SECRET_KEY 0:�Hardcoded variable `SECRET_KEY` detected. Use environment variables or config files instead
 extra lines: app.config['SECRET_KEY'] = '123aa8a93bdde342c871564a62282af857bda14b3359fde95d0c5e4b321610c1'Bunknownb �
file://good/vulpy.py:53-53ptop-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a functionPpython.flask.security.audit.app-run-security-config.avoid_using_app_run_directly 0:�top-level app.run(...) is ignored by flask. Consider putting app.run(...) behind a guard, like inside a function
 extra lines: app.run(debug=True, host='127.0.1.1', port=5001, extra_files='csp.txt')Bunknownb �
file://good/vulpy.py:53-53�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.7python.flask.security.audit.debug-enabled.debug-enabled 0:�Detected Flask app with debug=True. Do not deploy to production with this flag enabled as it will leak sensitive information. Instead, consider using Flask configuration variables or setting 'debug' using system environment variables.
 extra lines: app.run(debug=True, host='127.0.1.1', port=5001, extra_files='csp.txt')Bunknownb 