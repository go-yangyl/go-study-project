package builder

/**
简介：
创建者模式的核心思想是将一个"复杂对象的构建算法"与它的"部件及组装方式"分离，使得构件算法和组装方式可以独立应对变化；
复用同样的构建算法可以创建不同的表示，不同的构建过程可以复用相同的部件组装方式。

组成：
建造者（Builder）：对复杂对象的创建过程加以抽象，给出一个抽象接口，以规范产品对象的各个组成部分的建造。
具体创建者（ConcreateBuilder）：实现Builder接口，针对不同的业务逻辑，具体化复杂对象的各个部分的创建。在建造过程完成后，提供产品的实例
指导者（Dierctor）：调用具体建造者来创建复杂对象的各个部分，在指导者中不设计具体产品的信息，只负责保证对象各部分完整创建或者按某种顺序创建。
产品（Product）：要创建的复杂对象，一般来说包含多个部分。
*/

//Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}
