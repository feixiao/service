## 使用autotools生成Makefile

### 常用命令

#### autoreconf --install

根据Makefile.am和configure.ac生成configure等文件。 autoreconf -i 

#### configure

生成Makefile文件。

#### make all

生成全部产物。

#### make install / make install-strip

安装。

#### make clean 

删除make clean生成的产物。

#### make  uninstall

卸载。

#### make distclean

删除configure生成的产物。

#### make dist

打包程序。



### 参考资料

+ [《例解 autoconf 和 automake 生成 Makefile 文件》](https://www.ibm.com/developerworks/cn/linux/l-makefile/)
+ [《使用autotools生成Makefile学习笔记》](https://geesun.github.io/posts/2015/02/autotool.html)
+ [automake sample](https://github.com/raywill/automake)
+ [《GNU构建系统与Autotool概念分析》](https://www.linuxprobe.com/system-gnu-autotool.html)


