# touchpad_toggle 实现命令行开关触摸板

## 1. 安装

```sh
go install github.com/qingtao/touchpadtoggle/cmd/touchpad_toggle@latest
```

## 2. 使用方法

```
$ touchpad_toggle -h
Usage of touchpad_toggle:
  -debug
        启用debug
  -id string
        指定touchpad的id而不是使用自动查询
```

touchpad的id可以通过命令手工查询:

```sh
$ xinput list
```

## 3. 设置快捷键

1. 桌面环境: "ubuntu 20.04"
2. touchpad_toggle.sh目录: "/home/tom/bin/touchpadtoggle.sh"
3. 找到: "设置" -> "键盘快捷键" -> "自定义快捷键" -> 点击"+", 按照需要设置快捷键:

   - 名称: "切换触摸板"
   - 命令: "/home/tom/bin/touchpadtoggle.sh"
   - 快捷键: Super + F9


    设置完成后, 按 "Super + F9" 可以开启或者关闭触摸板.


