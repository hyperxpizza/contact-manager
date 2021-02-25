import React from 'react';

const ContactCard = ({pic, name1, name2, surname, email, phone, website, company}) => {
    return (
        <div className="card">
            <div className="card_inner`">
                <h2>{name1} {name2} {surname}</h2>  
                <h5>Email: {email}</h5>
                <h5>Phone: {phone}</h5>
                <h5>Website: {website}</h5>
                <h5>Company: {company}</h5>
            </div>
        </div>
    );
}


export default ContactCard;