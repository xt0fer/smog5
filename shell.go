package smog5

// public class Shell {

// 	private final Universe    universe;
// 	private final Interpreter interpreter;
  
// 	public Shell(final Universe universe, final Interpreter interpreter) {
// 	  this.universe = universe;
// 	  this.interpreter = interpreter;
// 	}
  
// 	private SMethod bootstrapMethod;
  
// 	public void setBootstrapMethod(SMethod method) {
// 	  bootstrapMethod = method;
// 	}
  
// 	public SAbstractObject start() {
  
// 	  BufferedReader in;
// 	  String stmt;
// 	  int counter;
// 	  int bytecodeIndex;
// 	  SClass myClass;
// 	  SAbstractObject myObject;
// 	  SAbstractObject it;
// 	  Frame currentFrame;
  
// 	  counter = 0;
// 	  in = new BufferedReader(new InputStreamReader(System.in));
// 	  it = universe.nilObject;
  
// 	  Universe.println("SOM Shell. Type \"quit\" to exit.\n");
  
// 	  // Create a fake bootstrap frame
// 	  currentFrame = interpreter.pushNewFrame(bootstrapMethod);
  
// 	  // Remember the first bytecode index, e.g. index of the HALT instruction
// 	  bytecodeIndex = currentFrame.getBytecodeIndex();
  
// 	  while (true) {
// 		try {
// 		  Universe.print("---> ");
  
// 		  // Read a statement from the keyboard
// 		  stmt = in.readLine();
// 		  if (stmt.equals("quit")) {
// 			return it;
// 		  }
  
// 		  // Generate a temporary class with a run method
// 		  stmt = "Shell_Class_" + counter++ + " = ( run: it = ( | tmp | tmp := ("
// 			  + stmt + " ). 'it = ' print. ^tmp println ) )";
  
// 		  // Compile and load the newly generated class
// 		  myClass = universe.loadShellClass(stmt);
  
// 		  // If success
// 		  if (myClass != null) {
// 			currentFrame = interpreter.getFrame();
  
// 			// Go back, so we will evaluate the bootstrap frames halt
// 			// instruction again
// 			currentFrame.setBytecodeIndex(bytecodeIndex);
  
// 			// Create and push a new instance of our class on the stack
// 			myObject = universe.newInstance(myClass);
// 			currentFrame.push(myObject);
  
// 			// Push the old value of "it" on the stack
// 			currentFrame.push(it);
  
// 			// Lookup the run: method
// 			SInvokable initialize = myClass.lookupInvokable(
// 				universe.symbolFor("run:"));
  
// 			// Invoke the run method
// 			initialize.invoke(currentFrame, interpreter);
  
// 			// Start the interpreter
// 			interpreter.start();
  
// 			// Save the result of the run method
// 			it = currentFrame.pop();
// 		  }
// 		} catch (Exception e) {
// 		  Universe.errorPrintln("Caught exception: " + e.getMessage());
// 		  Universe.errorPrintln("" + interpreter.getFrame().getPreviousFrame());
// 		}
// 	  }
// 	}
//   }
  