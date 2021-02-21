import React from 'react';
import logo from './logo.svg';
import './App.css';


let conn = new WebSocket("localhost:6666");
conn.addEventListener("open", (ev: Event) => {
  console.log("Opened Websocket: " + ev);
});
conn.addEventListener("message", ev => {
  console.log("Received message: " + ev.data)
})

function App() {
  console.log()
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
