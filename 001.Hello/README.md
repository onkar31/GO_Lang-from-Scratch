<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
  <!-- Box around the title -->
  <div class="box">
    <h2>Hello World Program</h2>
  </div>

  <!-- Description -->
  <h3>Description</h3>
  <p>This is a simple "Hello, World!" program written in Go.</p>

  <!-- Usage -->
  <h3>Usage</h3>
  <p>To run the program, make sure you have Go installed on your system. Then, navigate to the directory containing<br> <code>hello-world.go</code> in your terminal and execute the following command:</p>
  <pre><code>go run hello-world.go</code></pre>
  <p>This will compile and execute the program, displaying "Hi there, Welcome to Go learning!" in the terminal.</p>

  <!-- About go.mod -->
  <h3>About <code>go.mod</code></h3>
  <p>The <code>go.mod</code> file is used in Go modules, introduced in Go 1.11, to manage dependencies for your project. It defines the module's path and lists its dependencies with their specific versions.</p>
  <p>If you're working on a larger project that requires external packages, you may use the <code>go mod init</code> command to initialize a new module and create the <code>go.mod</code> file. Then, you can use <code>go get</code> to add dependencies, and <code>go build</code> or <code>go run</code> to build and run your project.</p>

  <!-- About helloworld.exe -->
  <h3>About <code>hello-world.exe</code></h3>
  <p>After running <code>go build</code> in the terminal with the <code>hello-world.go</code> file, a binary executable file named <code>hello-world.exe</code> will be generated in the same directory. This executable file contains the compiled version of your Go program.</p>
  <p>To run the executable, simply execute it in the terminal:</p>
  <pre><code>./hello-world.exe</code></pre>
  <p>This will produce the same output as running the program with <code>go run hello-world.go</code>.</p>
</body>
</html>
