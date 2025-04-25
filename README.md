![GoHave_copy-removebg-preview](https://github.com/user-attachments/assets/c0bdb2b0-3194-48f0-a892-9fb10da22803)
<h1>GoWave - A Golang Backend freameworks</h1>
<br>
<br>
<h3>How to run the server?</h3>
<p>Just type in the terminal <code>go run main.go</code></p>
<br>
<h3>Where should we define routes?</h3>
<p>Go to the "/routes/routes.go". that is where all route should defined.</p>
<br>
<h3>How to Generate Controllers, Models, and middleware?</h3>
<p>Just type in the terminal <code>go run wave.go generate {controllers, or models, or middlewares} {controller name, or models name, or middlewares name}</code></p>
<p>Examples:</p>
<code>go run wave.go generate controllers User</code>
<br>
<code>go run wave.go generate model Users</code>
<br>
<code>go run wave.go generate middlewares Auth</code>