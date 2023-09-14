package smog5

// Path: smog.go
// smog5 is a package for the 5th generation of the smog programming language.

type Object struct {
	Hash      uint64
	Name      string
	Fields    []*Object
	Clazz     *Object
	Primitive interface{}
}

//
// VM Runtime Structs
//

// Heap is the one place where all objects are stored.
type HeapStruct struct {
	Objects []*Object
	Index   map[string]int // map of object names to heap indexes
}
// not sure I need this.

// Stack is used for FIFO in the VM.
type StackStruct struct {
	Stack []*Object
	PC    int
}

type Frame struct {
	StackStruct
	HeapStruct
	Locals []*Object
	Args   []*Object
}

type CallStack []*Frame
type Globals map[string]*Object

type Universe struct {
	Globals
	CallStack
	// and the system class singletons
	ObjectClass *Object
	ClassClass  *Object
	NilClass    *Object
	TrueClass   *Object
	FalseClass  *Object

	// system object singletons
	True  *Object
	False *Object
	Nil   *Object
}

func (u *Universe) NewObject(clazzname string) *Object {
	o := new(Object)
	cls := u.GetGlobal(clazzname)
	o.Clazz = cls
	o.Hash = 0
	o.Fields = make([]*Object, 0)
	o.Primitive = nil
	return o
}

func NewUniverse() *Universe {
	u := new(Universe)
	u.Globals = make(map[string]*Object)
	u.CallStack = make([]*Frame, 0)
	return u
}

func (u *Universe) Initialize() *Universe {
	u.ObjectClass = u.NewClass("Object")
	u.ClassClass = u.NewClass("Class")
	u.NilClass = u.NewClass("Nil")
	u.TrueClass = u.NewClass("True")
	u.FalseClass = u.NewClass("False")
	u.True = u.NewObject("True")
	u.False = u.NewObject("False")
	u.Nil = u.NewObject("Nil")
	return u
}

func (u *Universe) NewClass(name string) *Object {
	c := new(Object)
	c.Hash = hash(name)
	c.Name = name
	c.Fields = make([]*Object, 0)
	c.Primitive = nil
	u.Globals[name] = c
	return c
}

func hash(s string) uint64 {
	var h uint64
	for i := 0; i < len(s); i++ {
		h = h*31 + uint64(s[i])
	}
	return h
}

// Globals Symbol Table routines
func (u *Universe) GetGlobal(name string) *Object {
	cls, ok := u.Globals[name]
	if !ok {
		panic("Object not found: " + name)
	}
	return cls
}

func (u *Universe) SetGlobal(name string, value *Object) {
	u.Globals[name] = value
}

//
// OBJECT routines
//

func (o *Object) SetField(name string, value *Object) {
	for i, f := range o.Fields {
		if f.Clazz.Name == name {
			o.Fields[i] = value
			return
		}
	}
	o.Fields = append(o.Fields, value)
}

func (o *Object) GetField(name string) *Object {
	for _, f := range o.Fields {
		if f.Clazz.Name == name {
			return f
		}
	}
	return nil
}
