package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _container(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Container", name, desc, method)
}
