package rtda

import (
    "testing"
    rtc "jvmgo/jvm/rtda/class"
    . "jvmgo/test"
)

func TestPushAndPop(t *testing.T) {
    stack := newOperandStack(6)
    stack.PushNull()
    stack.PushRef(&rtc.Obj{})
    stack.PushInt(-37)
    stack.PushLong(0xabcd1234ff)
    stack.PushFloat(3.14)
    stack.PushDouble(-2.71828)
    //stack.PushInt(0)

    AssertEquals(-2.71828, stack.PopDouble())
    AssertEquals(3.14, stack.PopFloat())
    AssertEquals(int64(0xabcd1234ff), stack.PopLong())
    AssertEquals(-37, stack.PopInt())
    AssertNotNil(stack.PopRef())
    // if x := stack.PopRef(); x != nil {
    //     t.Errorf("not nil: %v", x)
    // }
    AssertNil(stack.PopRef())
}

func TestPopN(t *testing.T) {
    stack := newOperandStack(6)
    stack.PushInt(4)
    stack.PushInt(5)
    stack.PushInt(6)
    stack.PushInt(8)
    stack.PushInt(9)

    top3 := stack.popN(3)
    if top3[0].(int32) != 6 || top3[1].(int32) != 8 || top3[2].(int32) != 9 {
        t.Errorf("top3: %v", top3)
    }

    if newTop := stack.PopInt(); newTop != 5 {
        t.Errorf("newTop: %v", newTop)
    }
}

func TestTop(t *testing.T) {
    stack := newOperandStack(3)
    stack.PushInt(1)
    stack.PushInt(2)
    stack.PushInt(3)
    
    if top0 := stack.Top(0).(int32); top0 != 3 {
        t.Errorf("top0: %v", top0)
    }
    if top1 := stack.Top(1).(int32); top1 != 2 {
        t.Errorf("top1: %v", top1)
    }
    if top2 := stack.Top(2).(int32); top2 != 1 {
        t.Errorf("top2: %v", top2)
    }
}