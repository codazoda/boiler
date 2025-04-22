# Boiler

_Boiler_ is an application that automatically prints basic boilerplate code that you can copy and paste into your project.

![Boiler CLI Examples](docs/image/boiler.gif)

In the future, _Boiler_ may download it's templates from a public website. When we do that we'll use GitHub for contributing to the project. You should also be able to use your own URL as an additonal template repository.


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
sudo cp boiler /usr/local/bin/
```


## Template Paths

You must set a `BOILER` environment variable to your template directory and your template directory must end in a slash.

```
export BOILER=/some/template/path/
```
