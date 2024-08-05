# Writing Test

```go
e := widget.NewEntry()
test.Type(e, "Hello")
assert.Equal(t, "Hello", e.Text)

test.DoubleTap(e)
assert.Equal(t, "Hello", e.SelectedText())
assert.Equal(t, 5, e.CursorColumn)
```

# Package
Commandline tool
```sh
go install fyne.io/fyne/v2/cmd/fyne@latest
```

Creates installer for the platform we are running on 
```sh
fyne package
```
We can also build for other platforms as below:
```sh
fyne package -os windows
```

Installs the installer on our system
```sh
fyne install
```

For android
```sh
fyne install -os android -appID com.mydomain.myProject
```

Use fine-cross tool instead of installing platform tools.

# Distribution
```sh
fyne release -appID com.mydomain.myProject
```

For ios:
```sh
fyne release -os ios -profile Xxx -certificate Yyy
```

For android:
```sh
fyne release -os android -keystore ~/Zzz.keystore
```

(Many options can be omitted using FyneApp.toml)

```toml
Website = "https://fyne.io"

[Details]
Icon = "img/fyne.png"
Name = "Fyne Demo"
ID = "io.fyne.demo"
Version = "1.0.0"
Build = 5
```

Checkout fyne_demo app for what's supported by it.
```sh
go install fyne.io/fyne/v2/cmd/fyne_demo@latest
```

# 3rd Party Components Too
Just import package and call constructor.
Works like any widget.
- ```map := xWidget.NewMap()```
- ```cmdline := terminal.New()```
- https://github.com/fyne-io/fyne-x

# Motivations
[apps.fyne.io](https://apps.fyne.io)
[app builder](https://fysion.app)

# Learn More
- https://docs.fyne.io
- https://www.youtube.com/@fyneio