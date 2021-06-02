#-*- coding: UTF-8 -*-
import sys

import pyzbar.pyzbar
import qrcode
import zxing
from PIL import Image

# 生成7089行的txt文件
def make_txt():
    with open("nmb.txt", "w") as f:
        list1 = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
        for x in range(7089):
            if x == 0:
                f.write(str(x) + '\n')
            else:
                str1 = ""
                for y in range(x + 1):
                    if y < 10:
                        str1 += str(list1[y])
                    else:
                        str1 += str(list1[y % 10])
                f.write(str1 + '\n')


# 生成二维码
def make_qrcode(data, filename):
    # 二维码内容
    # data = "123456789"
    # 生成二维码
    img = qrcode.make(data=data)
    # 直接显示二维码
    img.show()
    # 保存二维码为文件
    img.save(filename)
    return filename


# 解码
def get_data(img):
    reader = zxing.BarCodeReader()
    barcode = reader.decode(img)
    return barcode.parsed


def get_data_pyzbar(img):
    return str(pyzbar.pyzbar.decode(Image.open(img), symbols=[pyzbar.pyzbar.ZBarSymbol.QRCODE])[0].data, 'utf-8')

if __name__ == '__main__':
    # the first arg is the file name for decode.
    file_name = sys.argv[1]
    print(get_data_pyzbar(file_name))

# 安装依赖库
# pip install qrcode pillow image zxing pyzbar
# mac: brew install zbar
# linux: yum install zbar-devel
# ubuntu: sudo apt-get install libzbar-dev
