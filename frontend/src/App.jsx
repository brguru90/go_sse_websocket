import React, {Component} from "react"
import "./App.scss"
import * as $ from "jquery"
import SSEDemo from "./SSEDemo.jsx"

export default class App extends Component {
    state = {}

    client = null

    componentDidMount() {
        console.log("hi")
        let timeout = setTimeout(() => {
            $(".App").addClass("loaded")
            clearTimeout(timeout)
        }, 1000)
    }

    connect = () => {
        const resp = document.querySelector("pre")
        resp.innerHTML = "guru"

        try {
            this.client.close()
            this.client=null
        } catch (error) {
            
        }

        const client = new WebSocket("ws://localhost:8000/ws/chat/")

        client.onopen = function () {
            console.log("[open] Connection established")
            resp.innerHTML += "[open] Connection established\n"
            // client.send("My name is John");
        }

        client.onmessage = function (event) {
            console.log(`[message] Data received from server: ${event.data}`)
            resp.innerHTML += `${event.data}\n`
        }

        client.onclose = function (event) {
            if (event.wasClean) {
                console.log(
                    `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`
                )
                resp.innerHTML += `[close] Connection closed cleanly, code=${event.code} reason=${event.reason}\n`
            } else {
                // e.g. server process killed or network down
                // event.code is usually 1006 in this case
                console.log("[close] Connection died")
                resp.innerHTML += "[close] Connection died\n"
            }
        }

        client.onerror = function (error) {
            console.log(error)
            alert(`[error] ${error.message}`)
            resp.innerHTML += `[error] ${error.message}\n`
        }

        this.client = client
    }

    disconnect = () => {
        this.client.close()
    }

    send = () => {
        const text = document.querySelector("textarea").value

        try {
            this.client.send(text)
        } catch (error) {
            alert("client is not connected")
        }
        
    }

    render() {
        return (
            <div className="App">
                <br />
                <br />
                <button onClick={this.connect}>Connect</button>
                <button onClick={this.disconnect}>Disconnect</button>
                <br />
                <br />

                <textarea />
                <br />
                <br />
                <button onClick={this.send}>Send</button>
                <pre />

                <br /><br />
                <SSEDemo />
            </div>
        )
    }
}
