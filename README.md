# Boiler

_Boiler_ is an application that automatically prints basic boilerplate code that you can copy and paste into your project.

In the future, _Boiler_ will download it's templates from a public website. We'll use GitHub for contributing to the project. You'll also be able to use your own URL as an additonal template repository.


## Examples

A few examples of commands to show specific templates.

```
boiler html

boiler go

boiler go/http

boiler license/mit
```

Create an `index.html` file using the `html` template.

```
boiler html > index.html
```


## Installing

Clone this repo, build _Boiler_ for your platform, then copy (or symlink) the binary into your bin directory. You'll need Go installed to compile. After you have the code just run the following two commands to build and copy.

```
go build boiler.go
cp boiler /usr/local/bin/boiler
```


## Template Paths

The default template path is `~/boiler/template/`. You can put your own templates there or you can change the path by creating a `BOILER` environment variable. Your template directory must end in a slash.

```
export BOILER=/some/template/path/
```
