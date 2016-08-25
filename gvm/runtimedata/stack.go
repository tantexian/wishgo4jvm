/*
	Copyright (c) 2015-2018 All rights reserved.
	本软件源代码版权归 my.oschina.net/tantexian 所有,允许复制与学习借鉴.
 */
package runtimedata
/*
    Description:
    	在运行Java程序时，Java虚拟机需要使用内存来存放各式各样的数据。
    	Java虚拟机规范把这些内存区域叫作运行时数据区。运行时数据区可以分为两类：
    	一类是多线程共享的，另一类则是线程私有的。
    	多线程共享的的运行时数据区需要在Java虚拟机启动时创建好，在Java虚拟机退出时销毁。
    	线程私有的运行时数据区则在创建线程时才创建，线程退出时销毁。

    	多线程共享的内存区域主要存放两类数据：类数据和类实例（也就是对象）。
    	对象数据存放在堆（Heap）中，类数据存放在方法区（Method Area）中。
    	堆由垃圾收集器定期清理，所以程序员不需要关心对象空间的释放。
    	类数据包括字段和方法信息、方法的字节码、运行时常量池，等等。
    	从逻辑上来讲，方法区其实也是堆的一部分。

    	线程私有的运行时数据区用于辅助执行Java字节码。每个线程都有自己的pc寄存器（Program Counter）和Java虚拟机栈（JVM Stack）。
    	Java虚拟机栈又由栈帧（Stack Frame，后面简称帧）构成，帧中保存方法执行的状态，包括局部变量表（Local Variable）
    	和操作数栈（Operand Stack）等。在任一时刻，某一线程肯定是在执行某个方法。这个方法叫作该线程的当前方法；执行该方法的帧叫作线程的当前帧；
    	声明该方法的类叫作当前类。如果当前方法是Java方法，则pc寄存器中存放当前正在执行的Java虚拟机指令的地址，
    	否则，当前方法是本地方法，pc寄存器中的值没有明确定义。

    	更加形象的图片描述请参考本项目中：wishgo4jvm/resource/stackAndFrame.png
    	或者博文：http://my.oschina.net/tantexian/blog/737610

    Author: tantexian
    Since: 2016/8/24
 */


/*
    Description: 结构：{JVM Stack [Frame][Frame][Frame]... }。
		JVM Stack在每个线程被创建时被创建，用来存放一组栈帧（StackFrame/Frame）。
		JVM Statck的大小可以是固定的，也可以是动态扩展的。如果线程需要一个比固定大小大的Stack，会发生StackOverflowError；
		如果动态扩展Stack时没有足够的内存或者系统没有足够的内存为新线程创建Stack，发生OutOfMemoryError。

		栈中包括了Frame及当前栈的maxSize
		用经典的链表（linked list）数据结构来实现Java虚拟机栈，
		这样栈就可以按需使用内存空间，而且弹出的帧也可以及时被Go的垃圾收集器回收。

    Author: tantexian
    Since:  2016/8/25
 */
type Stack struct {
	maxSize uint // 保存栈的容量（最多可以容纳多少帧）
	size uint  // 保存栈的当前大小
	head *Frame // 保存栈顶指针
}

// 创建Stack，最多可以容纳maxSize帧
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 把帧推入栈顶
func (self * Stack) push(frame *Frame){
	if self.size >= self.maxSize {
		panic("java.lang.StackOverfnextror")
	}

	if self.head != nil {
		frame.next = self.head
	}

	self.head = frame
	self.size++
}

// 弹出栈顶帧
func (self * Stack) pop() *Frame{
	if self.head == nil {
		panic("jvm stack is empty！")
	}
	top := self.head
	self.head = top.next
	top.next = nil
	self.size--
	return top
}

// 取出栈顶帧，并不弹出
func (self * Stack) top() *Frame{
	if self.head == nil {
		panic("jvm stack is empty！")
	}
	return self.head
}