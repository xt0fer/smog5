package smog5

// type Universe struct {
// 	NilObject *SObject
// 	TrueObject *SObject
// 	FalseObject *SObject
// 	ObjectClass *SClass
// 	ClassClass *SClass
// 	MetaclassClass *SClass
//   NilClass *SClass
//   IntegerClass *SClass
//   ArrayClass *SClass
//   MethodClass *SClass
//   SymbolClass *SClass
// 	PrimitiveClass *SClass
// 	StringClass *SClass
// 	SystemClass *SClass
// 	BlockClass *SClass
// 	DoubleClass *SClass
// 	  TrueClass *SClass
// 	  FalseClass *SClass
  
//   Globals map[*SSymbol]*SAbstractObject
  
//   ClassPath []string
// 	DumpBytecodes bool
//   PathSeparator string
//   FileSeparator string
//   Interpreter *Interpreter
//   SymbolTable map[string]*SSymbol
	
//   // 	// TODO: this is not how it is supposed to be... it is just a hack to cope
//   // 	// with the use of system.exit in SOM to enable testing
//   AvoidExit bool
//   LastExitCode int
//   }
  
func main() {
	// Create Universe
	u := NewUniverse()
	// Start interpretation
	u.Interpret(arguments)
	// Exit with error code 0
	u.Exit(0)
}

func (u *Universe) Interpret(arguments []string) {
	// Check for command line switches
	arguments = u.HandleArguments(arguments)
	// Initialize the known universe
	return u.Initialize(arguments)
}

func NewUniverse() *Universe {
	u := new(Universe)
	u.Interpreter = NewInterpreter(u)
	u.SymbolTable = make(map[string]*SSymbol)
	u.AvoidExit = false
	u.LastExitCode = 0
	return u
}

func NewUniverse(avoidExit bool) *Universe {
	u := new(Universe)
	u.Interpreter = NewInterpreter(u)
	u.SymbolTable = make(map[string]*SSymbol)
	u.AvoidExit = avoidExit
	u.LastExitCode = 0
	return u
}

  
func Current() *Universe {
	return current
}

func (u *Universe) GetInterpreter() *Interpreter {
	return u.Interpreter
}

func exit(errorCode int) {
	// Exit from the Java system
	if !avoidExit {
		System.exit(errorCode)
	} else {
		lastExitCode = errorCode
	}
}	

func (u *Universe) LastExitCode() int {
	return u.LastExitCode
}

func errorExit(message string) {
	errorPrintln("Runtime Error: " + message)
	exit(1)
}	

func (u *Universe) HandleArguments(arguments []string) {
	gotClasspath := false
	remainingArgs := make([]string, 0)
	// read dash arguments only while we haven't seen other kind of arguments
	sawOthers := false
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "-cp" && !sawOthers {
			if i + 1 >= len(arguments) {
				printUsageAndExit()
			}
			setupClassPath(arguments[i + 1])
			// Checkstyle: stop
			i++ // skip class path
			// Checkstyle: resume
			gotClasspath = true
		} else if arguments[i] == "-d" && !sawOthers {
			dumpBytecodes = true
		} else {
			sawOthers = true
			remainingArgs = append(remainingArgs, arguments[i])
		}
	}
	if !gotClasspath {
	// Get the default class path of the appropriate size
		classPath = setupDefaultClassPath(0)
	}
// 	  // check first of remaining args for class paths, and strip file extension
	if !remainingArgs.isEmpty() {
		split := getPathClassExt(remainingArgs[0])	 
		if split[0] != "" { // there was a path
			tmp := make([]string, len(classPath) + 1)
			copy(tmp, classPath)
			tmp[0] = split[0]
			classPath = tmp
		}
	return remainingArgs
	}
}
  
// 	// take argument of the form "../foo/Test.som" and return
// 	// "../foo", "Test", "som"
func getPathClassExt(arg string) []string {
	// Create a new tokenizer to split up the string of dirs
	tokenizer := new StringTokenizer(arg, fileSeparator, true)
	cp := ""
	for tokenizer.countTokens() > 2 {
		cp = cp + tokenizer.nextToken()
	}
	if tokenizer.countTokens() == 2 {
		tokenizer.nextToken() // throw out delimiter
	}
	file := tokenizer.nextToken()
	tokenizer = new StringTokenizer(file, ".")
	if tokenizer.countTokens() > 2 {
		println("Class with . in its name?")
		exit(1)
	}
	result := make([]string, 3)
	result[0] = cp
	result[1] = tokenizer.nextToken()
	result[2] = tokenizer.hasMoreTokens() ? tokenizer.nextToken() : ""
	return result
}

func setupClassPath(cp string) {
	// Create a new tokenizer to split up the string of dirs
	tokenizer := new StringTokenizer(cp, pathSeparator)
	classPath = make([]string, tokenizer.countTokens())
	// Get the directories and put them into the class path array
	for i := 0; tokenizer.hasMoreTokens(); i++ {
		classPath[i] = tokenizer.nextToken()
	}
}

func setupDefaultClassPath(directories int) []string {
	// Get the default system class path
	systemClassPath := System.getProperty("system.class.path")
	// Compute the number of defaults
	defaults := 1
	if systemClassPath != nil {
		defaults = 2
	}
	// Allocate an array with room for the directories and the defaults
	result := make([]string, directories + defaults)
	// Insert the system class path into the defaults section
	if systemClassPath != nil {
		result[directories] = systemClassPath
	}
	// Insert the current directory into the defaults section
	result[directories + defaults - 1] = "."

	return result
}	

func printUsageAndExit() {
	// Print the usage
	println("Usage: som [-options] [args...]                          ")
	println("                                                         ")
	println("where options include:                                   ")
	println("    -cp <directories separated by " + pathSeparator	+ ">")
	println("                  set search path for application classes")
	println("    -d            enable disassembling")

	// Exit
	System.exit(0)
}

// 	/**
// 	 * Start interpretation by sending the selector to the given class.
// 	 * This is mostly meant for testing currently.
// 	 *
// 	 * @param className
// 	 * @param selector
// 	 * @return
// 	 * @throws ProgramDefinitionError
// 	 */

func (u *Universe) Interpret(className string, selector string) (){
		u.InitializeObjectSystem()
		clazz := u.LoadClass(u.SymbolFor(className))
		// Lookup the initialize invokable on the system class
		initialize := clazz.GetSOMClass().LookupInvokable(u.SymbolFor(selector))
		if initialize == nil {
			panic("Lookup of " + className + ">>#" + selector + " failed")
		}
		return u.InterpretMethod(clazz, initialize, nil)
}

func (u *Universe) Initialize(arguments []string) *SAbstractObject {
	u.InitializeObjectSystem()
	if len(arguments) == 0 {
		shell := NewShell(u, u.Interpreter)
		bootstrapMethod := u.CreateBootstrapMethod()
		shell.SetBootstrapMethod(bootstrapMethod)
		return shell.Start()
	}
	initialize := u.SystemClass.LookupInvokable(u.SymbolFor("initialize:"))
	argumentsArray := u.NewArray(arguments)
	return u.InterpretMethod(u.SystemObject, initialize, argumentsArray)
}

func (u *Universe) CreateBootstrapMethod() *SMethod {
	bootstrapMethod := u.NewMethod(u.SymbolFor("bootstrap"), 1, 0, 2, nil)
	bootstrapMethod.SetBytecode(0, HALT)
	bootstrapMethod.SetHolder(u.SystemClass)
	return bootstrapMethod
}

func (u *Universe) InterpretMethod(receiver *SAbstractObject, invokable *SInvokable, arguments *SArray) *SAbstractObject {
	bootstrapMethod := u.CreateBootstrapMethod()
	bootstrapFrame := u.Interpreter.PushNewFrame(bootstrapMethod)
	bootstrapFrame.Push(receiver)
	if arguments != nil {
		bootstrapFrame.Push(arguments)
	}
	invokable.Invoke(bootstrapFrame, u.Interpreter)
	return u.Interpreter.Start()
}

func (u *Universe) InitializeObjectSystem() {
	// Allocate the nil object
	u.NilObject = NewSObject(nil)
	// Allocate the Metaclass classes
	u.MetaclassClass = NewMetaclassClass()
	// Allocate the rest of the system classes
	u.ObjectClass = NewSystemClass()
	u.NilClass = NewSystemClass()
	u.ClassClass = NewSystemClass()
	u.ArrayClass = NewSystemClass()
	u.SymbolClass = NewSystemClass()
	u.MethodClass = NewSystemClass()
	u.IntegerClass = NewSystemClass()
	u.PrimitiveClass = NewSystemClass()
	u.StringClass = NewSystemClass()
	u.DoubleClass = NewSystemClass()
	// Setup the class reference for the nil object
	u.NilObject.SetClass(u.NilClass)
// Initialize the system classes.
// Note: The order here is important.
	u.InitializeSystemClass(u.ObjectClass, nil, "Object")
	u.InitializeSystemClass(u.ClassClass, u.ObjectClass, "Class")
	u.InitializeSystemClass(u.MetaclassClass, u.ClassClass, "Metaclass")
	u.InitializeSystemClass(u.NilClass, u.ObjectClass, "Nil")
	u.InitializeSystemClass(u.ArrayClass, u.ObjectClass, "Array")
	u.InitializeSystemClass(u.MethodClass, u.ArrayClass, "Method")
	u.InitializeSystemClass(u.StringClass, u.ObjectClass, "String")
	u.InitializeSystemClass(u.SymbolClass, u.StringClass, "Symbol")
	u.InitializeSystemClass(u.IntegerClass, u.ObjectClass, "Integer")
	u.InitializeSystemClass(u.PrimitiveClass, u.ObjectClass, "Primitive")
	u.InitializeSystemClass(u.DoubleClass, u.ObjectClass, "Double")
	
// 	  // Load methods and fields into the system classes
	u.LoadSystemClass(u.ObjectClass)
	u.LoadSystemClass(u.ClassClass)
	u.LoadSystemClass(u.MetaclassClass)
	u.LoadSystemClass(u.NilClass)
	u.LoadSystemClass(u.ArrayClass)
	u.LoadSystemClass(u.MethodClass)
	u.LoadSystemClass(u.StringClass)
	u.LoadSystemClass(u.SymbolClass)
	u.LoadSystemClass(u.IntegerClass)
	u.LoadSystemClass(u.PrimitiveClass)
	u.LoadSystemClass(u.DoubleClass)
// 	  // Load the generic block class
  
// 	  // Fix up objectClass
	u.ObjectClass.SetSuperClass(u.NilObject)
// 	  // Load the generic block class
	u.BlockClass = u.LoadClass(u.SymbolFor("Block"), nil)

// Setup the true and false objects
	u.TrueSymbol := u.SymbolFor("True")
	u.TrueClass = u.LoadClass(u.TrueSymbol, nil)
	u.TrueObject = u.NewInstance(u.TrueClass)
  
	u.FalseSymbol := u.SymbolFor("False")
	u.FalseClass = u.LoadClass(u.FalseSymbol, nil)
	u.FalseObject = u.NewInstance(u.FalseClass)
  
// 	  // Load the system class and create an instance of it
	u.SystemClass = u.LoadClass(u.SymbolFor("System"), nil)
	u.SystemObject = u.NewInstance(u.SystemClass)
  
	// 	  // Put special objects and classes into the dictionary of globals
	u.SetGlobal(u.SymbolFor("nil"), u.NilObject)
	u.SetGlobal(u.SymbolFor("true"), u.TrueObject)
	u.SetGlobal(u.SymbolFor("false"), u.FalseObject)
	u.SetGlobal(u.SymbolFor("system"), u.SystemObject)
	u.SetGlobal(u.SymbolFor("System"), u.SystemClass)
	u.SetGlobal(u.SymbolFor("Block"), u.BlockClass)
	
	u.SetGlobal(u.TrueSymbol, u.TrueClass)
	u.SetGlobal(u.FalseSymbol, u.FalseClass)
}


func (u *Universe) SymbolFor(string string) *SSymbol {
	  result := u.SymbolTable[string]
	  if result != nil {
		  return result
	  }
	  result = NewSSymbol(string)
	  return result
  }

func (u *Universe) NewArray(length int) *SArray {
	  return NewSArray(u.NilObject, length)
  }

  func (u *Universe) NewArray(list []SAbstractObject) *SArray {
	result := u.NewArray(len(list))
	for i := 0; i < len(list); i++ {
		result.SetIndexableField(i, list[i])
	}
	return result
}

func (u *Universe) NewArray(stringArray []string) *SArray {
	result := u.NewArray(len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		result.SetIndexableField(i, u.NewString(stringArray[i]))
	}
	return result
} 

func (u *Universe) NewBlock(method *SMethod, context *Frame, arguments int) *SBlock {
	result := NewSBlock(method, context, u.GetBlockClass(arguments))
	return result
} 

func (u *Universe) NewClass(classClass *SClass) *SClass {
	result := NewSClass(classClass.GetNumberOfInstanceFields(), u)
	result.SetClass(classClass)
	return result
}

func (u *Universe) NewFrame(previousFrame *Frame, method *SMethod, context *Frame) *Frame {
	length := method.GetNumberOfArguments() + method.GetNumberOfLocals() + method.GetMaximumNumberOfStackElements() + 2
	result := NewFrame(u.NilObject, previousFrame, context, method, length)
	return result
}

func (u *Universe) NewMethod(signature *SSymbol, numberOfBytecodes int, numberOfLocals int, maxNumStackElements int, literals []SAbstractObject) *SMethod {
	result := NewSMethod(signature, numberOfBytecodes, numberOfLocals, maxNumStackElements, literals)
	return result
}

func (u *Universe) NewInstance(instanceClass *SClass) *SObject {
	result := NewSObject(instanceClass.GetNumberOfInstanceFields(), u.NilObject)
	result.SetClass(instanceClass)
	return result
}

func (u *Universe) NewInteger(value int) *SInteger {
	result := NewSInteger(value)
	return result
}
  
func (u *Universe) NewBigInteger(value *big.Int) *SBigInteger {
	result := NewSBigInteger(value)
	return result
}

func (u *Universe) NewDouble(value float64) *SDouble {
	result := NewSDouble(value)
	return result
}  

func (u *Universe) NewMetaclassClass() *SClass {
	result := NewSClass(u)
	result.SetClass(NewSClass(u))
	result.GetSOMClass().SetClass(result)
	return result
}
  
func (u *Universe) NewString(embeddedString string) *SString {
	result := NewSString(embeddedString) 
	return result
}
  
func (u *Universe) NewSymbol(string string) *SSymbol {
	result := NewSSymbol(string)
	u.SymbolTable[string] = result
	return result
}
  
func (u *Universe) NewSystemClass() *SClass {
	systemClass := NewSClass(u)
	systemClass.SetClass(NewSClass(u))
	systemClass.GetSOMClass().SetClass(u.MetaclassClass)
	return systemClass
}

func (u *Universe) InitializeSystemClass(systemClass *SClass, superClass *SClass, name string) {
	if superClass != nil {
		systemClass.SetSuperClass(superClass)
		systemClass.GetSOMClass().SetSuperClass(superClass.GetSOMClass())
	} else {
		systemClass.GetSOMClass().SetSuperClass(u.ClassClass)
	}
	systemClass.SetInstanceFields(u.NewArray(0))
	systemClass.GetSOMClass().SetInstanceFields(u.NewArray(0))
	systemClass.SetInstanceInvokables(u.NewArray(0))
	systemClass.GetSOMClass().SetInstanceInvokables(u.NewArray(0))
	systemClass.SetName(u.SymbolFor(name))
	systemClass.GetSOMClass().SetName(u.SymbolFor(name + " class"))
	u.SetGlobal(systemClass.GetName(), systemClass)
}
  
func (u *Universe) GetGlobal(name *SSymbol) *SAbstractObject {
	if u.HasGlobal(name) {
		return u.Globals[name]
	}
	return nil
}

func (u *Universe) SetGlobal(name *SSymbol, value *SAbstractObject) {
	u.Globals[name] = value
}

func (u *Universe) HasGlobal(name *SSymbol) bool {
	return u.Globals[name] != nil
}

func (u *Universe) GetBlockClass() *SClass {
	return u.BlockClass
}

func (u *Universe) GetBlockClass(numberOfArguments int) *SClass {
	name := u.SymbolFor("Block" + strconv.Itoa(numberOfArguments))
	if u.HasGlobal(name) {
		return u.GetGlobal(name)
	}
	result := u.LoadClass(name, nil)
	result.AddInstancePrimitive(SBlock.GetEvaluationPrimitive(numberOfArguments, u))
	u.SetGlobal(name, result)
	return result
}

func (u *Universe) LoadClass(name *SSymbol) *SClass {
	if u.HasGlobal(name) {
		return u.GetGlobal(name)
	}
	result := u.LoadClass(name, nil)
	if result != nil && result.HasPrimitives() {
		result.LoadPrimitives()
	}
	u.SetGlobal(name, result)
	return result
}
  
func (u *Universe) LoadSystemClass(systemClass *SClass) *SClass {
	result := u.LoadClass(systemClass.GetName(), systemClass)
	if result == nil {
		panic("Failed to load the " + systemClass.GetName().GetEmbeddedString() + " class. This is unexpected and may indicate that the classpath is not set correctly, or that the core library is not available.")
	}
	if result.HasPrimitives() {
		result.LoadPrimitives()
	}
	return result
}
 
func (u *Universe) LoadClass(name *SSymbol, systemClass *SClass) *SClass {
	for _, cpEntry := range u.ClassPath {
		result, err := CompileClass(cpEntry, name.GetEmbeddedString(), systemClass, u)
		if err != nil {
			continue
		}
		if u.DumpBytecodes {
			Disassembler.Dump(result.GetSOMClass(), u)
			Disassembler.Dump(result, u)
		}
		return result
	}
	return nil
}

func (u *Universe) loadShellClass(stmt string) *SClass {
	result, err := CompileClass(stmt, nil, u)
	if err != nil {
		errorExit(err.Error())
		panic(err)
	}
	if u.DumpBytecodes {
		Disassembler.Dump(result, u)
	}
	return result
}
  
func (u *Universe) ErrorPrint(msg string) {
	fmt.Print(msg)
}

func (u *Universe) ErrorPrintln(msg string) {
	fmt.Println(msg)
}

func (u *Universe) ErrorPrintln() {
	fmt.Println()
}

func (u *Universe) Print(msg string) {
	fmt.Print(msg)
}

func (u *Universe) Println(msg string) {
	fmt.Println(msg)
}

func (u *Universe) Println() {
	fmt.Println()
}
