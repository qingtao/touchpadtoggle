#!/bin/bash

CMD="$HOME/go/bin/touchpad_toggle"
TITLE="触摸板"
ENABLED="已启用"
DISABLED="已禁用"

# toggle 切换触摸板的启用或者启用
toggle() {
    local result=`$CMD -debug 2>&1`
    op=`echo $result| grep -o -E "set state to (enable|disable) success"| awk '{ print $4 }'`
    if [ "$op" = "enable" ]; then
        echo $ENABLED
    else
        echo $DISABLED
    fi
}

state=`toggle`
# 桌面通知
notify-send -t 50 "$TITLE" "$state"

