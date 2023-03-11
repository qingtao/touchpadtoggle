#!/bin/bash
# author: wuqingtao
# description: 切换笔记本触摸板的状态, 适用于Xwayland

TITLE="触摸板"
ENABLED="已启用"
DISABLED="已禁用"
# 过期时间毫秒
EXPIRES=20

# 通过 gsettings 读取当前的touchpad状态
getState() {
    echo -n `gsettings get org.gnome.desktop.peripherals.touchpad send-events | grep -o -E "(enabled|disabled)"`
}

# 设置touchpad的状态
setState() {
    local state=$1
    echo $state
    if [ "$state" = "enabled" -o "$state" = "disabled" ]; then
        gsettings set org.gnome.desktop.peripherals.touchpad send-events $state
    fi
}

# toggle 切换触摸板的启用或者启用
toggle() {
    # 读取触摸板是启用还是禁用
    local state=`getState`
    local action=""
    local msg=""

    # find action
    case $state in
    "enabled")
        action="disabled"
        msg="$DISABLED"
        ;;
    "disabled")
        action="enabled"
        msg="$ENABLED"
        ;;
    *)
    ;;
    esac

    # 若action非空才执行
    if [ "a$action" != "a" ]; then
        setState $action
    fi
    # 再次获取状态, 如果与预期一致发送桌面通知
    state=`getState`
    if [ "$state" = "$action" ]; then
        notify-send -c device -t $EXPIRES "$TITLE" "$msg"
    fi
}

# 执行切换
toggle
