import zipfile
import glob, os

"""
压缩
"""
path = '/usr/local/etc/nginx/*'
zfile = zipfile.ZipFile('/tmp/test.zip', 'w')
for name in glob.glob(path):
    zfile.write(name, os.path.basename(name), zipfile.ZIP_DEFLATED)
zfile.close()


"""
查看
"""
path = '/tmp/test.zip'
listzip = zipfile.ZipFile(path, 'r')
print(listzip.namelist())
for info in listzip.infolist():
    print((info.filename, info.file_size, info.compress_size))


"""
解压
"""
path = '/tmp/test.zip'
unzip = zipfile.ZipFile(path)
unzip.extractall('/tmp')
unzip.close()
