Go Maps:

Maps are used to store data values in key:value pairs.
Each element in a map is a key:value pair.
Maps are unordered data structures.
A map is an unordered and changeable collection that does not allow duplicates.
The length of a map is the number of its elements. You can find it using the len() function.
The default value of a map is nil.
Maps hold references to an underlying hash table.
Go has multiple ways for creating maps.

Note: The order of the map elements defined in the code is different from the way that they are stored. 
The data are stored in a way to have efficient data retrieval from the map.

NOTE:
maps are reference type
if any change occurs in derived ype reflecs same changes in base type also.

Note: The make()function is the right way to create an empty map. 
If you make an empty map in a different way and write to it, it will causes a runtime panic.


1. Creating Maps Using var and :=
Syntax:
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}


2. Creating Maps Using Using make()Function:
Syntax
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)



3. Check For Specific Elements in a Map:
You can check if a certain key exists in a map using
Syntax:
val, ok :=map_name[key]
If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.