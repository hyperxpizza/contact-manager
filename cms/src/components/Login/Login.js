import React, {useState} from 'react';
import { gql, useMutation} from '@apollo/client';

const LOGIN_MUTATION = gql`
    mutation login($username: String!, $password: String!){
        login(username: $username, password: $password){
            token
        }
    } 
`;

export default function Login({setToken}){
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    return(
        <div className="login-wrapper">
            <h1>Log in</h1>
            <form onSubmit>
                <label>
                    <p>Username</p>
                    <input type="text" onChange={e => setUsername(e.target.value)}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type="password" onChange={e => setPassword(e.target.value)}/>
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
            </form>
        </div>
    );
}