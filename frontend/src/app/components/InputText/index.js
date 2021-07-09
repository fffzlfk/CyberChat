import React from "react";
import "./styles.css"

export class InputText extends React.Component {
    render() {
        const { type, placeholder, onChange, defaultValue, send } = this.props;

        return (
            <input
                className="input"
                type={type ? type : "text"}
                placeholder={placeholder ? placeholder : ''}
                onChange={e => onChange(e.target.value)}
                value={defaultValue}
                onKeyDown={send}
            />
        )
    }
}