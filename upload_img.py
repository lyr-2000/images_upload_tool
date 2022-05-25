
import argparse
from datetime import date
from operator import truediv
import os

# cwd = os.getcwd()

# print (cwd)
# print (os.path.dirname(os.path.realpath(__file__)))
# sdir = (os.path.dirname(os.path.realpath(__file__)))
def path_exists(dir):
    return os.path.exists(dir)
def check_path(dir):
    if not os.path.exists(dir):
        os.makedirs(dir)
def copy_file(src,dst):
    import shutil
    shutil.copy2(src,dst)

def copy_clipboard(s):
    import pyperclip
    pyperclip.copy(s)
    # paste 


def upload():

    from PIL import ImageGrab, Image
    im = ImageGrab.grabclipboard()
    if isinstance(im, Image.Image):
        check_path("./tmp/")
        im.save('./tmp/tmp.jpg')
        print("save ok")
    else:
        print("not an image")
        return 

    # upload , parse args ，解析命令行参数
    parser = argparse.ArgumentParser("For test the parser")
    parser.add_argument('-dir', '--dir', default="", help='upload img dir')
    parser.add_argument("-url","--url",default="https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/",help="custom urlfile")
    parser.add_argument("-format","--format",default="raw",help="format")
    args = parser.parse_args()

    # git status
    if args.dir == "":
        dir = "./temp"
    else:
        dir = args.dir
    # print(args.dir)
 
    from datetime import datetime
    now = datetime.now() # current date and time
    date_time = now.strftime("%Y_%m_%d_%H_%M_%S")
    disk_file_name = dir+"/"+date_time+".jpg"
    urlp =  args.url +date_time+".jpg"
    # print(disk_file_name)
    # return 
    copy_file("./tmp/tmp.jpg",disk_file_name)
    if not path_exists(dir):
        print("path not exists {}".format(dir))
        return 
    os.chdir(dir)

    x = True
    # /f/git_image_blog/region0/staticFS/IMAGES/2021_7_31
    if x:
        try:
            i = os.system("git status")
            if i is not 0:
                raise Exception("error occur on git status")
            os.system("git add .")
            i = os.system("git commit -m 'addimg'")
            print(urlp)
            if args.format=='md' or args.format=="markdown":
                copy_clipboard("![]({})".format(urlp))
            else:
                copy_clipboard(urlp)
            


            if i is not 0:
                raise Exception("commit error")
            os.system("git push")

        except Exception as e:
            print("error occur {}".format(e))
        # os.system("git commit -m add image")
        # os.system("git pull")
        # os.system("git push")

upload()

#  print("date and time:",date_time)
