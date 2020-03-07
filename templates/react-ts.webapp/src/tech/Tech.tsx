import React from "react";
import axios from "axios";
import "./Tech.css"

export interface ITechnology {
    name: string;
    details: string;
}

export interface ITechProps {}

export interface ITechState {
    technologies: ITechnology[];
}

export class Tech extends React.Component<ITechProps, ITechState> {
    public state = {
        technologies: []
    };

    componentDidMount() {
        axios.get(`${process.env.REACT_APP_API_URL}/api/technologies`)
            .then(resp => this.setState({
                technologies: resp.data as ITechnology[]
            }));
    }

    render() {
        const technologies = this.state.technologies.map((technology: ITechnology, index: number) =>
            <li key={index}>
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
