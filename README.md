# directory structure checker
## description
Check whether the directory structure defined by the user matches the actual directory structure.
## installation
```
$ git clone https://github.com/akisatoon1/directory-structure-checker.git
$ cd directory-structure-checker/src/
$ make
```
You cat get the binary named 'dir-struct-check' in directory-structure-checker/src/.
## how to use
This tool check the directory indicated by the first argument path. Pass the directory structure definition written in json to standard input. The syntax of the definition is explained in syntax seciton below. The result is true or false.
  
example:
```
dir-struct-check ./path/to/dir < structure_definition.json
```
output:
```
true
```
## syntax
The definition of directory structure is written in json. The key in json represents the file and directory name.
  
example json:
```
{
    "file1": null,
    "dir1": {
        "file01": null,
        "file02": null,
        "dir2": {
            "file001": null,
            "file002": null
        }
    }
}
```
### file
The key is file name, and the value must be null.
```
{
    "file name": null
}
```
### directory
The key is directory name, and the value must be object.
```
{
    "directory name": {
        "file1": null
    }
}
```
