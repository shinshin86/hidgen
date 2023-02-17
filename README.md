# hidgen
CLI tool for automatic heading tag ID generation.

## Install

```sh
go install github.com/shinshin86/hidgen@latest
```

## Usage
By default, only h2 and h3 are given the id

`index.html`

```html
<div>div tag</div>
<h1>h1 tag</h1>
<h2>h2 tag</h2>
<h3>h3 tag</h3>
<h4>h4 tag</h4>
<h5>h5 tag</h5>
<h6>h6 tag</h6>
```

```sh
# hidgen <input.html> <output.html>
hidgen index.html output.html
```

`output.html`

```html
<div>div tag</div>
<h1>h1 tag</h1>
<a id="h2 tag"></a>
<h2>h2 tag</h2>
<a id="h3 tag"></a>
<h3>h3 tag</h3>
<h4>h4 tag</h4>
<h5>h5 tag</h5>
<h6>h6 tag</h6>
```

### Assign IDs to all header tags

```sh
hidgen index.html output.html "h1,h2,h3,h4,h5,h6"
```


```html
<div>div tag</div>
<a id="h1 tag"></a>
<h1>h1 tag</h1>
<a id="h2 tag"></a>
<h2>h2 tag</h2>
<a id="h3 tag"></a>
<h3>h3 tag</h3>
<a id="h4 tag"></a>
<h4>h4 tag</h4>
<a id="h5 tag"></a>
<h5>h5 tag</h5>
<a id="h6 tag"></a>
<h6>h6 tag</h6>
```

### Specify any tag to replace

You can also specify any tag to replace, not just headings.

```sh
hidgen index.html output.html "div"
```

```html
<a id="div tag"></a>
<div>div tag</div>
<h1>h1 tag</h1>
<h2>h2 tag</h2>
<h3>h3 tag</h3>
<h4>h4 tag</h4>
<h5>h5 tag</h5>
<h6>h6 tag</h6>
```