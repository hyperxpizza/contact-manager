import react from 'react';

const NewContactCard = () => {
    return (
        <div className="card">
            <i className="fa fa-plus fa-2x text-lightblue" aria-hidden="true"></i>
            <div className="card_inner">
                <p className="text-primary-p">Add a new contact</p>
            </div>
        </div>
    );
}

export default NewContactCard;