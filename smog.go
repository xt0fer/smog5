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

func (u *Universe) NewObject(clazzname string) *Object {
	o := new(Object)
	cls := u.GetGlobal(clazzname)
	o.Clazz = cls
	o.Hash = 0
	o.Fields = make([]*Object, 0)
	o.Primitive = nil
	return o
}

//
// VM Runtime Structs
//

type Frame struct {
	Stack         Stack[*Object]
	Locals        []*Object
	Args          []*Object
	MethodContext *Object // method
	Receiver      *Object
	Return        *Object
}

type Universe struct {
	Globals        map[string]*Object
	CallStack      Stack[*Frame]
	ExecutionStack Stack[*Object]
	// and the system class singletons
	ObjectClass *Object
	ClassClass  *Object
	NilClass    *Object
	TrueClass   *Object
	FalseClass  *Object
	//
	IntegerClass *Object
	DoubleClass  *Object
	StringClass  *Object
	//

	// system object singletons
	True  *Object
	False *Object
	Nil   *Object
}

func NewUniverse() *Universe {
	u := new(Universe)
	u.Globals = make(map[string]*Object)
	u.CallStack = *NewStack[*Frame]()
	u.ExecutionStack = *NewStack[*Object]()
	return u
}

func (u *Universe) Initialize() *Universe {
	u.ObjectClass = u.NewClass("Object")
	u.ClassClass = u.NewClass("Class")
	u.NilClass = u.NewClass("Nil")
	u.TrueClass = u.NewClass("True")
	u.FalseClass = u.NewClass("False")
	u.IntegerClass = u.NewClass("Integer")
	u.DoubleClass = u.NewClass("Double")
	u.StringClass = u.NewClass("String")
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

func (u *Universe) ClassFor(name string) *Object {
	return u.GetGlobal(name)
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

// OBJECT routines
func (o *Object) GetFieldAt(i int) *Object {
	return o.Fields[i]
}
func (o *Object) SetFieldAt(i int, value *Object) {
	o.Fields[i] = value
}

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

func (u *Universe) NewString(s string) *Object {
	o := u.NewObject("String")
	o.Primitive = s
	return o
}
func (u *Universe) NewSymbol(s string) *Object {
	o := u.NewObject("Symbol")
	o.Primitive = s
	return o
}

func (o *Object) String() string {
	if o.Clazz.Name == "String" || o.Clazz.Name == "Symbol" {
		return o.Primitive.(string)
	}
	return o.Clazz.Name
}
func (u *Universe) NewInteger(i int) *Object {
	o := u.NewObject("Integer")
	o.Primitive = i
	return o
}
func (u *Universe) NewDouble(d float64) *Object {
	o := u.NewObject("Double")
	o.Primitive = d
	return o
}
func (u *Universe) NewBlock(b *Object) *Object {
	o := u.NewObject("Block")
	o.Primitive = b
	return o
}

func (o *Object) Integer() int {
	if o.Clazz.Name == "Integer" {
		return o.Primitive.(int)
	}
	return 0
}
func (o *Object) Double() float64 {
	if o.Clazz.Name == "Double" {
		return o.Primitive.(float64)
	}
	return 0.0
}
func (o *Object) Block() *Object {
	if o.Clazz.Name == "Block" {
		return o.Primitive.(*Object)
	}
	return nil
}

func (o *Object) IsNil() bool {
	return o.Clazz.Name == "Nil"
}
func (o *Object) IsTrue() bool {
	return o.Clazz.Name == "True"
}
func (o *Object) IsFalse() bool {
	return o.Clazz.Name == "False"
}
func (o *Object) IsInteger() bool {
	return o.Clazz.Name == "Integer"
}
func (o *Object) IsDouble() bool {
	return o.Clazz.Name == "Double"
}
func (o *Object) IsString() bool {
	return o.Clazz.Name == "String"
}
func (o *Object) IsSymbol() bool {
	return o.Clazz.Name == "Symbol"
}
func (o *Object) IsBlock() bool {
	return o.Clazz.Name == "Block"
}
