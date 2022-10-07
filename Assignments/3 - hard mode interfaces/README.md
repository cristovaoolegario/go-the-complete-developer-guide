# Hard mode assignment

Create a program that reads the contents of a text file then prints its contents to the terminal.

The file to open should be provided as an argument to the program when it is executed at the terminal. For example, 'go run main.go myfile.txt' should open the myfile.txt file.

To read in the arguments provided to a program, you can reference the variables 'os.Args', which is a slice of type string.

To open a file, check out the documentation for the 'Open' function in the 'os' package - <https://golang.org/pkg/os/#Open>

What interfaces does the 'File' type implement?

If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy function!
