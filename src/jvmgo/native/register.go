package native

import (
    "fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
    _ "jvmgo/native/java/lang"
    _ "jvmgo/native/java/security"
)

// register native methods
func init() {
    rtc.SetRegisterNatives(registerNatives)
    // hack!
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}

func registerNatives(operandStack *rtda.OperandStack) {
    // todo
}

// hack
func jvmgo_SystemOut_println(stack *rtda.OperandStack) {
    str := stack.PopRef()
    this := stack.PopRef()
    this.Class()
    chars := str.Class().GetField("value", "[C").GetValue(str).(*rtc.Obj).Fields().([]uint16)
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}