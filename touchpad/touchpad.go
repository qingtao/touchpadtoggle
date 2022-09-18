// Package touchpad 基于 xinput设置触摸板的启用/禁用
package touchpad

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sync/atomic"
)

var (
	// xinput 命令
	_xinput = "xinput"
	// _touchpadIdRegexp 搜索触摸板id
	_touchpadIdRegexp = regexp.MustCompile(`.+TOUCHPAD\s+ID=(\w+)\s+.*`)
	// _touchpadStateRegexp 搜索触摸板状态
	_touchpadStateRegexp = regexp.MustCompile(`.+DEVICE\s*ENABLED.+:\s(\d+).*`)
	// 调试开关
	debug = new(atomic.Bool)
)

// SetDebug 设置调试开关
func SetDebug(b bool) {
	debug.Store(b)
}

// FindTouchpad 找到触摸板
//
//	xinput list
func FindTouchpad(ctx context.Context) (id string, err error) {
	c := exec.CommandContext(ctx, _xinput, "list")
	b, err := c.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("get id of touchpad failed: %w, %s", err, b)
		return
	}
	br := bufio.NewScanner(bytes.NewReader(bytes.ToUpper(b)))
	for br.Scan() {
		s := br.Text()
		ss := _touchpadIdRegexp.FindStringSubmatch(s)
		if len(ss) == 2 {
			id = ss[1]
			break
		}
	}
	if id == "" {
		err = errors.New("touchpad not found")
	}
	return
}

// GetStateOfTouchpad 获取触摸板的当前状态
//
//	xinput list-props 15
func GetStateOfTouchpad(ctx context.Context, id string) (enabled bool, err error) {
	c := exec.CommandContext(ctx, _xinput, "list-props", id)
	b, err := c.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("get state of touchpad %s failed: %w, %s", id, err, b)
		return
	}
	br := bufio.NewScanner(bytes.NewReader(bytes.ToUpper(b)))
	for br.Scan() {
		s := br.Text()
		ss := _touchpadStateRegexp.FindStringSubmatch(s)
		if len(ss) == 2 {
			if ss[1] == "1" {
				enabled = true
			}
			break
		}
	}
	return
}

// touchpadEnable 启用触摸板
//
//	xinput enable 15
func touchpadEnable(ctx context.Context, id string) error {
	c := exec.CommandContext(ctx, _xinput, "enable", id)
	b, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("enable touchpad failed: %w %s", err, b)
	}
	return nil
}

// touchpadDisable 禁用触摸板
//
//	xinput disable 15
func touchpadDisable(ctx context.Context, id string) error {
	c := exec.CommandContext(ctx, _xinput, "disable", id)
	b, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("enable touchpad failed: %w %s", err, b)
	}
	return nil
}

func getToggleAction(enabled bool) string {
	if enabled {
		return "disable"
	}
	return "enable"
}

// ToggleStateOfTouchpad 切换触摸板的状态
// 如果触摸版已经启用, 则设置其为禁用, 否则启用它
func ToggleStateOfTouchpad(ctx context.Context, id string) error {
	enabled, err := GetStateOfTouchpad(ctx, id)
	if err != nil {
		return err
	}
	if debug.Load() {
		log.Printf("[DEBUG] touchpad id = %s, enabled is %t", id, enabled)
	}
	if enabled {
		err = touchpadDisable(ctx, id)
	} else {
		err = touchpadEnable(ctx, id)
	}
	if debug.Load() {
		action := getToggleAction(enabled)
		if err != nil {
			log.Printf("[DEBUG] touchpad id = %s, set state to %s failed: %s", id, action, err)
		} else {
			log.Printf("[DEBUG] touchpad id = %s, set state to %s success", id, action)
		}
	}
	return err
}
