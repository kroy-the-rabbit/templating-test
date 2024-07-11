
Go:

```
  > go mod init github.com/kroy-the-rabbit/templating-test
    go: creating new go.mod: module github.com/kroy-the-rabbit/templating-test
    go: to add module requirements and sums:
        go mod tidy
  > go mod tidy
  > go run main.go
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Hello MyName</title>
    </head>
    <body>
        <h1>Hello MyName</h1>
        <ul>
            
                <li>Item1</li>
            
                <li>Item2</li>
            
                <li>Item3</li>
            
        </ul>
    </body>
```
 
Python:

```
  > pyenv virtualenv 3.12 templating
  > pyenv local templating 
    (templating) 
  > pip install jinja2           
    Collecting jinja2
      Obtaining dependency information for jinja2 from https://files.pythonhosted.org/packages/31/80/3a54838c3fb461f6fec263ebf3a3a41771bd05190238de3486aae8540c36/jinja2-3.1.4-py3-none-any.whl.metadata
      Using cached jinja2-3.1.4-py3-none-any.whl.metadata (2.6 kB)
    Collecting MarkupSafe>=2.0 (from jinja2)
      Obtaining dependency information for MarkupSafe>=2.0 from https://files.pythonhosted.org/packages/0a/0d/2454f072fae3b5a137c119abf15465d1771319dfe9e4acbb31722a0fff91/MarkupSafe-2.1.5-cp312-cp312-manylinux_2_17_x86_64.manylinux2014_x86_64.whl.metadata
      Using cached MarkupSafe-2.1.5-cp312-cp312-manylinux_2_17_x86_64.manylinux2014_x86_64.whl.metadata (3.0 kB)
    Using cached jinja2-3.1.4-py3-none-any.whl (133 kB)
    Using cached MarkupSafe-2.1.5-cp312-cp312-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (28 kB)
    Installing collected packages: MarkupSafe, jinja2
    Successfully installed MarkupSafe-2.1.5 jinja2-3.1.4

    [notice] A new release of pip is available: 23.2.1 -> 24.1.2
    [notice] To update, run: python3.12 -m pip install --upgrade pip
  > python3 main.py
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Hello MyName</title>
    </head>
    <body>
        <h1>Hello MyName</h1>
        <ul>
            
                <li>Item1</li>
            
                <li>Item2</li>
            
                <li>Item3</li>
            
        </ul>
    </body>
    </html>
    (templating)  
``` 


