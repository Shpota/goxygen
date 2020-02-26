import React, {Component} from "react";
import axios from "axios";
import "./Tech.css"

export class Tech extends Component {
    state = {
        technologies: []
    };

    componentDidMount() {
        axios.get(`${process.env.REACT_APP_API_URL}/api/technologies`)
            .then(resp => this.setState({
                technologies: resp.data
            }));
    }

    render() {
        const technologies = this.state.technologies.map((technology, i) =>
            <li key={i}>
                <b>{technology.name}</b>: {technology.details}
            </li>
        );
        return (
            <ul className="technologies">
                {technologies}
            </ul>
        );
    }
}
