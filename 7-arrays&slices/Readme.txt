Arrays:

	Arrays are used to store multiple values of the same type in a single variable,
	instead of declaring separate variables for each value.

	two ways to declare arrays in go:
		var array_name = [length]datatype{values} // here length is defined
		or
		var array_name = [...]datatype{values} // here length is inferred


Slices:

	Slices are similar to arrays, but are more powerful and flexible.
	Like arrays, slices are also used to store multiple values of the same type in a single variable.
	However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	In Go, there are several ways to create a slice:
		Using the []datatype{values} format
		Create a slice from an array
		Using the make() function