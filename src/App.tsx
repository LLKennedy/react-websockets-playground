import React from 'react';
import logo from './logo.svg';
import './App.css';

export class App extends React.Component<{}, { lastMessage: string, conn: WebSocket }> {
  constructor(props: {}) {
    super(props);

    let conn = new WebSocket("ws://localhost:9943");

    this.state = {
      lastMessage: "No messages received",
      conn: conn
    };

    conn.addEventListener("open", (ev: Event) => {
      console.log("Opened Websocket: " + ev);
    });
    conn.addEventListener("error", ev => {
      console.log("Received error: " + ev);
    })
    conn.addEventListener("close", ev => {
      console.log("Received close: " + ev.code);
    })
  }
  componentDidMount() {
    this.state.conn.addEventListener("message", ev => {
      this.setState({
        lastMessage: `Waited ${JSON.parse(ev.data).delay}`,
      })
      // console.log("Received message: " + ev.data);
    })
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>
            {this.state.lastMessage}
          </p>
        </header>
      </div>
    );
  }

}

export default App;
