import {useState} from 'react';
import "./Main.css";
import { gql, useQuery } from '@apollo/client';
import NumberOfContacts from '../Cards/NumberOfContacts/NumberOfContacts';
import NewContactCard from '../Cards/NewContactCard/NewContactCard';
import Contacts from '../Contacts/Contacts';

const GET_CONTACTS = gql`
  query getContacts{
    getContacts(filter:{}){
      ObjectID
      name1
      name2
      surname
      email
      phone
      website
      company
      createdAt
      updatedAt
    }
  }
`;

const Main = () => {
    
  return (
    <main>
      <div className="main__container">
        {/* <!-- MAIN CARDS STARTS HERE --> */}
        <div className="main__cards">
          <NumberOfContacts />
          <NewContactCard />
        </div>
        {/* <!-- MAIN CARDS ENDS HERE --> */}

        {/* <!-- CHARTS STARTS HERE --> */}
        <div className="charts">
          <div className="charts__left">
            <div className="charts__left__title">
              <div>
                
              </div>
              <i className="fa fa-user-o fa-2x text-lightblue" aria-hidden="true"></i>
            </div>
            <Contacts />
          </div>

          <div className="charts__right">
            <div className="charts__right__title">
              <div>
                <h1>Stats Reports</h1>
                <p>Cupertino, California, USA</p>
              </div>
              <i className="fa fa-usd" aria-hidden="true"></i>
            </div>

            <div className="charts__right__cards">
              <div className="card1">
                <h1>Income</h1>
                <p>$75,300</p>
              </div>

              <div className="card2">
                <h1>Sales</h1>
                <p>$124,200</p>
              </div>

              <div className="card3">
                <h1>Users</h1>
                <p>3900</p>
              </div>

              <div className="card4">
                <h1>Orders</h1>
                <p>1881</p>
              </div>
            </div>
          </div>
        </div>
        {/* <!-- CHARTS ENDS HERE --> */}
      </div>
    </main>
  );
};

export default Main;