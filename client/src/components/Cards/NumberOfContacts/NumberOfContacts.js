import react, {useState} from 'react';
import { gql, useQuery } from '@apollo/client';

const COUNT_CONTACTS = gql`
    query countContacts{
        countContacts(filter:{}){
            count
        }
    }
`;

const NumberOfContacts = () => {

    const { loading, error, data } = useQuery(COUNT_CONTACTS);

    if(loading){
        return (
        <div className="card">
            <i className="fa fa-user-o fa-2x text-lightblue" aria-hidden="true"></i>
            <div className="card_inner">
                <p className="text-primary-p">Loading...</p>
            </div>
        </div>
        );
    }

    return (
        <div className="card">
            <i className="fa fa-user-o fa-2x text-lightblue" aria-hidden="true"></i>
            <div className="card_inner">
                <p className="text-primary-p">Number of Contacts</p>
                <span className="font-bold text-title">{data.countContacts.count}</span>
            </div>
        </div>
    );
    
}

export default NumberOfContacts;