
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TDragObject struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewDragObject
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewDragObject() *TDragObject {
    d := new(TDragObject)
    d.instance = DragObject_Create()
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DragObjectFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func DragObjectFromInst(inst uintptr) *TDragObject {
    d := new(TDragObject)
    d.instance = inst
    d.ptr = unsafe.Pointer(inst)
    return d
}

// DragObjectFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func DragObjectFromObj(obj IObject) *TDragObject {
    d := new(TDragObject)
    d.instance = CheckPtr(obj)
    d.ptr = unsafe.Pointer(d.instance)
    return d
}

// DragObjectFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func DragObjectFromUnsafePointer(ptr unsafe.Pointer) *TDragObject {
    d := new(TDragObject)
    d.instance = uintptr(ptr)
    d.ptr = ptr
    return d
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (d *TDragObject) Free() {
    if d.instance != 0 {
        DragObject_Free(d.instance)
        d.instance = 0
        d.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (d *TDragObject) Instance() uintptr {
    return d.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (d *TDragObject) UnsafeAddr() unsafe.Pointer {
    return d.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (d *TDragObject) IsValid() bool {
    return d.instance != 0
}

// TDragObjectClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TDragObjectClass() TClass {
    return DragObject_StaticClassType()
}

// Assign
func (d *TDragObject) Assign(Source *TDragObject) {
    DragObject_Assign(d.instance, CheckPtr(Source))
}

// HideDragImage
func (d *TDragObject) HideDragImage() {
    DragObject_HideDragImage(d.instance)
}

// ShowDragImage
func (d *TDragObject) ShowDragImage() {
    DragObject_ShowDragImage(d.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (d *TDragObject) DisposeOf() {
    DragObject_DisposeOf(d.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (d *TDragObject) ClassType() TClass {
    return DragObject_ClassType(d.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (d *TDragObject) ClassName() string {
    return DragObject_ClassName(d.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (d *TDragObject) InstanceSize() int32 {
    return DragObject_InstanceSize(d.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (d *TDragObject) InheritsFrom(AClass TClass) bool {
    return DragObject_InheritsFrom(d.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (d *TDragObject) Equals(Obj IObject) bool {
    return DragObject_Equals(d.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (d *TDragObject) GetHashCode() int32 {
    return DragObject_GetHashCode(d.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (d *TDragObject) ToString() string {
    return DragObject_ToString(d.instance)
}

// AlwaysShowDragImages
func (d *TDragObject) AlwaysShowDragImages() bool {
    return DragObject_GetAlwaysShowDragImages(d.instance)
}

// SetAlwaysShowDragImages
func (d *TDragObject) SetAlwaysShowDragImages(value bool) {
    DragObject_SetAlwaysShowDragImages(d.instance, value)
}

// Cancelling
func (d *TDragObject) Cancelling() bool {
    return DragObject_GetCancelling(d.instance)
}

// SetCancelling
func (d *TDragObject) SetCancelling(value bool) {
    DragObject_SetCancelling(d.instance, value)
}

// DragHandle
func (d *TDragObject) DragHandle() HWND {
    return DragObject_GetDragHandle(d.instance)
}

// SetDragHandle
func (d *TDragObject) SetDragHandle(value HWND) {
    DragObject_SetDragHandle(d.instance, value)
}

// DragPos
func (d *TDragObject) DragPos() TPoint {
    return DragObject_GetDragPos(d.instance)
}

// SetDragPos
func (d *TDragObject) SetDragPos(value TPoint) {
    DragObject_SetDragPos(d.instance, value)
}

// DragTarget
func (d *TDragObject) DragTarget() uintptr {
    return DragObject_GetDragTarget(d.instance)
}

// SetDragTarget
func (d *TDragObject) SetDragTarget(value uintptr) {
    DragObject_SetDragTarget(d.instance, value)
}

// DragTargetPos
func (d *TDragObject) DragTargetPos() TPoint {
    return DragObject_GetDragTargetPos(d.instance)
}

// SetDragTargetPos
func (d *TDragObject) SetDragTargetPos(value TPoint) {
    DragObject_SetDragTargetPos(d.instance, value)
}

// Dropped
func (d *TDragObject) Dropped() bool {
    return DragObject_GetDropped(d.instance)
}

// MouseDeltaX
func (d *TDragObject) MouseDeltaX() float64 {
    return DragObject_GetMouseDeltaX(d.instance)
}

// MouseDeltaY
func (d *TDragObject) MouseDeltaY() float64 {
    return DragObject_GetMouseDeltaY(d.instance)
}

// RightClickCancels
func (d *TDragObject) RightClickCancels() bool {
    return DragObject_GetRightClickCancels(d.instance)
}

// SetRightClickCancels
func (d *TDragObject) SetRightClickCancels(value bool) {
    DragObject_SetRightClickCancels(d.instance, value)
}

