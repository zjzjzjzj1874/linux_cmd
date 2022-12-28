- chmod
  - `change mode`:控制用户对文件的权限的命令
  - `only file owner and root can change mode`.只有文件所有者和超级用户可以修改文件权限
  - [mode结构](./../images/mode-struct.png)
  - [list -l in computer](./../images/mode.png)
  - [detail see at](https://www.runoob.com/linux/linux-comm-chmod.html)
  - syntax `chmod [-cfvR] [--help] [--version] mode file..`
  - example
    - 给所有用户读写执行权限:`chmod 777 a.txt` = `chmod a+rwx a.txt` = `chmod ugo+rwx` = `chmod a=rwx a.txt`
    - 删除文件所有用户的执行权限: `chmod a-x a.txt` = `chmod ugo-x a.txt`
    - 给文件拥有者和群组加rwx,其他用户执行权限:`chmod ug+rwx,o+x a.txt` = `chmod 771 a.txt`
    - 给文件拥有者所有权限(rwx),群组用户增加写权限,其他用户删除写权限: `chmod u=rwx,g+w,o-w a.txt`
    - 将当前目录下所有文件和文件夹设置为任何人都可以读: `chmod -R a+r *`
  - 总结:使用符号模式可以设置多个项目：who（用户类型），operator（操作符）和 permission（权限），
    每个项目的设置可以用逗号隔开。 命令 chmod 将修改 who 指定的用户类型对文件的访问权限，用户类型由一个或者多个字母在 who 的位置来说明
    - who的符号模式

      | Who |  用户类型  | 说明 |
      |:----:| :----: | :---- |
      | u |  user  | 文件拥有者 |
      | g | group  | 文件所有者所在组 |
      | o | others | 其他用户 |
      | a |  all   | 所有用户,=ugo |
    - operator的符号模式
    
      | Operator |    说明     |
      |:-----:| :----- |
      | + | 为指定用户增加权限 |
      | - | 去除指定用户权限  |
      | = | 设置指定用户权限  |
    - permission的符号模式

      | 模式   |  名字  |   说明    |
      |:----:|:-------:| :---- |
      | r   |  读  | 设置可读权限  |
      | w   | 写  | 设置可写权限  |
      | x   | 执行 | 设置可执行权限 |
      | X   |  特殊执行   |    -    |
  - 八进制语法
  
    | # | 权限 | rwx   | 二进制 |
    | :----: |:------| :----: | :----: |
    | 7 | 读+写+执行 | rwx   | 111 |
    | 6 | 读+写 | rw-   | 110 |
    | 5 | 读+执行 | r-x   | 101 |
    | 4 | 只读 | r--   | 100 |
    | 3 | 写+执行 | -wx   | 011 |
    | 2 | 只写 | -w-   | 010 |
    | 1 | 只执行 | --x   | 001 |
    | 0 | 无 | ---   | 000 |