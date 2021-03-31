import React from "react"
import { InputText } from "../../components/InputText";
import { Message } from "../../components/Message"
import { Status } from "../../components/Status"
import "./styles.css"

const baseURL = 'ws://localhost:8080/chat';

export class Chat extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            ws: undefined,
            username: '',
            message: '',
            messages: [],
        }
        // this.handleEnter = this.handleEnter.bind(this)
    }

    render() {
        const { ws, messages } = this.state;

        return (
            <div className="chat">
                <h1>CyberChat</h1>
                <Status status={ws ? "connected" : "disconnected"}/>

                {
                    ws && <Message messages={messages}/>
                }

                <div className="chat-inputs">
                    <InputText
                        placeholder={ws ? "写下你要发送的消息" : "请输入你的昵称"}
                        onChange={value => ws ? this.setMessage(value) : this.setUsername(value) }
                        defaultValue={ws ? this.state.message : this.state.username}
                        send={this.handleEnter}
                    />
                    <button className="button" onClick={() => ws ? this.sendMessage() : this.enterChat()}>
                        {ws ? "Send" : "Enter"}
                    </button>
                </div>
            </div>
        )
    }

    handleEnter = (event) => {
        if (event.keyCode === 13) {
            const { ws } = this.state; 
            console.log(ws)
            ws ? this.sendMessage() : this.enterChat();
        }
    }

    enterChat() {
        const { username } = this.state;

        let ws = new WebSocket(baseURL + `?username=${username}`)

        ws.onopen = event => {
            console.log("WebSocket opened", {event})
        }

        ws.onclose = event => {
            console.log("WebSocket closed", {event})
        }

        ws.onmessage = msg => {
            console.log("WebSocket message", {msg})
            this.setMessages(msg.data);
        }

        ws.onerror = error => {
            console.log("WebSocket error", {error})
        }

        this.setState({ws, username: ''});
    }

    sendMessage() {
        const { ws, message } = this.state;
        if (message === '') {
            alert("消息不能为空!");
            return;
        }
        ws.send(message);
        this.setMessage('');
    }

    setUsername(value) {
        this.setState({
            username: value,
        })
    }

    setMessage(value) {
        this.setState({
            message: value,
        })
    }

    setMessages(value) {
        let messages = this.state.messages.concat([JSON.parse(value)]);
        this.setState({messages});
    }
}
