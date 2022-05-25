# images_upload_tool
github 图床上传工具  【喜欢的欢迎 点击 star  】

##  typora 代码配置

![image](https://user-images.githubusercontent.com/46613910/127744213-a712db07-42d1-4723-a74f-a61515461b66.png)


![image](https://user-images.githubusercontent.com/46613910/127744226-ec110750-cac7-4b6e-8003-5e0a416c0968.png)

打包后 的 exe ,的文件路径 复制到typora 设置 那里就可以了

上传图片后的效果：
https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/2021_07_31_22__59_01a0.png
![](https://cdn.jsdelivr.net/gh/lyr-2000/images_repo_2021_ASUS/2021_07_31_22__59_01a0.png)


 


 
 ## python 脚本


```bash
alias upimg='pushd_c; cd $script_home;python upload_img.py -format="md" -dir="/f/git_image_blog/staticFS/IMAGES/2021_7_31"; popd_c;  '

upimg
```


