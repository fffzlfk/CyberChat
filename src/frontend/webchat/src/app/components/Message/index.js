import React from "react";

export class Message extends React.Component {
    render() {
        const { messages } = this.props;

        return messages && messages.map(msg => <p key={msg.ID}><strong>{msg.sender + ": "}</strong>{msg.body} <small>{msg.time}</small></p>)
    }
}