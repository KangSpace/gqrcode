import qrcode
import zxing


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
    img = input("图片路径：")
    reader = zxing.BarCodeReader()
    barcode = reader.decode(img)
    return barcode.parsed



# 安装依赖库
# pip install qrcode pillow image zxing