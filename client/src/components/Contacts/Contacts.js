import { gql, useQuery } from '@apollo/client';
import ContactCard from '../Cards/ContactCard/ContactCard';

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


const Contacts = () => {
    const { loading, error, data } = useQuery(GET_CONTACTS);

    console.log(data.getContacts);

    return (
      <div>
      {data.getContacts.map(contact => (
        <ContactCard
          name1={contact.name1}
          name2={contact.name2}
          surname={contact.surname}
          email={contact.email}
          phone={contact.phone}
          website={contact.website}
          company={contact.company}
        />
      ))}
      </div>
    );


}

export default Contacts;