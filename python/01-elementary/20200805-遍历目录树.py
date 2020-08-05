import os

"""
/usr/local/etc/nginx
['conf.bak', 'conf.d']
['uwsgi_params.default', 'nginx.conf.default', 'mime.types.default', 'fastcgi_params', 'scgi_params', 'scgi_params.default', 'fastcgi.conf.default', 'fastcgi.conf', 'fastcgi_params.default', 'uwsgi_params', 'koi-win', 'mime.types', 'mime.types.default.default', 'koi-utf', 'win-utf', 'nginx.conf']

/usr/local/etc/nginx/conf.bak
[]
['mydoc.conf', 'package.test.conf', 'symfony3.conf', 'edusoho.conf', 'phalcon-admin-module.conf', 'phalcon-micro.conf', 'lumen-api.conf', 'thinkcms.conf', 'laravel-bbs.test.conf', 'payment.conf']

/usr/local/etc/nginx/conf.d
[]
['sample.com.conf', 'example.com.conf']
"""
path = '/usr/local/etc/nginx'
for dirname, subdirname, filename in os.walk(path):
    print(dirname)
    print(subdirname)
    print(filename)
    print()
